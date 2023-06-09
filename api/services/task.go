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

type UpdateTaskRequest struct {
	Title string `json:"title" binding:"required"`
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

func GetAllTasks(ctx *gin.Context) {
	// Get all tasks.
	var tasks []models.Task
	result := database.DB.Find(&tasks)
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
	// Get Task by id i.e. Primary Key.
	var task models.Task
	result := database.DB.First(&task, id)
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

func UpdateTaskById(ctx *gin.Context) {
	// Get Id from route parameters.
	id := ctx.Param("id")
	// Get Task by id i.e. Primary Key.
	var task models.Task
	result := database.DB.First(&task, id)
	// Error Handling.
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Task could not be found",
		})
		return
	}
	// Validate request body.
	var body UpdateTaskRequest
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Update "task" model with the given "body".
	updateResult := database.DB.Model(&task).Updates(body)
	if updateResult.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not update task",
		})
		return
	}
	// Send found and updated "Task" as response.
	ctx.JSON(http.StatusOK, gin.H{
		"data": task,
	})
}

func DeleteTaskById(ctx *gin.Context) {
	// Get Id from route parameters.
	id := ctx.Param("id")
	// Delete Task by id i.e. Primary Key.
	var task models.Task
	/*
		💡Note: If not used "Unscoped()", then "Delete()" will only do soft-delete because our "tasks" table contain "deleted_at" column.
		When we created our "Task" model, we included "gorm.Model" which will append "deleted_at" column in our "tasks" table.
		See: https://gorm.io/docs/delete.html#Soft-Delete for more info.
	*/
	result := database.DB.Unscoped().Delete(&task, id)
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
