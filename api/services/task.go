package services

import (
	"net/http"
	"rnd-surajan-gin/database"
	"rnd-surajan-gin/models"

	"github.com/gin-gonic/gin"
)

type CreateTaskRequest struct {
	models.Task
}

func CreateTaskService(ctx *gin.Context) {
	var body CreateTaskRequest
	// Validate request body.
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Create Task.
	task := models.Task{Title: body.Title}
	result := database.DB.Create(&task)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not create task",
		})
		return
	}
	// Send created "Task" as response.
	ctx.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

func GetAllTasksService(ctx *gin.Context) {
	// Get all tasks
	var tasks []models.Task
	result := database.DB.Find(&tasks)
	// Error Handling
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Tasks could not be found",
		})
		return
	}
	// Send found "Tasks" as response.
	ctx.JSON(http.StatusOK, gin.H{
		"data": tasks,
	})
}
