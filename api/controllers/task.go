package controllers

import (
	"fmt"
	"math"
	"net/http"
	"rnd-surajan-gin/api/services"
	"rnd-surajan-gin/constants"
	"rnd-surajan-gin/dtos"
	"rnd-surajan-gin/models"
	"rnd-surajan-gin/pagination"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Dependency Injection
type TaskController struct {
	taskService services.TaskService
}

//	func NewTaskController() TaskController {
//		return TaskController{taskService: services.NewTaskService()}
//	}
func NewTaskController(taskService services.TaskService) TaskController {
	return TaskController{taskService: taskService}
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
	task := models.Task{Title: body.Title, UserId: body.UserId}
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
	// Get Query Params if any available, otherwise set them to defaults
	defaultPageNo, defaultPageSize := pagination.DefaultPageVariables()
	pageNo := ctx.DefaultQuery("page", strconv.Itoa(defaultPageNo))
	pageSize := ctx.DefaultQuery("pageSize", strconv.Itoa(defaultPageSize))
	// For counting total data in the database
	var count int64
	// Task Service
	tasks, totalDataResult, result := cc.taskService.GetAllTasks(pageNo, pageSize, defaultPageSize, &count)
	// Error Handling.
	if result.Error != nil && totalDataResult.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Tasks could not be found",
		})
		return
	}
	// Convert pageSize from query params to int
	pg, _ := strconv.Atoi(pageSize)
	// Send found "Tasks" as response.
	ctx.JSON(http.StatusOK, gin.H{
		"data":       tasks,
		"page":       pageNo,
		"pageSize":   pageSize,
		"totalData":  count,
		"totalPages": math.Ceil(float64(count) / float64(pg)),
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

func (cc TaskController) UpdateTaskStatus(ctx *gin.Context) {
	// Get UserId from JWT (set by jwt middleware using ctx.Set())
	userId := fmt.Sprintf("%v", ctx.MustGet(constants.UserId))
	// Get Id from route parameters.
	id := ctx.Param("id")
	// Validate request body.
	var body dtos.UpdateTaskStatus
	// This validates payload's key
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// This validates payload's value i.e. status values if they match the 3 statuses provided by us.
	if !body.IsValidStatus() {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid status",
		})
		return
	}
	// Only that user who owns this task can update it's status
	data, findErr, updateErr := cc.taskService.UpdateTaskStatus(id, userId, string(body.Status))
	// Error Handling.
	if findErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Task could not be found",
		})
		return
	}
	if updateErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not update task status",
		})
		return
	}
	// Send found and updated "Task" as response.
	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (cc TaskController) GetTaskByUserIdAndStatus(ctx *gin.Context) {
	userId := ctx.Param("id")
	status := ctx.DefaultQuery("status", "")
	tasks, result := cc.taskService.GetTaskByUserIdAndStatus(userId, status)
	// Error Handling.
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Tasks could not be found",
		})
		return
	}
	// Send found "Task" as response.
	ctx.JSON(http.StatusOK, gin.H{
		"data": tasks,
	})
}

func (cc TaskController) GetTaskReportByUserId(ctx *gin.Context) {
	userId := ctx.Param("id")
	var todoCount, inProgressCount, doneCount int64
	todoResult, inProgressResult, doneResult := cc.taskService.GetTaskReportByUserId(userId, &todoCount, &inProgressCount, &doneCount)
	// Error Handling.
	if todoResult.Error != nil || inProgressResult.Error != nil || doneResult.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Tasks could not be found",
		})
		return
	}
	// Send found "Task" as response.
	ctx.JSON(http.StatusOK, gin.H{
		"to_do":       todoCount,
		"in_progress": inProgressCount,
		"done":        doneCount,
	})
}
