package resolvers

import (
	"golang_graphql_postgres/configs"
	"golang_graphql_postgres/internal/models"

	"github.com/graphql-go/graphql"
)

func UserResolver(p graphql.ResolveParams) (interface{}, error) {
	var users []models.User
	if err := configs.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// ============REQUEST=============
// query Users {
//     users {
//         id,
//         firstname
//         lastname
//         email
//         mobile
//         isactivated
//         isblocked
//         userpicture
//         secret
//         qrcodeurl
//     }
// }

func UserByIDResolver(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(int)
	if !ok {
		return nil, nil
	}

	var user models.User
	if err := configs.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// USER REGISTRATION
// func CreateUserMutaion((p graphql.ResolveParams) (interface{}, error) {

// }
