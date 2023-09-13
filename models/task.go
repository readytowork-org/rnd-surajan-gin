package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID     uuid.UUID `gorm:"primarykey" json:"ID"`
	Title  string    `json:"title" binding:"required"`
	UserId uuid.UUID `binding:"required"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (task *Task) BeforeCreate(db *gorm.DB) (err error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	task.ID = id
	return
}
