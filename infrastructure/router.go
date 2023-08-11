package infrastructure

import "github.com/gin-gonic/gin"

type Router struct {
	Gin *gin.Engine
}

func NewRouter() Router {
	// Gin Server
	r := gin.Default()
	return Router{
		Gin: r,
	}
}
