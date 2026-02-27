package middleware

import (
	"context"
	"fmt"
	types "golang_graphql_postgres/internal/gql/schema/types"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.Next()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.ParseWithClaims(tokenString, &types.UserClaims{}, func(t *jwt.Token) (interface{}, error) {
			return types.JwtSecret, nil
		})

		if err == nil && token.Valid {

			if claims, ok := token.Claims.(*types.UserClaims); ok {
				fmt.Printf("Authenticated user: %s\n", claims.Email)
				ctx := context.WithValue(c.Request.Context(), "user", claims)
				c.Request = c.Request.WithContext(ctx)
			} else {
				fmt.Println("Failed to assert claims")
			}

		}
		c.Next()
	}
}

// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
// 			c.Next()
// 			return
// 		}

// 		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
// 		token, _ := jwt.ParseWithClaims(tokenString, &types.Claims{}, func(t *jwt.Token) (interface{}, error) {
// 			return types.JwtSecret, nil
// 		})

// 		if claims, ok := token.Claims.(*types.Claims); ok && token.Valid {
// 			ctx := context.WithValue(c.Request.Context(), "user", claims)
// 			c.Request = c.Request.WithContext(ctx)
// 		}
// 		c.Next()
// 	}
// }

// import (
// 	"net/http"
// 	utils "src/golang_postgres/util"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// )

// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Get the Authorization header value
// 		authHeader := c.GetHeader("Authorization")

// 		if authHeader == "" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized Access."})
// 			c.Abort() // Stop further processing
// 			return
// 		}

// 		// Check if the header format is "Bearer <token>"
// 		if !strings.HasPrefix(authHeader, "Bearer ") {
// 			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized Access."})
// 			c.Abort() // Stop further processing
// 			return
// 		}

// 		// Extract the token string by trimming the "Bearer " prefix
// 		token := strings.TrimPrefix(authHeader, "Bearer ")

// 		_, err := middleware.VerifyJWT(token)
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Bearer Token."})
// 			c.Abort()
// 			return
// 		}

// 		// store the token or relevant user info in the context for handlers
// 		c.Set("authToken", token)

// 		// Continue to the next handler
// 		c.Next()
// 	}
// }
