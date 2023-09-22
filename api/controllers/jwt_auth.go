package controllers

import (
	"fmt"
	"net/http"
	"rnd-surajan-gin/api/services"
	"rnd-surajan-gin/dtos"
	"rnd-surajan-gin/environment"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type JwtAuthController struct {
	userService    services.UserService
	jwtAuthService services.JwtAuthService
}

func NewJwtController(userService services.UserService, jwtAuthService services.JwtAuthService) JwtAuthController {
	return JwtAuthController{userService: userService, jwtAuthService: jwtAuthService}
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

	// Create a new JWT claims (payload) for the token
	accessClaims := services.JwtClaims{RegisteredClaims: jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
		ID:        fmt.Sprintf("%v", user.ID),
	}}

	// Generate access token
	accessToken, tokenErr := cc.jwtAuthService.GenerateToken(accessClaims, environment.GetJwtSecret())
	if tokenErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "User could not be logged in. Error creating token.",
		})
		return
	}

	// Send found "User" as response.
	ctx.JSON(http.StatusOK, gin.H{
		"data":        "User successfully logged in",
		"accessToken": accessToken,
	})
}
