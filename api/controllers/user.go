package controllers

import (
	"net/http"
	"rnd-surajan-gin/api/services"
	"rnd-surajan-gin/dtos"
	"rnd-surajan-gin/models"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{userService: userService}
}

func (cc UserController) CreateUser(ctx *gin.Context) {
	var body dtos.CreateUserRequest
	// Validate request body.
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Create User.
	user := models.User{Name: body.Name, Email: body.Email, Password: body.Password}
	data, err := cc.userService.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not create user",
		})
		return
	}
	// Send created "User" as response.
	ctx.JSON(http.StatusOK, gin.H{
		"user": data,
	})
}

func (cc UserController) GetAllUsers(ctx *gin.Context) {
	users, result := cc.userService.GetAllUsers()
	// Error Handling.
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Users could not be found",
		})
		return
	}
	// Send found "Users" as response.
	ctx.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
