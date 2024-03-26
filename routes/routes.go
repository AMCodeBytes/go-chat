package routes

import (
	"net/http"

	"github.com/AMCodeBytes/go-chat/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Home
	server.GET("/", home)

	server.GET("/ws", handlers.Chat)
}

func home(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Welcome home."})
}
