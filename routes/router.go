package routes

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()

	InitializeRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // Porta padrão se não for definida
	}
	router.Run(":" + port)

	//router.Run("localhost:8080")
}
