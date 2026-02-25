package queries

import (
	"github.com/graphql-go/graphql"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"getUsers":    GetUsersQuery(),
		"getUserByID": GetUserByIDQuery(),
	},
})
