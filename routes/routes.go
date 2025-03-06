package routes

import (
	"github.com/Thomika1/RestfulGoAPI/handlers"
	"github.com/gin-gonic/gin"
)

// set the routes
func InitializeRoutes(router *gin.Engine) {

	router.GET("/albums", handlers.GetAlbums)
	router.GET("/albums/:id", handlers.GetAlbumByID)
	router.POST("/albums", handlers.PostAlbum)

}
