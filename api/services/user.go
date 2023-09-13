package services

import (
	"rnd-surajan-gin/infrastructure"
	"rnd-surajan-gin/models"

	"gorm.io/gorm"
)

type UserService struct {
	db infrastructure.Database
}

func NewUserService(db infrastructure.Database) UserService {
	return UserService{db: db}
}

func (cc UserService) CreateUser(user models.User) (data models.User, err error) {
	return user, cc.db.DB.Create(&user).Error
}

func (cc UserService) GetAllUsers() (data []models.User, result *gorm.DB) {
	var users []models.User
	return users, cc.db.DB.Find(&users)
}
