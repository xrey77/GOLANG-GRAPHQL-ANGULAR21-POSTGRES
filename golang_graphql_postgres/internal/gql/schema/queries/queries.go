package queries

import (
	"github.com/graphql-go/graphql"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"getUsers":      GetUsersQuery(),
		"getUserID":     GetUserIDQuery(),
		"productList":   ProductListQuery(),
		"productSearch": ProductSearchQuery(),
		"getSales":      GetSalesQuery(),
		"pdfQuery":      PdfaQuery(),
	},
})
