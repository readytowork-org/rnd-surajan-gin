package routes

import (
	"rnd-surajan-gin/api/controllers"

	"github.com/gin-gonic/gin"
)

type TaskRoutes struct {
	taskController controllers.TaskController
}

func NewTaskRoutes(taskController controllers.TaskController) TaskRoutes {
	return TaskRoutes{taskController: taskController}
}

func (cc TaskRoutes) TasksRouteSetup(r *gin.Engine) {
	// Task Routes
	r.GET("/tasks", cc.taskController.GetAllTasks)
	r.GET("/tasks/:id", cc.taskController.GetTaskById)
	r.POST("/tasks", cc.taskController.CreateTask)
	r.PATCH("/tasks/:id", cc.taskController.UpdateTaskById)
	r.DELETE("/tasks/:id", cc.taskController.DeleteTaskById)
}
