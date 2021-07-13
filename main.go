package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// postAlbums adds an album from JSON received in the request body
func postAlbums(c *gin.Context) {
	var newAlbum album

	// call BindJSON to bind the received JSON
	// to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// add the new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()            // initialize a gin router
	router.GET("/albums", getAlbums)   // pass the function, without calling it
	router.POST("/albums", postAlbums) // attach the function to the request

	router.Run("localhost:8080") // attach the router to an http.Server and starts it
}

// getAlbums responds with the list of all albums as JSON
func getAlbums(c *gin.Context) {
	// serialize the JSON adding the status response
	c.IndentedJSON(http.StatusOK, albums)
}
