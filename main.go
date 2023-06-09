package main

import (
	"fmt"
	"net/http"
	"rnd-surajan-gin/api/services"
	"rnd-surajan-gin/database"
	"rnd-surajan-gin/infrastucture"
	"rnd-surajan-gin/models"

	"github.com/gin-gonic/gin"
)

// Album Struct
// 💡 Note: If named CamelCased, the struct can be imported and used by other modules.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Slice from "album" struct
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// GET albums
func getAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, albums)
}

// POST an album
func postAlbums(ctx *gin.Context) {
	var newAlbum album

	// 💡 Note: Here semicolon ";" is used to separate the initialization statement and the actual conditional statement.
	if err := ctx.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

// GET a specific album
func getAlbumById(ctx *gin.Context) {
	id := ctx.Param("id")
	for _, a := range albums {
		if a.ID == id {
			ctx.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{
		"message": "Album not found",
	})
}

// DELETE a specific album
func deleteAlbumById(ctx *gin.Context) {
	id := ctx.Param("id")
	indexToDelete := -1
	for i, a := range albums {
		if a.ID == id {
			indexToDelete = i
		}
	}
	// Remove the item from the slice if found
	if indexToDelete != -1 {
		albums = append(albums[:indexToDelete], albums[indexToDelete+1:]...)
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully deleted album of id: %s", id),
		})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Album not found",
		})
	}
}

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

	// Album Routes
	r.GET("/albums", getAlbums)
	r.GET("/albums/:id", getAlbumById)
	r.POST("/albums", postAlbums)
	r.DELETE("/albums/:id", deleteAlbumById)

	// Task Routes
	r.GET("/tasks", services.GetAllTasks)
	r.GET("/tasks/:id", services.GetTaskById)
	r.POST("/tasks", services.CreateTask)
	r.PATCH("/tasks/:id", services.UpdateTaskById)
	r.DELETE("/tasks/:id", services.DeleteTaskById)

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
