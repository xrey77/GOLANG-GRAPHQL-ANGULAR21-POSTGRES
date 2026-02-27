package queries

import (
	"golang_graphql_postgres/internal/gql/resolvers"
	"golang_graphql_postgres/internal/gql/schema/types"

	"github.com/graphql-go/graphql"
)

func GetSalesQuery() *graphql.Field {

	SaleType := types.SaleType
	return &graphql.Field{
		Type:        graphql.NewList(SaleType),
		Description: "Get sales data",
		Resolve:     resolvers.GetSalesResolver,
	}
}
