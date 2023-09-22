package services

import (
	"rnd-surajan-gin/dtos"
	"rnd-surajan-gin/infrastructure"
	"rnd-surajan-gin/models"
	"rnd-surajan-gin/pagination"

	"gorm.io/gorm"
)

// Dependency Injection
type TaskService struct {
	db infrastructure.Database
}

func NewTaskService(db infrastructure.Database) TaskService {
	return TaskService{db: db}
}

func (cc TaskService) CreateTask(task models.Task) (data models.Task, err error) {
	return task, cc.db.DB.Create(&task).Error
}

func (cc TaskService) GetAllTasks(pageNo, pageSize string, defaultPageSize int) (data []models.Task, result *gorm.DB) {
	var tasks []models.Task
	// Using Gorm scopes, we can reuse query logic, like the pagination logic => db.Offset().Limit()
	return tasks, cc.db.DB.Scopes(pagination.Paginate(pageNo, pageSize, defaultPageSize)).Find(&tasks)
}

func (cc TaskService) GetTaskById(id string) (data models.Task, result *gorm.DB) {
	var task models.Task
	// The commented code only works when our "id" is automatically generated by GORM and is an auto-incremented number.
	// result := cc.db.DB.First(&task, id)
	return task, cc.db.DB.Where("id = ?", id).First(&task)
}

func (cc TaskService) UpdateTaskById(id string, payload dtos.UpdateTaskRequest) (data models.Task, findErr error, updateErr error) {
	// Get Task by id i.e. Primary Key.
	var task models.Task
	// This code only works when our "id" is automatically generated by GORM and is an auto-incremented number.
	// result := cc.db.DB.First(&task, id)
	result := cc.db.DB.Where("id = ?", id).First(&task)
	// Update "task" model with the given "payload".
	updateResult := cc.db.DB.Model(&task).Updates(payload)
	return task, result.Error, updateResult.Error
}

func (cc TaskService) DeleteTaskById(id string) (result *gorm.DB) {
	// Delete Task by id i.e. Primary Key.
	var task models.Task
	/*
		💡Note: If not used "Unscoped()", then "Delete()" will only do soft-delete because our "tasks" table contain "deleted_at" column.
		When we created our "Task" model, we included "gorm.Model" which will append "deleted_at" column in our "tasks" table.
		See: https://gorm.io/docs/delete.html#Soft-Delete for more info.
	*/
	// This code only works when our "id" is automatically generated by GORM and is an auto-incremented number.
	// result := cc.db.DB.Unscoped().Delete(&task, id)
	return cc.db.DB.Where("id = ?", id).Unscoped().Delete(&task)
}
