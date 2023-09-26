package dtos

import "rnd-surajan-gin/models"

type CreateTaskRequest struct {
	Title string `json:"title" binding:"required"`
}

type CreateTaskByFormDataRequest struct {
	Title string `form:"title" binding:"required"`
}
type UpdateTaskRequest struct {
	Title string `json:"title" binding:"required"`
}

type UpdateTaskStatus struct {
	Status models.TaskStatus `json:"status" binding:"required"`
}

var AllowedStatus = []models.TaskStatus{models.TaskToDo, models.TaskInProgress, models.TaskDone}

// Validate UpdateTaskStatus payload
func (task *UpdateTaskStatus) IsValidStatus() bool {
	for _, allowed := range AllowedStatus {
		if task.Status == allowed {
			return true
		}
	}
	return false
}
