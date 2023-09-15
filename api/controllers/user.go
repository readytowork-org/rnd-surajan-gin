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

func (cc UserController) GetUserById(ctx *gin.Context) {
	// Get Id from route parameters.
	id := ctx.Param("id")
	user, result := cc.userService.GetUserById(id)
	// Error Handling.
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "User could not be found",
		})
		return
	}
	// Send found "User" as response.
	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (cc UserController) UpdateUserById(ctx *gin.Context) {
	// Get Id from route parameters.
	id := ctx.Param("id")
	// Validate request body.
	var body dtos.UpdateUserRequest
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	data, findErr, updateErr := cc.userService.UpdateUserById(id, body)
	// Error Handling.
	if findErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "User could not be found",
		})
		return
	}
	if updateErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not update user",
		})
		return
	}
	// Send found and updated "User" as response.
	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (cc UserController) DeleteUserById(ctx *gin.Context) {
	// Get Id from route parameters.
	id := ctx.Param("id")
	result := cc.userService.DeleteUserById(id)
	// Error Handling.
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "User could not be deleted",
		})
		return
	} else if result.RowsAffected < 1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "User could not be deleted because it could not be found.",
		})
		return
	}
	// Send success response.
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User successfully deleted",
	})
}

func (cc UserController) LoginUser(ctx *gin.Context) {
	var body dtos.LoginUserRequest
	// Validate request body.
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Login or Find User
	user, result := cc.userService.LoginUser(body.Email, body.Password)
	// Error Handling.
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "User with those credentials could not be found",
		})
		return
	}
	// Send found "User" as response.
	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
