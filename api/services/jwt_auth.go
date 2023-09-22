package services

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type JwtAuthService struct {
}

// Detailed explanation 👉: https://www.sohamkamani.com/golang/jwt-authentication/
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

func (cc JwtAuthService) GetTokenFromHeader(ctx *gin.Context) (string, error) {
	// We consider the header to have "Authorization": `Bearer ${token}`
	header := ctx.GetHeader("Authorization")
	if header == "" {
		return "", errors.New("authorization token required in header")
	}
	if !strings.Contains(header, "Bearer") {
		return "", errors.New("bearer is missing in authorization token")
	}
	tokenString := strings.TrimSpace(strings.Replace(header, "Bearer", "", 1))
	return tokenString, nil
}

func (cc JwtAuthService) ParseAndVerifyToken(tokenString string, secret string) (*jwt.Token, error) {
	/* 💡
	In JWT, we extract the jwt-token from Authorization Bearer Header.
	From that jwt-token, we generate signature by using header(string before 1st ".") and claims(payload)(string before 2nd ".") along with secretKey.
	If that generated signature matches the signature from the jwt-token(i.e. the coded string after the last "." in the token),
	then the jwt-token is considered verified.
	*/
	// Validate the jwtToken and extract claims from it using token and secretKey
	// It also checks expiry time
	// The claims is stored in &JwtClaims{} pointer. We can use the claims later (if we want) to regenerate previously stored data like "ID", "ExpiresAt" from it
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, errors.New("invalid token")
	}
	return token, nil
}
