package routes

import (
	"net/http"

	"github.com/AMCodeBytes/go-chat/handlers"
	"github.com/AMCodeBytes/go-chat/models"
	"github.com/AMCodeBytes/go-chat/utilities"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Home
	server.GET("/", home)

	// Sign up
	server.GET("/register", getRegisterPage)
	server.POST("/register", register)

	// Login
	server.GET("/login", getLoginPage)
	server.POST("/login", login)

	server.GET("/ws", handlers.Chat)
}

func home(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{
		"content": "You have reached the homepage of the chat application...",
	})
}

func getRegisterPage(context *gin.Context) {
	context.HTML(http.StatusOK, "register.html", gin.H{
		"content": "Please sign up to be able to chat.",
	})
}

func register(context *gin.Context) {
	var user models.User

	err := context.ShouldBind(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse the request data."})
		return
	}

	hashedPassword, err := utilities.HashPassword(user.Password)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to has password."})
		return
	}

	user.Password = hashedPassword

	userID, err := utilities.GenerateUUID()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate UUID."})
		return
	}

	user.ID = userID
	user.Create()

	context.HTML(http.StatusCreated, "register.html", gin.H{"message": "User created!", "user": user})
}

func getLoginPage(context *gin.Context) {
	context.HTML(http.StatusOK, "login.html", gin.H{
		"content": "Please use your credentials to log in.",
	})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBind(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	user.ID, err = user.Login()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials."})
		return
	}

	if user.ID == "" {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid credentials."})
		return
	}

	token, err := utilities.GenerateToken(user.Username, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate token."})
		return
	}

	context.HTML(http.StatusOK, "login.html", gin.H{"message": "User logged in!", "token": token})
}
