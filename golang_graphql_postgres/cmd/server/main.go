package main

import (
	"context"
	"encoding/json"
	"fmt"
	"golang_graphql_postgres/configs"
	mutate "golang_graphql_postgres/internal/gql/schema/mutations"
	gql "golang_graphql_postgres/internal/gql/schema/queries"
	middleware "golang_graphql_postgres/internal/middleware"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/joho/godotenv"
)

func init() {
	configs.Connection()
	err1 := godotenv.Load(".env")
	if err1 != nil {
		log.Fatalf("Error loading .env file")
	}

}

// func GraphQLHandler(schema graphql.Schema) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var requestBody struct {
// 			Query         string                 `json:"query"`
// 			Variables     map[string]interface{} `json:"variables"`
// 			OperationName string                 `json:"operationName"`
// 		}
// 		//=======
// 		form, err := c.MultipartForm()
// 		if err != nil {
// 			// Fallback to standard JSON if not multipart
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid multipart."})
// 		}

// 		if err := c.BindJSON(&requestBody); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
// 			return
// 		}

// 		var params struct {
// 			Query         string                 `json:"query" form:"query"`
// 			OperationName string                 `json:"operationName" form:"operationName"`
// 			Variables     map[string]interface{} `json:"variables" form:"variables"`
// 		}

// 		// 2. Parse variables (often sent as a JSON string in multipart)
// 		if varStr := c.PostForm("variables"); varStr != "" {
// 			json.Unmarshal([]byte(varStr), &params.Variables)
// 		}

// 		if form != nil && len(form.File["0"]) > 0 {
// 			params.Variables["userpicture"] = form.File["0"][0]
// 		}
// 		//======
// 		result := graphql.Do(graphql.Params{
// 			Schema:         schema,
// 			RequestString:  requestBody.Query,
// 			VariableValues: requestBody.Variables,
// 			OperationName:  requestBody.OperationName,
// 			Context:        context.WithValue(c.Request.Context(), "GinContext", c),
// 			// Context:        c.Request.Context(),
// 		})

// 		c.JSON(http.StatusOK, result)
// 	}
// }

func GraphQLHandler(schema graphql.Schema) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Initialize a single struct to hold data from any source
		var params struct {
			Query         string                 `json:"query" form:"query"`
			Variables     map[string]interface{} `json:"variables"`
			OperationName string                 `json:"operationName" form:"operationName"`
		}

		// 2. Determine the source based on Content-Type
		contentType := c.ContentType()

		if contentType == "application/json" {
			if err := c.ShouldBindJSON(&params); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON body"})
				return
			}
		} else {
			// Handle Multipart or standard Form
			params.Query = c.PostForm("query")
			params.OperationName = c.PostForm("operationName")

			// Parse Variables string from form if it exists
			if varStr := c.PostForm("variables"); varStr != "" {
				json.Unmarshal([]byte(varStr), &params.Variables)
			}
		}

		// 3. Initialize Variables map if it's still nil
		if params.Variables == nil {
			params.Variables = make(map[string]interface{})
		}

		// 4. Handle File Uploads (Multipart specific)
		form, _ := c.MultipartForm()
		if form != nil && len(form.File["0"]) > 0 {
			params.Variables["userpicture"] = form.File["0"][0]
		}

		// 5. Execute GraphQL with the unified 'params'
		result := graphql.Do(graphql.Params{
			Schema:         schema,
			RequestString:  params.Query,
			VariableValues: params.Variables,
			OperationName:  params.OperationName,
			Context:        context.WithValue(c.Request.Context(), "GinContext", c),
		})

		c.JSON(http.StatusOK, result)
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Apollo-Require-Preflight"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Use(static.Serve("/", static.LocalFile(",/templates", true)))
	router.LoadHTMLGlob("./templates/*.*")
	router.Static("/assets", "./templates/assets")

	schemaConfig := graphql.SchemaConfig{Query: gql.RootQuery, Mutation: mutate.RootMutation}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	router.Use(middleware.AuthMiddleware())
	router.POST("/graphql", func(c *gin.Context) {
		var requestBody struct {
			Query         string                 `json:"query"`
			Variables     map[string]interface{} `json:"variables"`
			OperationName string                 `json:"operationName"`
		}

		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		result := graphql.Do(graphql.Params{
			Schema:         schema,
			RequestString:  requestBody.Query,
			VariableValues: requestBody.Variables,
			OperationName:  requestBody.OperationName,
			Context:        c.Request.Context(),
		})

		c.JSON(http.StatusOK, result)
	})

	router.GET("/", func(c *gin.Context) {
		c.File("./templates/index.html")
	})

	host := "localhost"
	port := "5000"
	address := fmt.Sprintf("%s:%s", host, port)
	log.Print("Listening to ", address)
	log.Fatal(http.ListenAndServe(":5000", router))
}
