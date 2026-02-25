package mutations

import (
	"github.com/graphql-go/graphql"
)

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"signupUser":      SignUpField,
		"signinUser":      SignInField,
		"updateProfile":   UpdateProfileField,
		"updatePassword":  UpdatePasswordField,
		"mfaActivation":   MfaActivationField,
		"otpVerification": OtpVerificationField,
		"uploadPicture":   UploadPictureField,
	},
})
