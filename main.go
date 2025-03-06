package main

import (
	"github.com/Thomika1/RestfulGoAPI/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/albums", handlers.GetAlbums)
	router.GET("/albums/:id", handlers.GetAlbumByID)
	router.POST("/albums", handlers.PostAlbum)

	router.Run("localhost:8080")
}
