package resolvers

import (
	"golang_graphql_postgres/configs"
	"golang_graphql_postgres/internal/models"

	"github.com/graphql-go/graphql"
)

func GetSalesResolver(p graphql.ResolveParams) (interface{}, error) {

	var sales []models.Sale
	if err := configs.DB.Find(&sales).Error; err != nil {
		return nil, err
	}
	return &sales, nil
}

// ==================REQUEST=====================
// query GetSales {
//     getSales {
//         saleamount
//         salesdate
//     }
// }
