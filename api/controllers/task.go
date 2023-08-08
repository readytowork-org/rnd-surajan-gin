package controllers

import (
	"net/http"
	"rnd-surajan-gin/api/dtos"
	"rnd-surajan-gin/api/services"
	"rnd-surajan-gin/models"

	"github.com/gin-gonic/gin"
)

// Dependency Injection
type TaskController struct {
	taskService services.TaskService
}

func NewTaskController() TaskController {
	return TaskController{taskService: services.NewTaskService()}
}

func (cc TaskController) CreateTask(ctx *gin.Context) {
	var body dtos.CreateTaskRequest
	// Validate request body.
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Create Task.
	task := models.Task{Title: body.Title}
	data, err := cc.taskService.CreateTask(task)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not create task",
		})
		return
	}
	// Send created "Task" as response.
	ctx.JSON(http.StatusOK, gin.H{
		"task": data,
	})
}

func (cc TaskController) GetAllTasks(ctx *gin.Context) {
	tasks, result := cc.taskService.GetAllTasks()
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

func (cc TaskController) GetTaskById(ctx *gin.Context) {
	// Get Id from route parameters.
	id := ctx.Param("id")
	task, result := cc.taskService.GetTaskById(id)
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

func (cc TaskController) UpdateTaskById(ctx *gin.Context) {
	// Get Id from route parameters.
	id := ctx.Param("id")
	// Validate request body.
	var body dtos.UpdateTaskRequest
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	data, findErr, updateErr := cc.taskService.UpdateTaskById(id, body)
	// Error Handling.
	if findErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Task could not be found",
		})
		return
	}
	if updateErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not update task",
		})
		return
	}
	// Send found and updated "Task" as response.
	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (cc TaskController) DeleteTaskById(ctx *gin.Context) {
	// Get Id from route parameters.
	id := ctx.Param("id")
	result := cc.taskService.DeleteTaskById(id)
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
