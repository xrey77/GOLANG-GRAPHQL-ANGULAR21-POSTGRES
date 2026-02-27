package types

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

// var JwtSecret = []byte("your-secret-key")
var secret = os.Getenv("JWT_SECRET")
var JwtSecret = []byte(secret)

type UserClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// type UserClaims struct {
// 	Username string `json:"username"`
// 	jwt.RegisteredClaims
// }
