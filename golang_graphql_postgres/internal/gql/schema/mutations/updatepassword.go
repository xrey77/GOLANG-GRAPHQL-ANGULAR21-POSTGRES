package mutations

import (
	"errors"
	"fmt"
	"golang_graphql_postgres/configs"
	"golang_graphql_postgres/internal/dto"
	"golang_graphql_postgres/internal/gql/schema/types"
	"golang_graphql_postgres/internal/models"

	"github.com/graphql-go/graphql"
	"golang.org/x/crypto/bcrypt"
)

var passwordResponseType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PasswordResponse",
	Fields: graphql.Fields{
		"message": &graphql.Field{Type: graphql.String},
		"user":    &graphql.Field{Type: types.UserType},
	},
})

var UpdatePasswordField = &graphql.Field{
	Type:        passwordResponseType,
	Description: "Update user profile",
	Args: graphql.FieldConfigArgument{
		"id":       &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		"password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		user, ok := params.Context.Value("user").(*types.UserClaims)

		if !ok || user == nil {
			return nil, errors.New("Unauthorized Access, valid bearer token required.")
		}

		userid := params.Args["id"].(int)
		password, _ := params.Args["password"].(string)

		hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}

		userProfile := dto.UserDTO{
			Password: string(hashedPass),
		}

		result := configs.DB.Model(&models.User{}).Where("id = ?", userid).Updates(userProfile)

		if result.Error != nil {
			return nil, fmt.Errorf("%s", "Update failed: "+result.Error.Error())
		}

		if result.RowsAffected == 0 {
			return nil, fmt.Errorf("Update failed: User ID not found")
		}

		return map[string]interface{}{
			"message": "You have changed your password successfully.",
			"user":    userProfile,
		}, nil
	},
}

// ==========REQUEST================
// mutation UpdatePassword($id : Int!, $password: String!) {
//     updatePassword(id: $id, password: $password) {
//         message
//         user {
//             id
//         }
//     }
// }

// ===========VARIABLES=============
// {
//   "id": 31,
//   "password": "nald"
// }
