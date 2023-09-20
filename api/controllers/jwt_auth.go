package controllers

import (
	"net/http"
	"rnd-surajan-gin/api/services"
	"rnd-surajan-gin/dtos"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type JwtAuthController struct {
	userService services.UserService
}

func NewJwtController(userService services.UserService) JwtAuthController {
	return JwtAuthController{userService: userService}
}

func (cc JwtAuthController) LoginWithJwt(ctx *gin.Context) {
	var body dtos.LoginUserRequest
	// Validate request body.
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Login or Find User
	user, result := cc.userService.GetUserFromEmail(body.Email)
	// Error Handling.
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "User with those credentials could not be found",
		})
		return
	}
	// Check if password is correct
	pwErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if pwErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Password does not match",
		})
		return
	}

	// Send found "User" as response.
	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
