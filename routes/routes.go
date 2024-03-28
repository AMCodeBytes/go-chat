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
	context.HTML(http.StatusOK, "index.html", gin.H{
		"content": "You have reached the homepage of the chat application...",
	})
}
