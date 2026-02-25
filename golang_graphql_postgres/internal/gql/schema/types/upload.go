package types

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

var UploadScalar = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Upload",
	Description: "The `Upload` scalar type represents a file upload.",
	Serialize: func(value interface{}) interface{} {
		return value
	},
	ParseValue: func(value interface{}) interface{} {
		return value
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		return nil
	},
})
