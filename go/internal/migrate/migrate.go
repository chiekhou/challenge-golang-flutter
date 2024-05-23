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
	//Drop la BDD afin de faire de nouvelle migrations
	/*initializers2.DB.Migrator().DropTable(
	&models2.User{},
	&models2.Destination{},
	&models2.Hotel{},
	&models2.Feedback{},
	&models2.GroupeVoyage{},
	&models2.Option{},
	&models2.Role{})*/

	err := initializers2.DB.AutoMigrate(
		&models2.User{},
		&models2.Destination{},
		&models2.Hotel{},
		&models2.Feedback{},
		&models2.GroupeVoyage{},
		&models2.Option{})
	if err != nil {
		return
	}
}
