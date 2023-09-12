package routes

import (
	"rnd-surajan-gin/api/controllers"
	"rnd-surajan-gin/infrastructure"
)

type UserRoutes struct {
	router         infrastructure.Router
	userController controllers.UserController
}

func NewUserRoutes(router infrastructure.Router, userController controllers.UserController) UserRoutes {
	return UserRoutes{router: router, userController: userController}
}

func (cc UserRoutes) Setup() {
	// User Routes
	cc.router.Gin.POST("/users", cc.userController.CreateUser)
}
