package mutations

import (
	"fmt"
	"golang_graphql_postgres/internal/gql/schema/types"
	utils "golang_graphql_postgres/internal/middlware"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/pquerna/otp/totp"
)

var otpVerificationResponseType = graphql.NewObject(graphql.ObjectConfig{
	Name: "OtpVerificationResponse",
	Fields: graphql.Fields{
		"message": &graphql.Field{Type: graphql.String},
		"user":    &graphql.Field{Type: types.UserType},
	},
})

var OtpVerificationField = &graphql.Field{
	Type:        otpVerificationResponseType,
	Description: "One Time Password Verification",
	Args: graphql.FieldConfigArgument{
		"id":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		"otp": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		userid := params.Args["id"].(int)
		otp := params.Args["otp"].(string)

		user, err := utils.GetUserID(userid)
		if err != nil {
			return nil, fmt.Errorf("User ID not found: %w", err)
		}

		log.Println("SECRET......." + *user.Secret)
		log.Println("OTP..........." + otp)
		if user.Secret != nil {

			valid := totp.Validate(otp, *user.Secret)
			if valid {
				return map[string]interface{}{
					"message": "OTP code was verified successfully.",
					"user":    user,
				}, nil

			}
			return nil, fmt.Errorf("Invalid OTP code, please try again.")

		}
		return nil, fmt.Errorf("Multi-Factor Authenticator is not yet activated.")
	},
}

// REQUEST
// mutation OtpVerification($id: Int!, $otp: String!) {
//     otpVerification(id: $id, otp: $otp) {
//         message
//         user {
//             username
//         }
//     }
// }

// VARIABLES
// {
//     "id": 1,
//     "otp": "234234"
// }
