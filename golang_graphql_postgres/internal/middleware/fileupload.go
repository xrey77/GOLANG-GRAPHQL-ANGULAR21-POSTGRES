package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GraphQLMultipartMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !strings.HasPrefix(c.ContentType(), "multipart/form-data") {
			c.Next()
			return
		}

		// 1. Parse the multipart form
		form, err := c.MultipartForm()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
			return
		}

		// 2. Extract 'operations' (contains query and variables)
		operationsStr := c.PostForm("operations")
		var operations map[string]interface{}
		if err := json.Unmarshal([]byte(operationsStr), &operations); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid operations JSON"})
			return
		}

		// 3. Extract 'map' (maps form field names to variable paths)
		mapStr := c.PostForm("map")
		var fileMap map[string][]string
		if err := json.Unmarshal([]byte(mapStr), &fileMap); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid map JSON"})
			return
		}

		// 4. Map files back to the "variables" map
		for fileKey, paths := range fileMap {
			files := form.File[fileKey]
			if len(files) == 0 {
				continue
			}

			// For each path (e.g., "variables.file"), inject the file header
			for _, path := range paths {
				injectFile(operations, path, files[0])
			}
		}

		// 5. Update the request body for the GraphQL handler
		newBody, _ := json.Marshal(operations)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(newBody))
		c.Request.Header.Set("Content-Type", "application/json")
		c.Request.ContentLength = int64(len(newBody))

		c.Next()
	}
}

// Helper to inject file into nested map based on dot-notation path
func injectFile(operations map[string]interface{}, path string, file interface{}) {
	parts := strings.Split(path, ".")
	var current interface{} = operations

	for i, part := range parts {
		if i == len(parts)-1 {
			if m, ok := current.(map[string]interface{}); ok {
				m[part] = file
			}
			return
		}

		if m, ok := current.(map[string]interface{}); ok {
			current = m[part]
		} else {
			return
		}
	}
}
