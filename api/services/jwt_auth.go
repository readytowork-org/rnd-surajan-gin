package services

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtAuthService struct {
}

// JWT token has 3 parts separated by ".": "header.payload.signature".
// 👇 This JwtClaims struct is used to generate the "payload" part of the jwt token.
type JwtClaims struct {
	jwt.RegisteredClaims
	// Other claims like username, userid or even email is generally additionally put here.
}

func NewJwtAuthService() JwtAuthService {
	return JwtAuthService{}
}

func (cc JwtAuthService) GenerateToken(claims JwtClaims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create JWT string
	accessToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}
