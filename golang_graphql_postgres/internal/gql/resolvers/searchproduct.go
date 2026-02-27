package resolvers

import (
	"errors"
	"fmt"
	"golang_graphql_postgres/configs"
	"golang_graphql_postgres/internal/models"
	"math"
	"strconv"
	"strings"

	"github.com/graphql-go/graphql"
)

func GetSearchResolver(p graphql.ResolveParams) (interface{}, error) {

	xkey, ok := p.Args["key"].(string)
	if !ok {
		return nil, fmt.Errorf("argument 'key' must be a string")
	}
	key := strings.ToLower((xkey))
	keyWord := "%" + key + "%"

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
	var products []models.Product
	result := configs.DB.Where("LOWER(descriptions) ILIKE ?", keyWord).Find(&products)

	totrecs := result.RowsAffected
	total1 := float64(totrecs) / float64(perPage)
	totalPages := math.Ceil(total1)
	offset := (page - 1) * perPage

	var productData []models.Product
	if err := configs.DB.Where("descriptions ILIKE ?", keyWord).Order("id").Limit(perPage).Offset(offset).Find(&productData).Error; err != nil {
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
// query ProductSearch($key: String!, $page: Int!) {
//     productSearch(key: $key, page: $page) {
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
//     "key": "T500",
//     "page": 1
// }
