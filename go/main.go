package main

import (
	"example/hello/api/controllers/sockets"
	"example/hello/api/routes"
	"example/hello/internal/initializers"
	"example/hello/internal/migrate"

	_ "example/hello/docs"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	db = initializers.DB // Initialiser la variable db avec la connexion existante
}

func main() {
	// Servir des fichiers statiques depuis le dossier assets
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	server := gin.Default()

	// Configurer le chemin pour servir les fichiers statiques
	server.Static("/images", "./assets/images")

	// Enregistrer les routes
	routes.RegisterRoutes(server)
	routes.VoyageRoutes(server)
	routes.DestinationRoutes(server)
	routes.ActivityRoutes(server)
	routes.FlippingRoutes(server)
	routes.SocketRoutes(server)

	// Route pour gérer les connexions WebSocket
	server.GET("/ws", sockets.HandleConnections)

	// Lancer la gestion des messages en arrière-plan
	go sockets.HandleMessages()

	// Swagger documentation
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Appeler la fonction de migration après la connexion à la base de données
	if err := migrate.Migrate(db); err != nil {
		panic(err)
	}

	// Démarrer le serveur
	err := server.Run(":8080")
	if err != nil {
		panic(err)
	}
}
