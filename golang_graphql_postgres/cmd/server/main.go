package main

import (
	"fmt"
	"golang_graphql_postgres/configs"
	mutate "golang_graphql_postgres/internal/gql/schema/mutations"
	gql "golang_graphql_postgres/internal/gql/schema/queries"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func GraphQLHandler(schema graphql.Schema) gin.HandlerFunc {
	return func(c *gin.Context) {
		var params struct {
			Query         string                 `json:"query"`
			Variables     map[string]interface{} `json:"variables"`
			OperationName string                 `json:"operationName"`
		}

		if err := c.ShouldBindJSON(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		result := graphql.Do(graphql.Params{
			Schema:         schema,
			RequestString:  params.Query,
			VariableValues: params.Variables,
			OperationName:  params.OperationName,
			Context:        c.Request.Context(),
		})

		c.JSON(http.StatusOK, result)
	}
}

func main() {
	configs.Connection()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	schemaConfig := graphql.SchemaConfig{Query: gql.RootQuery, Mutation: mutate.RootMutation}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

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

	host := "localhost"
	port := "5000"
	address := fmt.Sprintf("%s:%s", host, port)
	log.Print("Listening to ", address)
	log.Fatal(http.ListenAndServe(":5000", router))
}
