package mutations

import (
	"fmt"
	"golang_graphql_postgres/configs"
	"golang_graphql_postgres/internal/gql/schema/types"
	middleware "golang_graphql_postgres/internal/middleware"
	"golang_graphql_postgres/internal/models"

	"github.com/graphql-go/graphql"
	"golang.org/x/crypto/bcrypt"
)

var signupResponseType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SignupResponse",
	Fields: graphql.Fields{
		"message": &graphql.Field{Type: graphql.String},
		"user":    &graphql.Field{Type: types.UserType},
	},
})

var SignUpField = &graphql.Field{
	Type:        signupResponseType,
	Description: "Register a new user",
	Args: graphql.FieldConfigArgument{
		"firstname": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"lastname":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"email":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"mobile":    &graphql.ArgumentConfig{Type: graphql.String},
		"username":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"password":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		email := params.Args["email"].(string)
		username := params.Args["username"].(string)
		password := params.Args["password"].(string)
		mobile, _ := params.Args["mobile"].(string)

		hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}

		userEmail, _ := middleware.SearchByEmail(email)
		if userEmail {
			userEmail = false
			return nil, fmt.Errorf("%s", "Email Address is already taken.")

			// return map[string]interface{}{
			// 	"message": "Email Address is already taken.",
			// 	"user":    nil,
			// }, nil
		}

		userName, _ := middleware.SearchByUsername(username)
		if userName {
			userName = false
			// return map[string]interface{}{
			// 	"message": "Username is already taken.",
			// 	"user":    nil,
			// }, nil
			return nil, fmt.Errorf("%s", "Username is already taken.")

		}

		newUser := models.User{
			Firstname: params.Args["firstname"].(string),
			Lastname:  params.Args["lastname"].(string),
			Email:     email,
			Mobile:    mobile,
			Username:  username,
			Password:  string(hashedPass),
			Role_id:   1,
		}

		if err := configs.DB.Create(&newUser).Error; err != nil {
			// return map[string]interface{}{
			// 	"message": "Registration failed: " + err.Error(),
			// 	"user":    nil,
			// }, nil
			return nil, fmt.Errorf("%s", "Registration failed, "+err.Error())

		}

		return map[string]interface{}{
			"message": "User created successfully",
			"user":    newUser,
		}, nil
	},
}

// ========REQUEST========
// mutation SignupUser(
//     $firstname: String!,
//     $lastname: String!,
//     $email: String!,
//     $mobile: String!,
//     $username: String!,
//     $password: String!) {
//     signupUser(
//         firstname: $firstname,
//         lastname: $lastname,
//         email: $email,
//         mobile: $mobile,
//         username: $username,
//         password: $password
//     ) {
//         message
//         user {
//             firstname
//             lastname
//         }
//     }
// }

// =======VARIABLES=========
// {
//     "firstname": "Jigoro",
//     "lastname": "Gragasin",
//     "email": "jigoro@yahoo.com",
//     "mobile": "2342342",
//     "username": "Jigoro",
//     "password": "rey"
// }
