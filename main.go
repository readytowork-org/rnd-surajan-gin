package main

import (
	"rnd-surajan-gin/bootstrap"
	"rnd-surajan-gin/database"
	"rnd-surajan-gin/infrastucture"
	"rnd-surajan-gin/models"

	"go.uber.org/fx"
)

func init() {
	// Initialize Env
	infrastucture.EnvInit()
	// DB connection
	database.ConnectDB()
	// Migrate the schema
	database.DB.AutoMigrate(&models.Task{})
}

// func main() {
// 	// Gin Server
// 	r := gin.Default()

// 	// Using Query Parameters
// 	r.GET("/welcome-home", func(ctx *gin.Context) {
// 		// This puts "Guest" in fName if "?firstname=" does not exist in our request query params.
// 		fName := ctx.DefaultQuery("firstname", "Guest")
// 		lName := ctx.Query("lastname")
// 		ctx.String(http.StatusOK, "Welcome %s %s", fName, lName)
// 	})

// 	// Routes
// 	routes.TasksRouteSetup(r)

// 	// Listen and serve on "0.0.0.0:8080"
// 	// r.Run()
// 	/*
// 		💡 Note: All the r.run code (above & below) will run on "localhost:8080".
// 		But, specifying "127.0.0.0:8080" or "localhost: 8080" will keep windows from prompting firewall popups everytime we run our server.
// 	*/
// 	// r.Run("127.0.0.1:8080")
// 	// r.Run("localhost:8080")
// 	r.Run(infrastucture.GetBaseUrl())
// }

func main() {
	fx.New(bootstrap.Module).Run()
}
