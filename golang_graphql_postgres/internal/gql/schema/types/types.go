package types

import (
	"github.com/graphql-go/graphql"
)

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.ID},
		"firstname":   &graphql.Field{Type: graphql.String},
		"lastname":    &graphql.Field{Type: graphql.String},
		"email":       &graphql.Field{Type: graphql.String},
		"mobile":      &graphql.Field{Type: graphql.String},
		"username":    &graphql.Field{Type: graphql.String},
		"password":    &graphql.Field{Type: graphql.String},
		"roles":       &graphql.Field{Type: graphql.String},
		"isactivated": &graphql.Field{Type: graphql.Int},
		"isblocked":   &graphql.Field{Type: graphql.Int},
		"mailtoken":   &graphql.Field{Type: graphql.Int},
		"userpicture": &graphql.Field{Type: graphql.String},
		"secret":      &graphql.Field{Type: graphql.String},
		"qrcodeurl":   &graphql.Field{Type: graphql.String},
	},
},
)
