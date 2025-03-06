package handlers

import (
	"net/http"

	"github.com/Thomika1/RestfulGoAPI/schemas"
	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, schemas.Albums)

}
