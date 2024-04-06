package api

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	router.GET("/albums", GetAlbums)
	router.GET("/albums/:id", GetAlbumByID)
	router.POST("/albums", PostAlbum)
	router.Run("localhost:9000")
}
