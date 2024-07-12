package main

import (
	initializers2 "example/hello/internal/initializers"
	models2 "example/hello/internal/models"
	"example/hello/config"
	"example/hello/internal/seed"

)

func init() {
    config.LoadConfig()
	initializers2.LoadEnvVariables()
	initializers2.ConnectToDatabase()
}

func main() {
	//Drop la BDD afin de faire de nouvelle migrations
    /*	initializers2.DB.Migrator().DropTable(

    		&models2.User{},
    		&models2.Activity{},
    		&models2.Destination{},
    		&models2.DestinationActivity{},
    		&models2.Hotel{},
    		&models2.Feedback{},
    		&models2.GroupeVoyage{},
    		&models2.Voyage{},
    		&models2.VoyageActivity{},
    		&models2.GroupeMembers{},
    		&models2.ChatMessage{},
    		&models2.DestinationHotel{},
    		&models2.VoyageHotel{})*/

    	// Supprimer explicitement les tables de jointure
    	//initializers2.DB.Migrator().DropTable("destination_activity", "destination_activities", "destination_hotels", "voyage_activities", "voyage_activity", "voyage_hotels", "group_voyage", "groupe_members")
    	err := initializers2.DB.AutoMigrate(
    		&models2.User{},
    		&models2.Activity{},
    		&models2.Destination{},
    		&models2.DestinationActivity{},
    		&models2.Hotel{},
    		&models2.Feedback{},
    		&models2.GroupeVoyage{},
    		&models2.Voyage{},
    		&models2.VoyageActivity{},
    		&models2.GroupeMembers{},
    		&models2.ChatMessage{},
    		&models2.DestinationHotel{},
    		&models2.VoyageHotel{})

		seed.SeedData(initializers2.DB,config.AppConfig)

	if err != nil {
		return
	}

}
