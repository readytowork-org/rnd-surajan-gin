package routes

import (
	"rnd-surajan-gin/api/controllers"
	"rnd-surajan-gin/api/infrastructure"
)

type TaskRoutes struct {
	router         infrastructure.Router
	taskController controllers.TaskController
}

func NewTaskRoutes(router infrastructure.Router, taskController controllers.TaskController) TaskRoutes {
	return TaskRoutes{router: router, taskController: taskController}
}

func (cc TaskRoutes) TasksRouteSetup() {
	// Task Routes
	cc.router.Gin.GET("/tasks", cc.taskController.GetAllTasks)
	cc.router.Gin.GET("/tasks/:id", cc.taskController.GetTaskById)
	cc.router.Gin.POST("/tasks", cc.taskController.CreateTask)
	cc.router.Gin.PATCH("/tasks/:id", cc.taskController.UpdateTaskById)
	cc.router.Gin.DELETE("/tasks/:id", cc.taskController.DeleteTaskById)
}
