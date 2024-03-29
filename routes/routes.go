package routes

import (
	"net/http"

	"github.com/AMCodeBytes/go-chat/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Home
	server.GET("/", home)

	// Sign up
	server.GET("/register", register)

	server.GET("/ws", handlers.Chat)
}

func home(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{
		"content": "You have reached the homepage of the chat application...",
	})
}

func register(context *gin.Context) {
	context.HTML(http.StatusOK, "register.html", gin.H{
		"content": "Please sign up to be able to chat.",
	})
}
