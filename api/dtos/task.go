package dtos

import "rnd-surajan-gin/models"

type CreateTaskRequest struct {
	models.Task
}

type UpdateTaskRequest struct {
	Title string `json:"title" binding:"required"`
}
