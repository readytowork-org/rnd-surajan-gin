package routes

import (
	"rnd-surajan-gin/infrastructure"

	"github.com/gin-gonic/gin"
)

type TestRoutes struct {
	router infrastructure.Router
}

func NewTestRoutes(router infrastructure.Router) TestRoutes {
	return TestRoutes{router: router}
}

func (cc TestRoutes) Setup() {
	cc.router.Gin.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Test",
		})
	})
}
