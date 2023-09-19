package dtos

import (
	"rnd-surajan-gin/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreateUserRequest struct {
	models.User
}

type UpdateUserRequest struct {
	Name string `json:"name" binding:"required"`
}

// Exclude "Password" field in this struct. This is called "Smart Select Fields" in Gorm.
type GetUserByIdResponse struct {
	gorm.Model
	ID    uuid.UUID `gorm:"primarykey" json:"ID"`
	Name  string    `json:"name" binding:"required"`
	Email string    `json:"email" binding:"required,email"`
	// As "Tasks" are originally from "models.User" and not "GetUserByIdResponse", we have to explicitly set "foreignKey:UserId".
	Tasks []models.Task `gorm:"foreignKey:UserId" json:"tasks"`
}
