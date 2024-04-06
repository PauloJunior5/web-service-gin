package api

import (
	"net/http"

	"github.com/PauloJunior5/web-service-gin/albums"
	"github.com/PauloJunior5/web-service-gin/albums/entity"
	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums.Albums)
}

func PostAlbum(c *gin.Context) {
	var newAlbum entity.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums.Albums = append(albums.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, album := range albums.Albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}
