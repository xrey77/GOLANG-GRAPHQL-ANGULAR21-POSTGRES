package resolvers

import (
	"errors"
	"golang_graphql_postgres/configs"
	"golang_graphql_postgres/internal/gql/schema/types"
	"golang_graphql_postgres/internal/models"

	"github.com/graphql-go/graphql"
)

func GetUsersResolver(p graphql.ResolveParams) (interface{}, error) {
	user, ok := p.Context.Value("user").(*types.UserClaims)

	if !ok || user == nil {
		return nil, errors.New("unauthorized: valid bearer token required")
	}

	var users []models.User
	if err := configs.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func UserIDResolver(p graphql.ResolveParams) (interface{}, error) {
	user, ok := p.Context.Value("user").(*types.UserClaims)

	if !ok || user == nil {
		return nil, errors.New("unauthorized: valid bearer token required")
	}

	idRaw := p.Args["id"]
	var id int
	switch v := idRaw.(type) {
	case int:
		id = v
	case string:
		// Convert if necessary: id, _ = strconv.Atoi(v)
	default:
		return nil, errors.New("invalid id type")
	}

	var userData models.User // Avoid naming the variable 'users' for a single result
	if err := configs.DB.First(&userData, id).Error; err != nil {
		return nil, err
	}
	return &userData, nil
}
