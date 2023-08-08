package routes

import (
	"rnd-surajan-gin/api/controllers"

	"github.com/gin-gonic/gin"
)

func TasksRouteSetup(r *gin.Engine) {
	// Task Routes
	r.GET("/tasks", controllers.NewTaskController().GetAllTasks)
	r.GET("/tasks/:id", controllers.NewTaskController().GetTaskById)
	r.POST("/tasks", controllers.NewTaskController().CreateTask)
	r.PATCH("/tasks/:id", controllers.NewTaskController().UpdateTaskById)
	r.DELETE("/tasks/:id", controllers.NewTaskController().DeleteTaskById)
}
