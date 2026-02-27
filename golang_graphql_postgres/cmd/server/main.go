package main

import (
	"fmt"
	"golang_graphql_postgres/configs"
	mutate "golang_graphql_postgres/internal/gql/schema/mutations"
	gql "golang_graphql_postgres/internal/gql/schema/queries"
	middleware "golang_graphql_postgres/internal/middleware"
	"log"
	"net/http"

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

func GraphQLHandler(schema graphql.Schema) gin.HandlerFunc {
	return func(c *gin.Context) {
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
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

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
