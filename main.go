package main

import (
	"net/http"
	"rnd-surajan-gin/api/controllers"
	"rnd-surajan-gin/database"
	"rnd-surajan-gin/infrastucture"
	"rnd-surajan-gin/models"

	"github.com/gin-gonic/gin"
)

func init() {
	// Initialize Env
	infrastucture.EnvInit()
	// DB connection
	database.ConnectDB()
	// Migrate the schema
	database.DB.AutoMigrate(&models.Task{})
}

func main() {
	// Gin Server
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Using Query Parameters
	r.GET("/welcome-home", func(ctx *gin.Context) {
		// This puts "Guest" in fName if "?firstname=" does not exist in our request query params.
		fName := ctx.DefaultQuery("firstname", "Guest")
		lName := ctx.Query("lastname")
		ctx.String(http.StatusOK, "Welcome %s %s", fName, lName)
	})

	// Task Routes
	r.GET("/tasks", controllers.GetAllTasks)
	r.GET("/tasks/:id", controllers.GetTaskById)
	r.POST("/tasks", controllers.CreateTask)
	r.PATCH("/tasks/:id", controllers.UpdateTaskById)
	r.DELETE("/tasks/:id", controllers.DeleteTaskById)

	// Test
	r.GET("/test-get", controllers.NewTaskController().GetAllTasksController)
	r.GET("/test-get/:id", controllers.NewTaskController().GetTaskByIdController)
	r.POST("/test-post", controllers.NewTaskController().CreateTaskController)
	r.PATCH("/test-patch/:id", controllers.NewTaskController().UpdateTaskByIdController)
	r.DELETE("/test-delete/:id", controllers.NewTaskController().DeleteTaskByIdController)

	// Listen and serve on "0.0.0.0:8080"
	// r.Run()
	/*
		💡 Note: All the r.run code (above & below) will run on "localhost:8080".
		But, specifying "127.0.0.0:8080" or "localhost: 8080" will keep windows from prompting firewall popups everytime we run our server.
	*/
	// r.Run("127.0.0.1:8080")
	// r.Run("localhost:8080")
	r.Run(infrastucture.GetBaseUrl())
}
