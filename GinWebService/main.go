package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

func getAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(context *gin.Context) {
	var newAlbum album

	if err := context.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	context.IndentedJSON(http.StatusCreated, newAlbum)

}

func getAlbumByID(context *gin.Context) {
	id := context.Param("id")

	for _, a := range albums {
		if a.ID == id {
			context.IndentedJSON(http.StatusOK, a)
			return 
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}


