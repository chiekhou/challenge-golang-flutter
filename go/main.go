package main

import (
	"example/hello/api/routes"
	_ "example/hello/docs"
	"example/hello/config"
	initializers2 "example/hello/internal/initializers"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
    config.LoadConfig()
	initializers2.LoadEnvVariables()
	initializers2.ConnectToDatabase()
}

func main() {


	// Servir des fichiers statiques depuis le dossier assets
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erreur lors du chargement du fichier .env : %v", err)
	}

	trustedProxiesEnv := os.Getenv("TRUSTED_PROXIES")
	var trustedProxies []string
	if trustedProxiesEnv != "" {
		trustedProxies = strings.Split(trustedProxiesEnv, ",")
	}

	server := gin.Default()

	if err := server.SetTrustedProxies(trustedProxies); err != nil {
		log.Fatalf("Erreur lors de la configuration des proxys de confiance : %v", err)
	}

	// Configure CORS
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	// Configurer le chemin pour servir les fichiers statiques
	server.Static("/images", "./assets/images")

	routes.RegisterRoutes(server)
	routes.VoyageRoutes(server)
	routes.DestinationRoutes(server)
	routes.ActivityRoutes(server)
	routes.FlippingRoutes(server)
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := server.Run(":8080")
	if err != nil {
		return
	}
}
