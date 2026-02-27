package queries

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
		Resolve:     resolvers.GetUsersResolver,
	}
}

// ============REQUEST=============
// query GetUsers {
//     getUsers {
//         id
//         firstname
//         lastname
//         email
//         mobile
//         username
//         isactivated
//         isblocked
//         userpicture
//     }
// }
