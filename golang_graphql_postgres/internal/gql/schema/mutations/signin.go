package mutations

import (
	"golang_graphql_postgres/internal/gql/schema/types"
	middleware "golang_graphql_postgres/internal/middleware"
	"golang_graphql_postgres/internal/models"

	"github.com/graphql-go/graphql"
	"golang.org/x/crypto/bcrypt"
)

var signinResponseType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SigninResponse",
	Fields: graphql.Fields{
		"token":    &graphql.Field{Type: graphql.String},
		"rolename": &graphql.Field{Type: graphql.String},
		"message":  &graphql.Field{Type: graphql.String},
		"user":     &graphql.Field{Type: types.UserType},
	},
})

var SignInField = &graphql.Field{
	Type:        signinResponseType,
	Description: "Login an existing user",
	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		username := params.Args["username"].(string)
		plainPwd := params.Args["password"].(string)

		user, err := middleware.GetUserInfo(username)
		if err != nil {
			return map[string]interface{}{
				"token":    nil,
				"rolename": nil,
				"message":  err.Error(),
				"user":     nil,
			}, nil

		}

		hashPwd := user.Password
		err = bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(plainPwd))
		if err != nil {
			return map[string]interface{}{
				"token":    nil,
				"rolename": nil,
				"message":  "Invalid password, please try again.",
				"user":     nil,
			}, nil
		}

		xtoken, err := middleware.GenerateToken(user.Email)
		token := ""
		if err == nil {
			token = xtoken
		}

		role, err := middleware.GetRolName(user.Role_id)
		roleName := ""
		if err == nil && role != nil {
			roleName = role.Name
		}

		userData := models.User{
			ID:          user.ID,
			Firstname:   user.Firstname,
			Lastname:    user.Lastname,
			Email:       user.Email,
			Mobile:      user.Mobile,
			Username:    user.Username,
			Userpicture: user.Userpicture,
			Isactivated: user.Isactivated,
			Isblocked:   user.Isblocked,
			Mailtoken:   user.Mailtoken,
			Qrcodeurl:   user.Qrcodeurl,
		}

		return map[string]interface{}{
			"token":    token,
			"rolename": roleName,
			"message":  "You have logged-in successfully",
			"user":     userData,
		}, nil
	},
}

// ================REQUEST==================
// mutation SigninUser($username: String!, $password: String!) {
//     signinUser(username: $username, password: $password) {
//         token
//         rolename
//         message
//         user {
//             id
//             firstname
//             lastname
//             email
//             mobile
//             roles
//             userpicture
//             qrcodeurl
//         }
//     }
// }

// ============VARIABLES==================
// {
//     "username": "Rey",
//     "password": "xrey"
// }
