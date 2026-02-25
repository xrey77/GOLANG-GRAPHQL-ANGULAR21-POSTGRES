package queries

import (
	"golang_graphql_postgres/internal/gql/resolvers"
	"golang_graphql_postgres/internal/gql/schema/types"

	"github.com/graphql-go/graphql"
)

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

// ===============REQUEST================
// query GetUsers {
//       getUsers {
//         id
//         firstname
//         lastname
//         email
//         mobile
//         userpicture
//         isactivated
//         isblocked
//         userpicture
//         qrcodeurl
//       }
// }

// ============VARIABLES================
// {
//     "id": 1
// }
