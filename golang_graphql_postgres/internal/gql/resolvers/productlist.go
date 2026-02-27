package resolvers

import (
	"errors"
	"fmt"
	"golang_graphql_postgres/configs"
	"golang_graphql_postgres/internal/models"
	"math"
	"strconv"

	"github.com/graphql-go/graphql"
)

func GetProductsResolver(p graphql.ResolveParams) (interface{}, error) {

	pageRaw := p.Args["page"]
	var page int
	switch v := pageRaw.(type) {
	case int:
		page = v
	case string:
		page, _ = strconv.Atoi(v)
	default:
		return nil, errors.New("invalid page type")
	}

	perPage := 5
	db := configs.Connection()
	var products []models.Product
	result := db.Find(&products)
	totrecs := result.RowsAffected
	total1 := float64(totrecs) / float64(perPage)
	totalPages := math.Ceil(total1)
	offset := (page - 1) * perPage

	var productData []models.Product
	if err := configs.DB.Order("id").Limit(perPage).Offset(offset).Find(&productData).Error; err != nil {
		return nil, err
	}

	if len(productData) == 0 {
		return nil, fmt.Errorf("%s", "No record(s) found.")
	}

	return map[string]interface{}{
		"page":         page,
		"totalpage":    totalPages,
		"totalrecords": totrecs,
		"products":     productData,
	}, nil
}

// ===================REQUEST==================
// query ProductList($page : Int!) {
//     productList(page: $page) {
//         page
//         totalpage
//         totalrecords
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

// ==================VARIABLES================
// {
//     "page": 1
// }
