package resolvers

import (
	"golang_graphql_postgres/configs"
	"golang_graphql_postgres/internal/models"

	"github.com/graphql-go/graphql"
)

func GetPdfResolver(p graphql.ResolveParams) (interface{}, error) {

	var products []models.Product
	if err := configs.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return &products, nil
}

// // ==================REQUEST=====================
// query PdfQuery {
//     pdfQuery {
//         id
//         category
//         descriptions
//         qty
//         unit
//         costprice
//         sellprice
//         saleprice
//         productpicture
//         alertstocks
//         criticalstocks
//     }
// }
