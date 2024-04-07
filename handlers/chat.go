package handlers

import (
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsChan = make(chan WsPayload)
var clients = make(map[WsConnection]string)

type WsConnection struct {
	*websocket.Conn
}

type WsResponse struct {
	Action         string   `json:"action"`
	Message        string   `json:"message"`
	MessageType    string   `json:"message_type"`
	ConnectedUsers []string `json:"connected_users"`
}

type WsPayload struct {
	Action   string       `json:"action"`
	Username string       `json:"username"`
	Message  string       `json:"message"`
	Conn     WsConnection `json:"-"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Chat(context *gin.Context) {
	ws, err := upgrader.Upgrade(context.Writer, context.Request, nil)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to establish a connection to the websocket."})
		return
	}

	// defer ws.Close()

	conn := WsConnection{Conn: ws}
	clients[conn] = ""

	// for {
	// conn.WriteMessage(websocket.TextMessage, []byte("Hello world! This is a websocket connection"))
	// conn.WriteJSON()
	// 	time.Sleep(time.Second)
	// }

	go ListenForWs(&conn)
}

func ListenForWs(conn *WsConnection) {
	var payload WsPayload

	for {
		err := conn.ReadJSON(&payload)
		fmt.Println(payload)

		if err != nil {
			// Do nothing
		} else {
			payload.Conn = *conn
			fmt.Println(payload.Conn)
			wsChan <- payload
		}
	}
}

func ListenToWsChannel() {
	var response WsResponse

	fmt.Println("chan")

	for {
		e := <-wsChan

		fmt.Println(e.Message)

		switch e.Action {
		case "username":
			// List of all users and send via broadcast
		case "left":
			// handle user leaving
			response.Action = "list_users"

			delete(clients, e.Conn)

			users := getUserList()
			response.ConnectedUsers = users

			broadcastToAll(response)
		case "broadcast":
			response.Action = "broadcast"
			response.Message = fmt.Sprintf("<strong>%s</strong>: %s", e.Username, e.Message)

			broadcastToAll(response)
		}
	}
}

func broadcastToAll(response WsResponse) {
	for client := range clients {
		err := client.WriteJSON(response)

		if err != nil {
			log.Println("websocket err")
			_ = client.Close()
			delete(clients, client)
		}
	}
}

func getUserList() []string {
	var userList []string

	for _, x := range clients {
		if x != "" {
			userList = append(userList, x)
		}
	}

	sort.Strings(userList)
	return userList
}
