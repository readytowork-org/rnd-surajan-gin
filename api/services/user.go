package services

import (
	"rnd-surajan-gin/dtos"
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
	// Preload all "Tasks" associated with each "User" and send them along with the found users.
	return users, cc.db.DB.Model(&models.User{}).Preload("Tasks").Find(&users)
}

func (cc UserService) GetUserById(id string) (data models.User, result *gorm.DB) {
	var user models.User
	// As we search for "uuid" rather than "Id" that is auto-generated by gorm, we have to use ".Where()".
	// Using "Preload()" with Where.
	return user, cc.db.DB.Model(&models.User{}).Preload("Tasks").Where("id = ?", id).First(&user)
}

func (cc UserService) UpdateUserById(id string, payload dtos.UpdateUserRequest) (data models.User, findErr error, updateErr error) {
	// Get User by id i.e. Primary Key.
	var user models.User
	result := cc.db.DB.Where("id = ?", id).First(&user)
	updateResult := cc.db.DB.Model(&user).Updates(payload)
	return user, result.Error, updateResult.Error
}

func (cc UserService) DeleteUserById(id string) (result *gorm.DB) {
	// Delete User by id i.e. Primary Key.
	var user models.User
	return cc.db.DB.Where("id = ?", id).Unscoped().Delete(&user)
}

func (cc UserService) LoginUser(email string, password string) (data models.User, result *gorm.DB) {
	var user models.User
	// We used "BINARY ?" because, only using "?" will match in case-insensitive way i.e. both "password" and "Password" will match.
	return user, cc.db.DB.Where("email = BINARY ? AND password = BINARY ?", email, password).First(&user)
}
