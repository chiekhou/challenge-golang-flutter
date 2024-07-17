package main

import (
	"example/hello/api/routes"
	"example/hello/config"
	_ "example/hello/docs"
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

func deploiement(c *gin.Context) {
	c.String(http.StatusOK, "Déploiement réussi!")
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

	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	origins := strings.Split(allowedOrigins, ",")
	// Configure CORS
	config := cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}

	server.Use(cors.New(config))

	// Configurer le chemin pour servir les fichiers statiques
	server.Static("/images", "./assets/images")

	/*server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Bienvenue à la racine!",
		})
	})*/

	server.GET("/", deploiement)

	routes.RegisterRoutes(server)
	routes.VoyageRoutes(server)
	routes.DestinationRoutes(server)
	routes.ActivityRoutes(server)
	routes.FlippingRoutes(server)
	routes.RootRoutes(server)

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on port %s", port)
	if err := server.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
