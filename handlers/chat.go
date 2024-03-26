package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Chat(context *gin.Context) {
	conn, err := upgrader.Upgrade(context.Writer, context.Request, nil)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to establish a connection to the websocket."})
		return
	}

	defer conn.Close()

	for {
		conn.WriteMessage(websocket.TextMessage, []byte("Hello world! This is a websocket connection"))
		time.Sleep(time.Second)
	}
}
