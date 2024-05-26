package main

import (
	"example/hello/api/routes"
	_ "example/hello/docs"
	initializers2 "example/hello/internal/initializers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializers2.LoadEnvVariables()
	initializers2.ConnectToDatabase()
}

func main() {
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := server.Run(":8080")
	if err != nil {
		return
	}
}
