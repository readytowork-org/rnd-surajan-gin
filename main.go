package main

import (
	"net/http"

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

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/albums", getAlbums)

	// listen and serve on 0.0.0.0:8080
	// r.Run()

	// Both will run on "localhost:8080", but specifying 127.0.0.0:8080 will keep windows from prompting firewall popups everytime we run our server.
	r.Run("127.0.0.1:8080")
}
