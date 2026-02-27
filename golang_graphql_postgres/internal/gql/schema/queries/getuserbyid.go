package queries

import (
	"golang_graphql_postgres/internal/gql/resolvers"
	"golang_graphql_postgres/internal/gql/schema/types"

	"github.com/graphql-go/graphql"
)

// GET USER BY ID
var userIdWrapperType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserIdWrapper",
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Type: types.UserType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source, nil
			},
		},
	},
})

func GetUserIDQuery() *graphql.Field {
	return &graphql.Field{
		Type:        userIdWrapperType,
		Description: "Get user by ID",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: resolvers.UserIDResolver,
	}
}

// ===============REQUEST================
// query GetUserID($id: Int!) {
//     getUserID(id: $id) {
//         user {
//             id
//             firstname
//             lastname
//             email
//             mobile
//             username
//         }
//     }
// }

// ============VARIABLES================
// {
//     "id": 1
// }
