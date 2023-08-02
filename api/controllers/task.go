package controllers

import (
	"net/http"
	"rnd-surajan-gin/api/services"
	"rnd-surajan-gin/models"

	"github.com/gin-gonic/gin"
)

type CreateTaskRequest struct {
	models.Task
}

func CreateTask(ctx *gin.Context) {
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
	result := services.CreateTaskService(task)
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

func GetAllTasks(ctx *gin.Context) {
	tasks, result := services.GetAllTasks()
	// Error Handling.
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

func GetTaskById(ctx *gin.Context) {
	// Get Id from route parameters.
	id := ctx.Param("id")
	task, result := services.GetTaskById(id)
	// Error Handling.
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Task could not be found",
		})
		return
	}
	// Send found "Task" as response.
	ctx.JSON(http.StatusOK, gin.H{
		"data": task,
	})
}

func DeleteTaskById(ctx *gin.Context) {
	// Get Id from route parameters.
	id := ctx.Param("id")
	result := services.DeleteTaskByIdService(id)
	// Error Handling.
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Task could not be deleted",
		})
		return
	} else if result.RowsAffected < 1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Task could not be deleted because it could not be found.",
		})
		return
	}
	// Send success response.
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Task successfully deleted",
	})
}
