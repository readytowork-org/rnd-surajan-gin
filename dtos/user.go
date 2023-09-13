package dtos

import "rnd-surajan-gin/models"

type CreateUserRequest struct {
	models.User
}

type UpdateUserRequest struct {
	Name string `json:"name" binding:"required"`
}
