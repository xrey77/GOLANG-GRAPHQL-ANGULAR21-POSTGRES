package types

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/shopspring/decimal"
)

var DecimalScalar = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Decimal",
	Description: "High-precision decimal",
	Serialize: func(value interface{}) interface{} {
		switch v := value.(type) {
		case decimal.Decimal:
			return v.String()
		case *decimal.Decimal:
			return v.String()
		default:
			return nil
		}
	},
	ParseValue: func(value interface{}) interface{} {
		if v, ok := value.(string); ok {
			d, err := decimal.NewFromString(v)
			if err != nil {
				return nil
			}
			return d
		}
		return nil
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		if v, ok := valueAST.(*ast.StringValue); ok {
			d, err := decimal.NewFromString(v.Value)
			if err != nil {
				return nil
			}
			return d
		}
		return nil
	},
})

var ProductType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Product",
	Fields: graphql.Fields{
		"id":             &graphql.Field{Type: graphql.Int},
		"category":       &graphql.Field{Type: graphql.String},
		"descriptions":   &graphql.Field{Type: graphql.String},
		"qty":            &graphql.Field{Type: graphql.Int},
		"unit":           &graphql.Field{Type: graphql.String},
		"costprice":      &graphql.Field{Type: DecimalScalar},
		"sellprice":      &graphql.Field{Type: DecimalScalar},
		"saleprice":      &graphql.Field{Type: DecimalScalar},
		"productpicture": &graphql.Field{Type: graphql.String},
		"alertstocks":    &graphql.Field{Type: graphql.Int},
		"criticalstocks": &graphql.Field{Type: graphql.Int},
	},
},
)
