package queries

import (
	"golang_graphql_postgres/internal/gql/resolvers"
	"golang_graphql_postgres/internal/gql/schema/types"

	"github.com/graphql-go/graphql"
)

var productResponseType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ProductResponse",
	Fields: graphql.Fields{
		"page":         &graphql.Field{Type: graphql.Int},
		"totalpage":    &graphql.Field{Type: graphql.Int},
		"totalrecords": &graphql.Field{Type: graphql.Int},
		"products":     &graphql.Field{Type: graphql.NewList(types.ProductType)},
		// "products": &graphql.Field{Type: types.ProductType},
	},
})

func ProductListQuery() *graphql.Field {
	return &graphql.Field{
		Type:        productResponseType,
		Description: "Get all products",
		Args: graphql.FieldConfigArgument{
			"page": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: resolvers.GetProductsResolver,
	}
}

// ========REQUEST================
// query ProductList($page: Int!) {
//     productList(page: $page) {
//         products {
//             id
//             category
//             descriptions
//             qty
//             unit
//             costprice
//             sellprice
//             saleprice
//             productpicture
//             alertstocks
//             criticalstocks
//         }
//     }
// }

// ======VARIABLES========
// {
// 	"page": 1
// }
