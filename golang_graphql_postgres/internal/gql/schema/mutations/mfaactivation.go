package mutations

import (
	"encoding/base64"
	"errors"
	"fmt"
	"golang_graphql_postgres/configs"
	"golang_graphql_postgres/internal/dto"
	"golang_graphql_postgres/internal/gql/schema/types"
	middleware "golang_graphql_postgres/internal/middleware"
	"golang_graphql_postgres/internal/models"

	"github.com/graphql-go/graphql"
	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
)

var activationResponseType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ActivationResponse",
	Fields: graphql.Fields{
		"message": &graphql.Field{Type: graphql.String},
		"user":    &graphql.Field{Type: types.UserType},
	},
})

var MfaActivationField = &graphql.Field{
	Type:        activationResponseType,
	Description: "Activate/De-Activate Multi-Factor Authenticator",
	Args: graphql.FieldConfigArgument{
		"id":               &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		"twofactorenabled": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Boolean)},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		user, ok := params.Context.Value("user").(*types.UserClaims)

		if !ok || user == nil {
			return nil, errors.New("Unauthorized Access, valid bearer token required.")
		}

		userid := params.Args["id"].(int)
		twofactorenabled, _ := params.Args["twofactorenabled"].(bool)

		if twofactorenabled {

			userData, err := middleware.GetUserID(userid)
			if err != nil {
				return nil, fmt.Errorf("%s", "User ID not found."+err.Error())
			}

			key, err := totp.Generate(totp.GenerateOpts{
				Issuer:      "APPLE INC.",
				AccountName: userData.Email,
			})

			if err != nil {
				return nil, fmt.Errorf("Failed to generate TOTP secret.")
			}
			secret := key.Secret()
			qrCodeURL := key.URL()

			pngBytes, err := qrcode.Encode(qrCodeURL, qrcode.Medium, 256)
			if err != nil {
				return nil, fmt.Errorf("Failed to generate QRCODE.")

			}
			base64Encoded := "data:image/png;base64," + string(base64.StdEncoding.EncodeToString(pngBytes))

			profileDto := dto.UserDTO{
				Secret:    &secret,
				Qrcodeurl: &base64Encoded,
			}

			result := configs.DB.Model(&models.User{}).Where("id = ?", userid).Updates(profileDto)

			if result.Error != nil {
				return nil, fmt.Errorf("%s", "Error Found, "+result.Error.Error())
			}

			var count int64
			configs.DB.Model(&models.User{}).Where("id = ?", userid).Count(&count)
			if count == 0 {
				return nil, fmt.Errorf("User ID not found")
			}
			userData.Qrcodeurl = &base64Encoded

			return map[string]interface{}{
				"message": "Multi-Factor Authenticator has been enabled successfully.",
				"user":    profileDto,
			}, nil

		} else {

			result := configs.DB.Model(&models.User{}).
				Where("id = ?", userid).
				Updates(map[string]interface{}{
					"secret":    nil,
					"qrcodeurl": nil,
				})

			if result.Error != nil {
				return nil, fmt.Errorf("%s", "Error Found, "+result.Error.Error())
			}

			if result.RowsAffected == 0 {
				return nil, fmt.Errorf("User ID not found.")
			}

			return map[string]interface{}{
				"message": "Multi-Factor has been disabled successfully.",
				"user":    nil,
			}, nil
		}
	},
}

// =============REQUEST================
// mutation MfaActivation($id: Int!, $twofactorenabled: Boolean!) {
//     mfaActivation(id: $id, twofactorenabled: $twofactorenabled) {
//         message
//         user {
//             qrcodeurl
//         }
//     }
// }

// ================VARIABLES==========
// {
//     "id": 1,
//     "twofactorenabled": false
// }
