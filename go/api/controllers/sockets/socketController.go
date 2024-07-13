package sockets

import (
	"example/hello/internal/initializers"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"example/hello/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSocketServer struct {
	clients   map[*websocket.Conn]bool
	broadcast chan models.ChatMessage
	mutex     sync.Mutex
}

var WebSocket = WebSocketServer{
	clients:   make(map[*websocket.Conn]bool),
	broadcast: make(chan models.ChatMessage),
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnections(c *gin.Context) {
	// Upgrade de la connexion HTTP en WebSocket
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Failed to upgrade to WebSocket: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade to WebSocket"})
		return
	}
	defer ws.Close()

	// Ajouter le client à la liste des clients connectés
	WebSocket.mutex.Lock()
	WebSocket.clients[ws] = true
	WebSocket.mutex.Unlock()
	log.Printf("New client connected")

	// Récupérer les messages pour ce groupe spécifique
	groupeVoyageID := extractGroupeVoyageIDFromRequest(c)
	messages, err := GetChatMessagesByGroupID(groupeVoyageID)
	if err != nil {
		log.Printf("Error retrieving chat messages: %v", err)
		// Gérer l'erreur selon votre besoin
	} else {
		// Envoyer les messages au client WebSocket connecté
		for _, msg := range messages {
			if err := ws.WriteJSON(msg); err != nil {
				log.Printf("Error sending JSON message: %v", err)
				WebSocket.mutex.Lock()
				delete(WebSocket.clients, ws)
				WebSocket.mutex.Unlock()
				return
			}
		}
	}

	// Boucle de lecture des messages entrants du client WebSocket
	for {
		var msg models.ChatMessage
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error reading JSON: %v", err)
			WebSocket.mutex.Lock()
			delete(WebSocket.clients, ws)
			WebSocket.mutex.Unlock()
			break
		}

		// Ajouter la date de création actuelle au message
		msg.Created = time.Now()

		// Enregistrer le message dans la base de données
		if err := saveChatMessage(&msg); err != nil {
			log.Printf("Error saving chat message: %v", err)
			continue
		}

		// Diffuser le message à tous les clients WebSocket connectés
		log.Printf("Message received and saved: %+v", msg)
		WebSocket.broadcast <- msg
	}

	log.Printf("Client disconnected")
}

func HandleMessages() {
	for {
		msg := <-WebSocket.broadcast
		log.Printf("Broadcasting message: %+v", msg)
		WebSocket.mutex.Lock()
		for client := range WebSocket.clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Error writing JSON: %v", err)
				client.Close()
				delete(WebSocket.clients, client)
			}
		}
		WebSocket.mutex.Unlock()
	}
}

func saveChatMessage(msg *models.ChatMessage) error {
	if err := initializers.DB.Create(msg).Error; err != nil {
		return err
	}

	// Charger les informations de l'utilisateur
	if err := initializers.DB.Preload("User").First(msg, msg.ID).Error; err != nil {
		return err
	}

	return nil
}

func GetChatMessagesByGroupID(groupeVoyageID uint) ([]models.ChatMessage, error) {
	var messages []models.ChatMessage
	if err := initializers.DB.Where("groupe_voyage_id = ?", groupeVoyageID).Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

func extractGroupeVoyageIDFromRequest(c *gin.Context) uint {
	groupeVoyageIDStr := c.Param("groupe_voyage_id")
	groupeVoyageID, err := strconv.ParseUint(groupeVoyageIDStr, 10, 64)
	if err != nil {
		// Gérer l'erreur de conversion ici, par exemple renvoyer une erreur HTTP 400
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid groupe voyage ID"})
		return 0
	}
	return uint(groupeVoyageID)
}

func GetPreviousMessages(c *gin.Context) {
	groupeVoyageID := extractGroupeVoyageIDFromRequest(c)
	messages, err := GetChatMessagesByGroupID(groupeVoyageID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve messages"})
		return
	}
	c.JSON(http.StatusOK, messages)
}
