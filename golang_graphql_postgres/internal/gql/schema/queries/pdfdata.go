package queries

import (
	"golang_graphql_postgres/internal/gql/resolvers"
	"golang_graphql_postgres/internal/gql/schema/types"

	"github.com/graphql-go/graphql"
)

func PdfaQuery() *graphql.Field {

	ProductType := types.ProductType
	return &graphql.Field{
		Type:        graphql.NewList(ProductType),
		Description: "Get PDF data",
		Resolve:     resolvers.GetPdfResolver,
	}
}

// ======REQUEST======
// query GetSales {
//     getSales {
//         saleamount
//         salesdate
//     }
// }
