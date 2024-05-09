package main

import (
	initializers2 "example/hello/internal/initializers"
	models2 "example/hello/internal/models"
)

func init() {
	initializers2.LoadEnvVariables()
	initializers2.ConnectToDatabase()
}

func main() {
	initializers2.DB.AutoMigrate(
		&models2.User{},
		&models2.Destination{},
		&models2.Hotel{},
		&models2.Feedback{},
		&models2.GroupeVoyage{})
}
