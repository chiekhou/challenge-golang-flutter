package main

import (
	initializers2 "example/hello/internal/initializers"
	models2 "example/hello/internal/models"
	"log"
)

func init() {
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
			&models2.VoyageHotel{})

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
		if err != nil {
			return
		}

		// Convertir la chaîne de date en une valeur time.Time
		/*dateString := "2024-07-01"
		today, err := time.Parse("2006-01-02", dateString)
		if err != nil {
			// Gérer l'erreur de parsing de la date
			// Par exemple, renvoyer une erreur ou fournir une valeur par défaut
			return
		}*/

	// Créer des instances de destinations et d'activités
	destinations := []models2.Destination{

		{
			Name:  "Phuket",
			Image: "http://10.0.2.2:8080/images/phuket.jpeg",
			Activities: []models2.Activity{
				{Image: "http://10.0.2.2:8080/images/activities/banana-beach.jpeg", Name: "Banana Beach", Destination: "Phuket", Price: 12.4, Address: "Phuket", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/activities/bangla-road-phuket.jpeg", Name: "Bangla Road", Destination: "Phuket", Price: 13.3, Address: "Phuket", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/activities/kata-beach-phuket.jpeg", Name: "Kata Beach", Destination: "Phuket", Price: 18.3, Address: "Phuket", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/activities/grand-bouddah.jpeg", Name: "Buddah Blanc", Destination: "Phuket", Price: 14.3, Address: "Phuket", Longitude: 41.9028, Latitude: 12.4964},
			},
			Hotels: []models2.Hotel{
				{Image: "http://10.0.2.2:8080/images/hotels/tidephuket-phuket.jpeg", Name: "Tide Phuket", Destination: "Phuket", Price: 60.4, Address: "Phuket", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/hotels/royalphuketcity-phuket.jpeg", Name: "Royal City", Destination: "Phuket", Price: 90.3, Address: "Phuket", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/hotels/pavillonhotel-phuket.jpeg", Name: "Pavillon Hotel", Destination: "Phuket", Price: 67.3, Address: "Phuket", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/hotels/islandescape-phuket.jpeg", Name: "Island Escade", Destination: "Phuket", Price: 89.3, Address: "Phuket", Longitude: 41.9028, Latitude: 12.4964},
			},
		},
		{
			Name:  "Dubaï",
			Image: "http://10.0.2.2:8080/images/dubai.jpeg",
			Activities: []models2.Activity{
				{Image: "http://10.0.2.2:8080/images/activities/burj-arab.jpeg", Name: "Burj Arab", Destination: "Dubaï", Price: 35.4, Address: "Dubaï", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/activities/burj-khalifa.jpeg", Name: "Burj Khalifa", Destination: "Dubaï", Price: 30.4, Address: "Dubaï", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/activities/dubai-fountain.jpeg", Name: "Fountain", Destination: "Dubaï", Price: 15.2, Address: "Dubaï", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/activities/dubai-aquarium-underwater.jpg", Name: "Aquarium", Destination: "Dubaï", Price: 30.2, Address: "Dubaï", Longitude: 41.9028, Latitude: 12.4964},
			},

			Hotels: []models2.Hotel{
				{Image: "http://10.0.2.2:8080/images/hotels/bitmore-dubai.jpeg", Name: "Bitmore Hotel", Destination: "Dubaï", Price: 55.4, Address: "Dubaï", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/hotels/canalcentral-dubai.jpeg", Name: "Canal Hôtel", Destination: "Dubaï", Price: 70.4, Address: "Dubaï", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/hotels/jumeirah-dubai.jpeg", Name: "Jumeirah Hôtel", Destination: "Dubaï", Price: 75.2, Address: "Dubaï", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/hotels/royalhideaway-dubai.jpeg", Name: "Royal Hideway", Destination: "Dubaï", Price: 90.2, Address: "Dubaï", Longitude: 41.9028, Latitude: 12.4964},
			},
		},
		{
			Name:  "Los Angeles",
			Image: "http://10.0.2.2:8080/images/los-angeles.jpeg",
			Activities: []models2.Activity{
				{Image: "http://10.0.2.2:8080/images/activities/griffith-observatory-la.jpeg", Name: "Griffith Observatory", Destination: "Los Angeles", Price: 5.30, Address: "Los Angeles", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/activities/hollywood-sign.jpg", Name: "Hollywood Sign", Destination: "Los Angeles", Price: 20.40, Address: "Los Angeles", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/activities/universal-studios-hollywood.jpg", Name: "Universal Studios", Destination: "Los Angeles", Price: 30.40, Address: "Los Angeles", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/activities/venice-beach.jpeg", Name: "Venise Beach", Destination: "Los Angeles", Price: 100.50, Address: "Los Angeles", Longitude: 41.9028, Latitude: 12.4964},
			},

        			Hotels: []models2.Hotel{
                          {Image: "http://10.0.2.2:8080/images/hotels/conrad-losangeles.jpeg", Name: "Conrad", Destination: "Los Angeles", Price: 65.30, Address: "Los Angeles",Longitude:41.9028 ,Latitude:12.4964},
                           {Image: "http://10.0.2.2:8080/images/hotels/figueroa-losangeles.jpeg", Name: "Figueroa Hôtel", Destination: "Los Angeles", Price: 80.40, Address: "Los Angeles",Longitude:41.9028 ,Latitude:12.4964},
                           {Image: "http://10.0.2.2:8080/images/hotels/proper-losangeles.jpeg", Name: "Proper Hôtel", Destination: "Los Angeles", Price: 90.40, Address: "Los Angeles",Longitude:41.9028 ,Latitude:12.4964},
                           {Image: "http://10.0.2.2:8080/images/hotels/beverlyhills-losangeles.jpeg", Name: "Venise Beach", Destination: "Los Angeles", Price: 100.50, Address: "Los Angeles",Longitude:41.9028 ,Latitude:12.4964},
                   },
        		},
       {
       			Name:  "Ibiza",
       			Image: "http://10.0.2.2:8080/images/ibiza.jpeg",
       			Activities: []models2.Activity{
       				{Image: "http://10.0.2.2:8080/images/activities/montgolfiere-ibiza.jpeg", Name: "Montgolfiere", Destination: "Ibiza", Price: 12.4, Address: "Ibiza", Longitude:41.9028 ,Latitude:12.4964},
       				{Image: "http://10.0.2.2:8080/images/activities/cala-bassa-ibiza.jpeg", Name: "Cala Bassa", Destination: "Ibiza", Price: 13.3, Address: "Ibiza", Longitude:41.9028 ,Latitude:12.4964},
       				{Image: "http://10.0.2.2:8080/images/activities/cala-saladeta.jpeg", Name: "Cala Saladeta", Destination: "Ibiza", Price: 18.3, Address: "Ibiza",Longitude:41.9028 ,Latitude:12.4964},
       				{Image: "http://10.0.2.2:8080/images/activities/es-vedra.jpeg", Name: "Es Vedra", Destination: "Ibiza", Price: 14.3, Address: "Ibiza", Longitude:41.9028 ,Latitude:12.4964},
                },
                Hotels: []models2.Hotel{
             		{Image: "http://10.0.2.2:8080/images/hotels/beachstar-ibiza.jpeg", Name: "Beach Star", Destination: "Ibiza", Price: 92.4, Address: "Ibiza", Longitude:41.9028 ,Latitude:12.4964},
             		{Image: "http://10.0.2.2:8080/images/hotels/blesshotel-ibiza.jpeg", Name: "Bless Hôtel", Destination: "Ibiza", Price: 73.3, Address: "Ibiza", Longitude:41.9028 ,Latitude:12.4964},
             		{Image: "http://10.0.2.2:8080/images/hotels/grupotelibizabeach-ibiza.jpeg", Name: "Grupotel Ibiza", Destination: "Ibiza", Price: 78.3, Address: "Ibiza",Longitude:41.9028 ,Latitude:12.4964},
             		{Image: "http://10.0.2.2:8080/images/hotels/ryanslolas-ibiza.jpeg", Name: "Ryanslolas", Destination: "Ibiza", Price: 94.3, Address: "Ibiza", Longitude:41.9028 ,Latitude:12.4964},
                      },
       		},
       		{
       			Name:  "Venise",
       			Image: "http://10.0.2.2:8080/images/venise.jpeg",
       			Activities: []models2.Activity{
       				{Image: "http://10.0.2.2:8080/images/activities/saint-marc.jpeg", Name: "Saint-Marc", Destination: "Venise", Price: 10.4, Address: "Venise", Longitude:41.9028 ,Latitude:12.4964},
       				{Image: "http://10.0.2.2:8080/images/activities/palais-doges-venise.jpeg", Name: "Palais Doges", Destination: "Venise", Price: 9.5, Address: "Venise",Longitude:41.9028 ,Latitude:12.4964},
       				{Image: "http://10.0.2.2:8080/images/activities/grand-canal-venise.jpeg", Name: "Grand Canal", Destination: "Venise", Price: 10.2, Address: "Venise",Longitude:41.9028 ,Latitude:12.4964},
       				{Image: "http://10.0.2.2:8080/images/activities/excursion-murano.jpeg", Name: "Excursion Murano", Destination: "Venise", Price: 30.2, Address: "Venise",Longitude:41.9028 ,Latitude:12.4964},
       			},
       			Hotels: []models2.Hotel{
                    {Image: "http://10.0.2.2:8080/images/hotels/cappellis-venise.jpeg", Name: "Cappellis", Destination: "Venise", Price: 60.4, Address: "Venise", Longitude:41.9028 ,Latitude:12.4964},
                     {Image: "http://10.0.2.2:8080/images/hotels/palazzinafortuny-venise.jpeg", Name: "Palazzina ", Destination: "Venise", Price: 69.5, Address: "Venise",Longitude:41.9028 ,Latitude:12.4964},
                     {Image: "http://10.0.2.2:8080/images/hotels/palazzocanova-venise.jpeg", Name: "Palazzo Canova", Destination: "Venise", Price: 50.2, Address: "Venise",Longitude:41.9028 ,Latitude:12.4964},
                     {Image: "http://10.0.2.2:8080/images/hotels/salutepalace-venise.jpeg", Name: "Salute Palace", Destination: "Venise", Price: 30.2, Address: "Venise",Longitude:41.9028 ,Latitude:12.4964},
                       			},
       		},
       		{
       			Name:  "Mykonos",
       			Image: "http://10.0.2.2:8080/images/mykonos.jpeg",
       			Activities: []models2.Activity{
                    {Image: "http://10.0.2.2:8080/images/activities/petite-venise-mykonos.jpeg", Name: "Petit Venise", Destination: "Mykonos", Price: 5.30, Address: "Mykonos",Longitude:41.9028 ,Latitude:12.4964},
                     {Image: "http://10.0.2.2:8080/images/activities/super-paradise-beach-mykonos.jpeg", Name: "Paradise Beach", Destination: "Mykonos", Price: 5.40, Address: "Mykonos",Longitude:41.9028 ,Latitude:12.4964},
                     {Image: "http://10.0.2.2:8080/images/activities/vioma-organic-mykonos.jpeg", Name: "Vioma Organic", Destination: "Mykonos", Price: 7.50, Address: "Mykonos",Longitude:41.9028 ,Latitude:12.4964},
                    {Image: "http://10.0.2.2:8080/images/activities/kato-mili-mykonos.jpeg", Name: "Kato Mili", Destination: "Mykonos", Price: 10.50, Address: "Mykonos",Longitude:41.9028 ,Latitude:12.4964},
                   },
                Hotels: []models2.Hotel{
                    {Image: "http://10.0.2.2:8080/images/hotels/aeolos-resort-mykonos.jpeg", Name: "Aeolos Resort", Destination: "Mykonos", Price: 55.30, Address: "Mykonos",Longitude:41.9028 ,Latitude:12.4964},
                    {Image: "http://10.0.2.2:8080/images/hotels/agrari-mykonos.jpeg", Name: "Agrari", Destination: "Mykonos", Price: 95.40, Address: "Mykonos",Longitude:41.9028 ,Latitude:12.4964},
                    {Image: "http://10.0.2.2:8080/images/hotels/fteliablack-mykonos.jpeg", Name: "Ftelia Black", Destination: "Mykonos", Price: 97.50, Address: "Mykonos",Longitude:41.9028 ,Latitude:12.4964},
                    {Image: "http://10.0.2.2:8080/images/hotels/millionstar-mykonos.jpeg", Name: "Million Star", Destination: "Mykonos", Price: 100.50, Address: "Mykonos",Longitude:41.9028 ,Latitude:12.4964},
                          			},
       		},

		{
			Name:  "Paris",
			Image: "http://10.0.2.2:8080/images/paris.jpg",
			Activities: []models2.Activity{
				{Image: "http://10.0.2.2:8080/images/activities/louvre.jpg", Name: "Louvre", Destination: "Paris", Price: 12.4, Address: "Paris", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/activities/chaumont.jpg", Name: "Chaumont", Destination: "Paris", Price: 13.3, Address: "Paris", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/activities/dame.jpg", Name: "Notre Dame", Destination: "Paris", Price: 18.3, Address: "Paris", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/activities/defense.jpg", Name: "La défense", Destination: "Paris", Price: 14.3, Address: "Paris", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/activities/effeil.jpg", Name: "Tour Eiffel", Destination: "Paris", Price: 15.3, Address: "Paris", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/activities/luxembourg.jpg", Name: "Jardin Luxembourg", Destination: "Paris", Price: 20.8, Address: "Paris", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/activities/mitterrand.jpg", Name: "Bibliothèque Mitterrand", Destination: "Paris", Price: 20.9, Address: "Paris", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/activities/montmartre.jpg", Name: "Montmartre", Destination: "Paris", Price: 23.9, Address: "Paris", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/activities/catacombe.jpg", Name: "Catacombes", Destination: "Paris", Price: 10.8, Address: "Paris", Longitude: 41.9028, Latitude: 12.4964},
			},

			Hotels: []models2.Hotel{
				{Image: "http://10.0.2.2:8080/images/hotels/bowmann-paris.jpeg", Name: "Bowman", Destination: "Paris", Price: 52.4, Address: "Paris", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/hotels/odalys-city-montmarte-paris.jpeg", Name: "Odalys City", Destination: "Paris", Price: 83.3, Address: "Paris", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/hotels/saint-petersbourg-paris.jpeg", Name: "Saint-Petersbourg", Destination: "Paris", Price: 88.3, Address: "Paris", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/hotels/scarlett-paris.jpeg", Name: "Scarlett", Destination: "Paris", Price: 140.3, Address: "Paris", Longitude: 41.9028, Latitude: 12.4964},
			},
		},
		{
			Name:  "Lyon",
			Image: "http://10.0.2.2:8080/images/lyon.jpg",
			Activities: []models2.Activity{
				{Image: "http://10.0.2.2:8080/images/activities/lyon_opera.jpg", Name: "Opéra", Destination: "Lyon", Price: 100.4, Address: "456 Via del Colosseo, Rome", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/activities/lyon_bellecour.jpg", Name: "Place Bellecour", Destination: "Lyon", Price: 30.4, Address: "456 Via del Colosseo, Rome", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/activities/lyon_basilique.jpg", Name: "Basilique St-Pierre", Destination: "Lyon", Price: 10.2, Address: "456 Via del Colosseo, Rome", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/activities/lyon_mairie.jpg", Name: "Mairie", Destination: "Lyon", Price: 30.2, Address: "456 Via del Colosseo, Rome", Longitude: 41.9028, Latitude: 12.4964},
			},

			Hotels: []models2.Hotel{
				{Image: "http://10.0.2.2:8080/images/hotels/boscolo-lyon.jpeg", Name: "Boscolo", Destination: "Lyon", Price: 100.4, Address: "456 Via del Colosseo, Rome", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/hotels/intercontinental-lyon.jpeg", Name: "Intercontinental", Destination: "Lyon", Price: 130.4, Address: "456 Via del Colosseo, Rome", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/hotels/sofitel-lyon.jpeg", Name: "Sofitel", Destination: "Lyon", Price: 100.2, Address: "456 Via del Colosseo, Rome", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/hotels/villa-florentine-lyon.jpeg", Name: "Villa Florentine", Destination: "Lyon", Price: 30.2, Address: "456 Via del Colosseo, Rome", Longitude: 41.9028, Latitude: 12.4964},
			},
		},
		{
			Name:  "Nice",
			Image: "http://10.0.2.2:8080/images/nice.jpg",
			Activities: []models2.Activity{
				{Image: "http://10.0.2.2:8080/images/activities/nice_orthodox.jpg", Name: "Eglise Orthodoxe", Destination: "Nice", Price: 5.30, Address: "Nice", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/activities/nice_riviera.jpg", Name: "Riviera", Destination: "Nice", Price: 20.40, Address: "Nice", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/activities/nice_promenade.jpg", Name: "Promenade des Anglais", Destination: "Nice", Price: 30.40, Address: "Nice", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/activities/nice_opera.jpg", Name: "Opéra", Destination: "Nice", Price: 100.50, Address: "Nice", Longitude: 41.9028, Latitude: 12.4964},
			},
			Hotels: []models2.Hotel{
				{Image: "http://10.0.2.2:8080/images/hotels/piedeau-nice.jpeg", Name: "Pied D'Eau", Destination: "Nice", Price: 55.30, Address: "Nice", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/hotels/saintgothard-nice.jpg", Name: "Saint Gothard", Destination: "Nice", Price: 60.40, Address: "Nice", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/hotels/yelo-mozart-nice.jpg", Name: "Yelo Mozart", Destination: "Nice", Price: 50.40, Address: "Nice", Longitude: 41.9028, Latitude: 12.4964},
				{Image: "http://10.0.2.2:8080/images/hotels/nicehome-nice.jpg", Name: "Nice Home", Destination: "Nice", Price: 40.50, Address: "Nice", Longitude: 41.9028, Latitude: 12.4964},
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
