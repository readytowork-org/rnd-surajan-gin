package routes

import (
	"rnd-surajan-gin/api/controllers"
	"rnd-surajan-gin/infrastructure"
)

type JwtAuthRoutes struct {
	router         infrastructure.Router
	userController controllers.UserController
}

func NewJwtAuthRoutes(router infrastructure.Router, userController controllers.UserController) JwtAuthRoutes {
	return JwtAuthRoutes{router: router, userController: userController}
}

func (cc JwtAuthRoutes) Setup() {
	// Jwt Auth Routes
	routes := cc.router.Gin.Group("/login")
	{
		routes.POST("", cc.userController.LoginUser)
	}
}
