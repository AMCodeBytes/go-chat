package main

import (
	"github.com/AMCodeBytes/go-chat/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := setupRouter()

	// Listen and serve on localhost:8080
	server.Run()
}

func setupRouter() *gin.Engine {
	// Default returns an engine instance with a logger & recovery middleware
	server := gin.Default()

	// Initialise the routes
	routes.RegisterRoutes(server)

	// Return the server instance
	return server
}
