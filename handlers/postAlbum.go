package handlers

import (
	"fmt"
	"net/http"

	"github.com/Thomika1/RestfulGoAPI/schemas"
	"github.com/gin-gonic/gin"
)

func PostAlbum(c *gin.Context) {
	var newAlbum schemas.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		fmt.Printf("Error binding the album: %v", err)
	}

	schemas.Albums = append(schemas.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)

}
