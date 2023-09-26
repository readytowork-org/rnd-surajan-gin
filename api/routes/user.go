package routes

import (
	"rnd-surajan-gin/api/controllers"
	"rnd-surajan-gin/api/middlewares"
	"rnd-surajan-gin/infrastructure"
)

type UserRoutes struct {
	router            infrastructure.Router
	userController    controllers.UserController
	jwtAuthMiddleware middlewares.JwtAuthMiddleware
}

func NewUserRoutes(router infrastructure.Router, userController controllers.UserController, jwtAuthMiddleware middlewares.JwtAuthMiddleware) UserRoutes {
	return UserRoutes{router: router, userController: userController, jwtAuthMiddleware: jwtAuthMiddleware}
}

func (cc UserRoutes) Setup() {
	// User Routes
	routes := cc.router.Gin.Group("/users", cc.jwtAuthMiddleware.HandleJwt)
	{
		routes.GET("", cc.userController.GetAllUsers)
		routes.GET("/withPw", cc.userController.GetAllUsersWithPw)
		routes.GET("/withPw/:id", cc.userController.GetUserByIdWithPw)
		routes.GET("/:id", cc.userController.GetUserById)
		routes.POST("", cc.userController.CreateUser)
		routes.PATCH("/:id", cc.userController.UpdateUserById)
		routes.DELETE("/:id", cc.userController.DeleteUserById)
	}
	// Register User
	cc.router.Gin.POST("/register", cc.userController.CreateUser)
}
