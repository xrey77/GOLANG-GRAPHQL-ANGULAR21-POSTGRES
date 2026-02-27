package types

import "github.com/graphql-go/graphql"

var SaleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Sale",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: graphql.Int},
		"saleamount": &graphql.Field{Type: DecimalScalar},
		"salesdate":  &graphql.Field{Type: graphql.DateTime},
	},
},
)
