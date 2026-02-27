package queries

import (
	"golang_graphql_postgres/internal/gql/resolvers"
	"golang_graphql_postgres/internal/gql/schema/types"

	"github.com/graphql-go/graphql"
)

var searchResponseType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ProductResponse",
	Fields: graphql.Fields{
		"page":         &graphql.Field{Type: graphql.Int},
		"key":          &graphql.Field{Type: graphql.String},
		"totalpage":    &graphql.Field{Type: graphql.Int},
		"totalrecords": &graphql.Field{Type: graphql.Int},
		"products":     &graphql.Field{Type: graphql.NewList(types.ProductType)},
	},
})

func ProductSearchQuery() *graphql.Field {
	return &graphql.Field{
		Type:        productResponseType,
		Description: "Get all products",
		Args: graphql.FieldConfigArgument{
			"page": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			"key":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: resolvers.GetSearchResolver,
	}
}
