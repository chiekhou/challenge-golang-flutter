package main

import (
	"example/hello/initializers"
	"example/hello/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	initializers.DB.AutoMigrate(
		&models.User{},
		&models.Destination{},
		&models.Hotel{},
		&models.Feedback{},
		&models.GroupeVoyage{})
}
