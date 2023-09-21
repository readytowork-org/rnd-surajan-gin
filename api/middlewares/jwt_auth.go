package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type JwtAuthMiddleware struct {
}

func NewJwtAuthMiddleware() JwtAuthMiddleware {
	return JwtAuthMiddleware{}
}

func (m JwtAuthMiddleware) HandleJwt(ctx *gin.Context) {
	fmt.Println("MIDDLEWARE ON USER ROUTES")
	ctx.Next()
}
