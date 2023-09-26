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
		// Get
		routes.GET("", cc.taskController.GetAllTasks)
		routes.GET("/byUser", cc.taskController.GetTaskByUserIdAndStatus)
		routes.GET("/report/byUser", cc.taskController.GetTaskReportByUserId)
		routes.GET("/:id", cc.taskController.GetTaskById)
		// Post
		routes.POST("", cc.taskController.CreateTask)
		routes.POST("/byFormdata", cc.taskController.CreateTaskByFormdata)
		// Patch
		routes.PATCH("/:id", cc.taskController.UpdateTaskById)
		routes.PATCH("/status/:id", cc.taskController.UpdateTaskStatus)
		// Delete
		routes.DELETE("/:id", cc.taskController.DeleteTaskById)
	}
}
