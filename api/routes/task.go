package routes

import (
	"rnd-surajan-gin/api/controllers"
	"rnd-surajan-gin/api/middlewares"
	"rnd-surajan-gin/infrastructure"
)

type TaskRoutes struct {
	router            infrastructure.Router
	taskController    controllers.TaskController
	jwtAuthMiddleware middlewares.JwtAuthMiddleware
}

func NewTaskRoutes(router infrastructure.Router, taskController controllers.TaskController, jwtAuthMiddleware middlewares.JwtAuthMiddleware) TaskRoutes {
	return TaskRoutes{router: router, taskController: taskController, jwtAuthMiddleware: jwtAuthMiddleware}
}

func (cc TaskRoutes) Setup() {
	// Task Routes
	routes := cc.router.Gin.Group("/tasks", cc.jwtAuthMiddleware.HandleJwt)
	{
		routes.GET("", cc.taskController.GetAllTasks)
		routes.GET("/:id", cc.taskController.GetTaskById)
		routes.POST("", cc.taskController.CreateTask)
		routes.PATCH("/:id", cc.taskController.UpdateTaskById)
		routes.DELETE("/:id", cc.taskController.DeleteTaskById)
	}
}
