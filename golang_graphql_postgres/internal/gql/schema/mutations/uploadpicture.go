package mutations

import (
	"fmt"
	"golang_graphql_postgres/configs"
	"golang_graphql_postgres/internal/gql/schema/types"
	middlware "golang_graphql_postgres/internal/middleware"
	"golang_graphql_postgres/internal/models"
	"log"
	"mime/multipart"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

var uploadResponseType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UploadResponse",
	Fields: graphql.Fields{
		"message": &graphql.Field{Type: graphql.String},
		"user":    &graphql.Field{Type: types.UserType},
	},
})

var UploadPictureField = &graphql.Field{
	Type:        uploadResponseType,
	Description: "Upload profile picture",
	Args: graphql.FieldConfigArgument{
		"id":          &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		"userpicture": &graphql.ArgumentConfig{Type: types.UploadScalar},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		userid := params.Args["id"].(int)
		file, ok := params.Args["userpicture"].(*multipart.FileHeader)
		if !ok {
			return nil, fmt.Errorf("invalid image file")
		}
		gc, ok := params.Context.Value("GinContext").(*gin.Context)
		if !ok {
			return nil, fmt.Errorf("could not find gin context")
		}

		userData, err := middlware.GetUserID(userid)
		if err != nil {
			return nil, fmt.Errorf("User ID not found.")
		}

		filename := filepath.Base(file.Filename)
		ext := filepath.Ext(filename)
		// newfile := "00" + string(rune(userid)) + ext
		newfile := fmt.Sprintf("00%d%s", userid, ext)

		log.Println("NEW FILE................", newfile)
		result := configs.DB.Model(&models.User{}).
			Where("id = ?", userid).
			Updates(map[string]interface{}{
				"userpicture": newfile,
			})

		if result.Error != nil {
			return nil, fmt.Errorf("%s", "Update failed: "+result.Error.Error())
		}

		dst := filepath.Join("./templates/assets/users", newfile)
		if err := gc.SaveUploadedFile(file, dst); err != nil {
			return nil, err
		}

		return map[string]interface{}{
			"message": "You have changed your profile picture successfully.",
			"user":    userData,
		}, nil
	},
}
