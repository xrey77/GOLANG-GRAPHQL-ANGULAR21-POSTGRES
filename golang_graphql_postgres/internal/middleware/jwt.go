package middleware

import (

	// "go/types"
	types "golang_graphql_postgres/internal/gql/schema/types"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// var jwtKey = []byte(secret)

func GenerateToken(email string) (string, error) {
	claims := &types.UserClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(types.JwtSecret)
}

// type Claims struct {
// 	Username string `json:"username"`
// 	jwt.RegisteredClaims
// }

// func GenerateToken(username string) (string, error) {

// 	// it will expire in 8 hours
// 	expirationTime := time.Now().Add(8 * time.Hour)

// 	claims := &types.Claims{
// 		Email: email,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(expirationTime),
// 			IssuedAt:  jwt.NewNumericDate(time.Now()),
// 			NotBefore: jwt.NewNumericDate(time.Now()),
// 			Issuer:    "WORLD BANK",
// 			Subject:   "user_authentication",
// 			ID:        "1",
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	tokenString, err := token.SignedString(types.JwtSecret)

// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil
// }

// func VerifyJWT(tokenString string) (*types.Claims, error) {
// 	claims := &types.Claims{}

// 	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}

// 		return types.JwtSecret, nil
// 	})

// 	if err != nil {
// 		if err == jwt.ErrSignatureInvalid {
// 			return nil, fmt.Errorf("invalid token signature")
// 		}
// 		return nil, err
// 	}

// 	if !token.Valid {
// 		return nil, fmt.Errorf("invalid token")
// 	}

// 	return claims, nil
// }
