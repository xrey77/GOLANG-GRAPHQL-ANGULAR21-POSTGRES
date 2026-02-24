package gql

import (
	"golang_graphql_postgres/internal/gql/resolvers"
	"golang_graphql_postgres/internal/gql/schema/types"

	"github.com/graphql-go/graphql"
)

// GET ALL USERS
func GetUsersQuery() *graphql.Field {

	UserType := types.UserType
	return &graphql.Field{
		Type:        graphql.NewList(UserType),
		Description: "Get all users",
		Resolve:     resolvers.UserResolver,
	}
}

// GET USER BY ID
var userWrapperType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserWrapper",
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Type: types.UserType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source, nil
			},
		},
	},
})

func GetUserByIDQuery() *graphql.Field {
	return &graphql.Field{
		Type:        userWrapperType,
		Description: "Get user by ID",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: resolvers.UserByIDResolver,
	}
}

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"users": GetUsersQuery(),
		"user":  GetUserByIDQuery(),
	},
})
