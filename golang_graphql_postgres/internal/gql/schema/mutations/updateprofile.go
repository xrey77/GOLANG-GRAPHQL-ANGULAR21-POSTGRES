package mutations

import (
	"errors"
	"fmt"
	"golang_graphql_postgres/configs"
	"golang_graphql_postgres/internal/dto"
	"golang_graphql_postgres/internal/gql/schema/types"
	"golang_graphql_postgres/internal/models"

	"github.com/graphql-go/graphql"
)

var profileResponseType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ProfileResponse",
	Fields: graphql.Fields{
		"message": &graphql.Field{Type: graphql.String},
		"user":    &graphql.Field{Type: types.UserType},
	},
})

var UpdateProfileField = &graphql.Field{
	Type:        profileResponseType,
	Description: "Update user profile",
	Args: graphql.FieldConfigArgument{
		"id":        &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		"firstname": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"lastname":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"mobile":    &graphql.ArgumentConfig{Type: graphql.String},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		user, ok := params.Context.Value("user").(*types.UserClaims)

		if !ok || user == nil {
			return nil, errors.New("Unauthorized Access, valid bearer token required.")
		}

		userid := params.Args["id"].(int)
		firstname := params.Args["firstname"].(string)
		lastname := params.Args["lastname"].(string)
		mobile, _ := params.Args["mobile"].(string)

		userProfile := dto.UserDTO{
			Firstname: firstname,
			Lastname:  lastname,
			Mobile:    mobile,
		}

		result := configs.DB.Model(&models.User{}).Where("id = ?", userid).Updates(userProfile)

		if result.Error != nil {
			return nil, fmt.Errorf("%s", "Update failed: "+result.Error.Error())
		}

		if result.RowsAffected == 0 {
			return nil, fmt.Errorf("Update failed: User ID not found")
		}

		return map[string]interface{}{
			"message": "You have updated your profile successfully.",
			"user":    userProfile,
		}, nil
	},
}

// ==============REQUEST================
// mutation UpdateProfile(
//     $id: Int!,
//     $firstname: String!,
//     $lastname: String!,
//     $mobile: String!) {
//         updateProfile(id: $id,
//         firstname: $firstname,
//         lastname: $lastname,
//         mobile: $mobile) {
//             message
//             user {
//                 id
//             }
//         }
// }

// VARIABLES
// {
//   "id": 1,
//    "firstname": "Reynaldo",
//    "lastname": "Marquez",
//    "mobile": "+63454545"
// }
