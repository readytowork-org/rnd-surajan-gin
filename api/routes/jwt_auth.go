package routes

import (
	"rnd-surajan-gin/api/controllers"
	"rnd-surajan-gin/infrastructure"
)

type JwtAuthRoutes struct {
	router            infrastructure.Router
	jwtAuthController controllers.JwtAuthController
}

func NewJwtAuthRoutes(router infrastructure.Router, jwtAuthController controllers.JwtAuthController) JwtAuthRoutes {
	return JwtAuthRoutes{router: router, jwtAuthController: jwtAuthController}
}

func (cc JwtAuthRoutes) Setup() {
	// Jwt Auth Routes
	routes := cc.router.Gin.Group("/login")
	{
		routes.POST("", cc.jwtAuthController.LoginWithJwt)
	}
}
