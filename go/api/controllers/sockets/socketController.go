package sockets

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Structure pour gérer les clients WebSocket et la diffusion des messages
type WebSocketServer struct {
	clients   map[*websocket.Conn]bool
	broadcast chan ChatMessage
	mutex     sync.Mutex
}

var WebSocket = WebSocketServer{
	clients:   make(map[*websocket.Conn]bool),
	broadcast: make(chan ChatMessage),
}

// Structure pour les messages de chat
type ChatMessage struct {
	GroupID uint   `json:"group_id"`
	UserID  uint   `json:"user_id"`
	Content string `json:"content"`
}

// Fonction pour upgrader les connexions HTTP en WebSocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Fonction pour gérer les connexions WebSocket
func HandleConnections(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade to WebSocket"})
		return
	}
	defer ws.Close()

	WebSocket.mutex.Lock()
	WebSocket.clients[ws] = true
	WebSocket.mutex.Unlock()

	for {
		var msg ChatMessage
		err := ws.ReadJSON(&msg)
		if err != nil {
			WebSocket.mutex.Lock()
			delete(WebSocket.clients, ws)
			WebSocket.mutex.Unlock()
			break
		}
		WebSocket.broadcast <- msg
	}
}

// Fonction pour gérer la diffusion des messages
func HandleMessages() {
	for {
		msg := <-WebSocket.broadcast
		WebSocket.mutex.Lock()
		for client := range WebSocket.clients {
			err := client.WriteJSON(msg)
			if err != nil {
				client.Close()
				delete(WebSocket.clients, client)
			}
		}
		WebSocket.mutex.Unlock()
	}
}
