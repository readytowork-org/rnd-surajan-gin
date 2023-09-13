package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"primarykey" json:"ID"`
	Name     string    `json:"name" binding:"required"`
	Email    string    `json:"email" binding:"required,email"`
	Password string    `json:"password" binding:"required"`
	Tasks    []Task    `json:"tasks"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (user *User) BeforeCreate(db *gorm.DB) (err error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	user.ID = id
	return
}
