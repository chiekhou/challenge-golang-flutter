package main

import (
	"example/hello/api/controllers/sockets"
	"example/hello/api/routes"
	_ "example/hello/docs"
	initializers2 "example/hello/internal/initializers"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializers2.LoadEnvVariables()
	initializers2.ConnectToDatabase()
}

func main() {
	// Servir des fichiers statiques depuis le dossier assets
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	server := gin.Default()

	// Configurer le chemin pour servir les fichiers statiques
	server.Static("/images", "./assets/images")

	routes.RegisterRoutes(server)
	routes.VoyageRoutes(server)
	routes.DestinationRoutes(server)
	routes.ActivityRoutes(server)

	// Route pour gérer les connexions WebSocket
	server.GET("/ws", sockets.HandleConnections) // Utiliser controllers.HandleConnections

	// Lancer la gestion des messages en arrière-plan
	go sockets.HandleMessages()

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := server.Run(":8080")
	if err != nil {
		return
	}
}
