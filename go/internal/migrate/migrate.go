package main

import (
	initializers2 "example/hello/internal/initializers"
	models2 "example/hello/internal/models"
	"fmt"
	"log"
	"time"
)

func init() {
	initializers2.LoadEnvVariables()
	initializers2.ConnectToDatabase()
}

func main() {
	//Drop la BDD afin de faire de nouvelle migrations
	// initializers2.DB.Migrator().DropTable(

	// 	&models2.Activity{},
	// 	&models2.Destination{},
	// 	&models2.Hotel{},
	// 	&models2.Feedback{},
	// 	&models2.GroupeVoyage{},
	// 	&models2.Option{},
	// 	&models2.Role{},
	// 	&models2.User{},
	// 	&models2.Voyage{},
	// 	&models2.DestinationActivity{},
	// 	&models2.VoyageActivity{},
	// )

	// Supprimer explicitement les tables de jointure
	// initializers2.DB.Migrator().DropTable("destination_activity", "destination_activities", "voyage_activities", "voyage_activity")
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
		&models2.Option{})
	if err != nil {
		return
	}

	// Exemple d'ajout de données
	activity1 := models2.Activity{
		Name:        "Paris Tour",
		Image:       "paris.jpg",
		Destination: "Paris",
		Price:       1500.00,
		Status:      0,
		Address:     "123 Rue de la Tour Eiffel, Paris",
		Longitude:   48.8584,
		Latitude:    2.2945,
	}
	activity2 := models2.Activity{
		Name:        "Rome Adventure",
		Image:       "rome.jpg",
		Destination: "Rome",
		Price:       1200.00,
		Status:      0,
		Address:     "456 Via del Colosseo, Rome",
		Longitude:   41.9028,
		Latitude:    12.4964,
	}

	initializers2.DB.Create(&activity1)
	initializers2.DB.Create(&activity2)

	// Convertir la chaîne de date en une valeur time.Time
	dateString := "2024-07-01"
	today, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		// Gérer l'erreur de parsing de la date
		// Par exemple, renvoyer une erreur ou fournir une valeur par défaut

		return
	}

	voyage := models2.Voyage{
		Destination: "Paris",
		Activities:  []models2.Activity{activity1, activity2},
		Date:        today,
	}

	voyage2 := models2.Voyage{
		Destination: "DIAGUILY",
		Activities:  []models2.Activity{activity1, activity2},
		Date:        today,
	}

	initializers2.DB.Create(&voyage)
	initializers2.DB.Create(&voyage2)

	// Associer les activités au voyage
	var voyages []models2.Voyage
	initializers2.DB.Preload("Activities").Preload("Destination").Find(&voyages)

	log.Printf("Voyages: %+v\n", voyages)
	fmt.Println("Voyage", voyages)

	// Créer des instances de destinations et d'activités
	destinations := []models2.Destination{
		{
			Name:  "Paris",
			Image: "localhost:8080/images/paris.jpg",
			Activities: []models2.Activity{
				{Image: "localhost:8080/images/activities/louvre.jpg", Name: "Louvre", Destination: "Paris", Price: 12.0, Address: "456 Via del Colosseo, Rome"},
				{Image: "localhost:8080/images/activities/chaumont.jpg", Name: "Chaumont", Destination: "Paris", Price: 0.0, Address: "456 Via del Colosseo, Rome"},
				{Image: "localhost:8080/images/activities/dame.jpg", Name: "Notre Dame", Destination: "Paris", Price: 0.0, Address: "456 Via del Colosseo, Rome"},
				{Image: "localhost:8080/images/activities/defense.jpg", Name: "La défense", Destination: "Paris", Price: 0.0, Address: "456 Via del Colosseo, Rome"},
				{Image: "localhost:8080/images/activities/effeil.jpg", Name: "Tour Eiffel", Destination: "Paris", Price: 15.0, Address: "456 Via del Colosseo, Rome"},
				{Image: "localhost:8080/images/activities/luxembourg.jpg", Name: "Jardin Luxembourg", Destination: "Paris", Price: 0.0, Address: "456 Via del Colosseo, Rome"},
				{Image: "localhost:8080/images/activities/mitterrand.jpg", Name: "Bibliothèque Mitterrand", Destination: "Paris", Price: 0.0, Address: "456 Via del Colosseo, Rome"},
				{Image: "localhost:8080/images/activities/montmartre.jpg", Name: "Montmartre", Destination: "Paris", Price: 0.0, Address: "456 Via del Colosseo, Rome"},
				{Image: "localhost:8080/images/activities/catacombe.jpg", Name: "Catacombes", Destination: "Paris", Price: 10.0, Address: "456 Via del Colosseo, Rome"},
			},
		},
		{
			Name:  "Lyon",
			Image: "localhost:8080/images/lyon.jpg",
			Activities: []models2.Activity{
				{Image: "localhost:8080/images/activities/lyon_opera.jpg", Name: "Opéra", Destination: "Lyon", Price: 100.0, Address: "456 Via del Colosseo, Rome"},
				{Image: "localhost:8080/images/activities/lyon_bellecour.jpg", Name: "Place Bellecour", Destination: "Lyon", Price: 0.0, Address: "456 Via del Colosseo, Rome"},
				{Image: "localhost:8080/images/activities/lyon_basilique.jpg", Name: "Basilique St-Pierre", Destination: "Lyon", Price: 10.0, Address: "456 Via del Colosseo, Rome"},
				{Image: "localhost:8080/images/activities/lyon_mairie.jpg", Name: "Mairie", Destination: "Lyon", Price: 0.0, Address: "456 Via del Colosseo, Rome"},
			},
		},
		{
			Name:  "Nice",
			Image: "localhost:8080/images/nice.jpg",
			Activities: []models2.Activity{
				{Image: "localhost:8080/images/activities/nice_orthodox.jpg", Name: "Eglise Orthodoxe", Destination: "Nice", Price: 5.0, Address: "456 Via del Colosseo, Rome"},
				{Image: "localhost:8080/images/activities/nice_riviera.jpg", Name: "Riviera", Destination: "Nice", Price: 0.0, Address: "456 Via del Colosseo, Rome"},
				{Image: "localhost:8080/images/activities/nice_promenade.jpg", Name: "Promenade des Anglais", Destination: "Nice", Price: 0.0, Address: "456 Via del Colosseo, Rome"},
				{Image: "localhost:8080/images/activities/nice_opera.jpg", Name: "Opéra", Destination: "Nice", Price: 100.0, Address: "456 Via del Colosseo, Rome"},
			},
		},
	}

	// Enregistrer les instances dans la base de données
	for _, destination := range destinations {
		result := initializers2.DB.Create(&destination)
		if result.Error != nil {
			log.Println("Failed to insert destination:", destination.Name, result.Error)
		}
	}

}
