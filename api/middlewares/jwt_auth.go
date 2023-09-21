package middlewares

import (
	"fmt"
	"net/http"
	"rnd-surajan-gin/api/services"
	"rnd-surajan-gin/environment"

	"github.com/gin-gonic/gin"
)

type JwtAuthMiddleware struct {
	jwtAuthService services.JwtAuthService
}

func NewJwtAuthMiddleware(jwtAuthService services.JwtAuthService) JwtAuthMiddleware {
	return JwtAuthMiddleware{jwtAuthService: jwtAuthService}
}

func (m JwtAuthMiddleware) HandleJwt(ctx *gin.Context) {
	// Get Access Token from Authorization Header
	token, err := m.jwtAuthService.GetTokenFromHeader(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		// If not aborted, any other ctx.JSON responses will also get thrown along with {"error": "Unauthorized"}
		ctx.Abort()
		return
	}
	// Verify the extracted "token".
	parsedToken, parseErr := m.jwtAuthService.ParseAndVerifyToken(token, environment.GetJwtSecret())
	if parseErr != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": parseErr.Error(),
		})
		ctx.Abort()
		return
	}
	if parsedToken != nil {
		fmt.Println("TOKEN VERIFIED:", parsedToken)
	}
	ctx.Next()
}
