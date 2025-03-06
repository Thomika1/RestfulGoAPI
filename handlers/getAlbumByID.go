package handlers

import (
	"net/http"

	"github.com/Thomika1/RestfulGoAPI/schemas"
	"github.com/gin-gonic/gin"
)

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range schemas.Albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})

}
