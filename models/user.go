package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
	// Create a "[]byte" type from user's password because that is the type "bcrypt" needs to generate hashed passwords.
	pwInBytes := []byte(user.Password)
	encryptedPw, pwErr := bcrypt.GenerateFromPassword(pwInBytes, bcrypt.DefaultCost)
	if err != nil || pwErr != nil {
		return err
	}
	user.ID = id
	user.Password = string(encryptedPw)
	return
}
