package seed

import (
	"fmt"
	"log"
	models2 "example/hello/internal/models"
	"example/hello/config"
	"gorm.io/gorm"
)


func SeedData(DB *gorm.DB,appConfig *config.Config) {
      baseURL := getBaseURL(appConfig)
      fmt.Printf("Base URL: %s\n", baseURL)

// Vérifier si les données existent déjà
	var count int64
	DB.Model(&models2.Destination{}).Count(&count)
	if count > 0 {
		log.Println("Les données sont déja insérer dans la base de données")
		return
	}

// Créer des instances de destinations et d'activités
	destinations := []models2.Destination{

		{
			Name:  "Phuket",
			Image: fmt.Sprintf("%s/images/phuket.jpeg",baseURL),
			Activities: []models2.Activity{
				{Image: fmt.Sprintf("%s/images/activities/banana-beach.jpeg", baseURL), Name: "Banana Beach", Destination: "Phuket", Price: 12.4, Address: "Phuket", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/bangla-road-phuket.jpeg", baseURL), Name: "Bangla Road", Destination: "Phuket", Price: 13.3, Address: "Phuket", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/kata-beach-phuket.jpeg", baseURL), Name: "Kata Beach", Destination: "Phuket", Price: 18.3, Address: "Phuket", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/grand-bouddah.jpeg", baseURL), Name: "Buddah Blanc", Destination: "Phuket", Price: 14.3, Address: "Phuket", Longitude: 41.9028, Latitude: 12.4964},
			},
			Hotels: []models2.Hotel{
				{Image: fmt.Sprintf("%s/images/hotels/tidephuket-phuket.jpeg", baseURL), Name: "Tide Phuket", Destination: "Phuket", Price: 60.4, Address: "Phuket", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/hotels/royalphuketcity-phuket.jpeg", baseURL), Name: "Royal City", Destination: "Phuket", Price: 90.3, Address: "Phuket", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/hotels/pavillonhotel-phuket.jpeg", baseURL), Name: "Pavillon Hotel", Destination: "Phuket", Price: 67.3, Address: "Phuket", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/hotels/islandescape-phuket.jpeg", baseURL), Name: "Island Escade", Destination: "Phuket", Price: 89.3, Address: "Phuket", Longitude: 41.9028, Latitude: 12.4964},
			},
		},
		{
			Name:  "Dubaï",
			Image: fmt.Sprintf("%s/images/dubai.jpeg",baseURL),
			Activities: []models2.Activity{
				{Image: fmt.Sprintf("%s/images/activities/burj-arab.jpeg", baseURL), Name: "Burj Arab", Destination: "Dubaï", Price: 35.4, Address: "Dubaï", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/burj-khalifa.jpeg", baseURL), Name: "Burj Khalifa", Destination: "Dubaï", Price: 30.4, Address: "Dubaï", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/dubai-fountain.jpeg", baseURL), Name: "Fountain", Destination: "Dubaï", Price: 15.2, Address: "Dubaï", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/dubai-aquarium-underwater.jpg", baseURL), Name: "Aquarium", Destination: "Dubaï", Price: 30.2, Address: "Dubaï", Longitude: 41.9028, Latitude: 12.4964},
			},

			Hotels: []models2.Hotel{
				{Image: fmt.Sprintf("%s/images/hotels/bitmore-dubai.jpeg", baseURL) ,Name: "Bitmore Hotel", Destination: "Dubaï", Price: 55.4, Address: "Dubaï", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/hotels/canalcentral-dubai.jpeg", baseURL) ,Name: "Canal Hôtel", Destination: "Dubaï", Price: 70.4, Address: "Dubaï", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/hotels/jumeirah-dubai.jpeg", baseURL), Name: "Jumeirah Hôtel", Destination: "Dubaï", Price: 75.2, Address: "Dubaï", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/hotels/royalhideaway-dubai.jpeg", baseURL), Name: "Royal Hideway", Destination: "Dubaï", Price: 90.2, Address: "Dubaï", Longitude: 41.9028, Latitude: 12.4964},
			},
		},
		{
			Name:  "Los Angeles",
			Image: fmt.Sprintf("%s/images/los-angeles.jpeg",baseURL),
			Activities: []models2.Activity{
				{Image: fmt.Sprintf("%s/images/activities/griffith-observatory-la.jpeg", baseURL), Name: "Griffith Observatory", Destination: "Los Angeles", Price: 5.30, Address: "Los Angeles", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/hollywood-sign.jpg", baseURL), Name: "Hollywood Sign", Destination: "Los Angeles", Price: 20.40, Address: "Los Angeles", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/universal-studios-hollywood.jpg", baseURL), Name: "Universal Studios", Destination: "Los Angeles", Price: 30.40, Address: "Los Angeles", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/venice-beach.jpeg", baseURL), Name: "Venise Beach", Destination: "Los Angeles", Price: 100.50, Address: "Los Angeles", Longitude: 41.9028, Latitude: 12.4964},
			},

			Hotels: []models2.Hotel{
				{Image: fmt.Sprintf("%s/images/hotels/conrad-losangeles.jpeg", baseURL), Name: "Conrad", Destination: "Los Angeles", Price: 65.30, Address: "Los Angeles", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/hotels/figueroa-losangeles.jpeg", baseURL), Name: "Figueroa Hôtel", Destination: "Los Angeles", Price: 80.40, Address: "Los Angeles", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/hotels/proper-losangeles.jpeg", baseURL), Name: "Proper Hôtel", Destination: "Los Angeles", Price: 90.40, Address: "Los Angeles", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/hotels/beverlyhills-losangeles.jpeg", baseURL), Name: "Venise Beach", Destination: "Los Angeles", Price: 100.50, Address: "Los Angeles", Longitude: 41.9028, Latitude: 12.4964},
			},
		},
		{
			Name:  "Ibiza",
			Image: fmt.Sprintf("%s/images/ibiza.jpeg",baseURL),
			Activities: []models2.Activity{
				{Image: fmt.Sprintf("%s/images/activities/montgolfiere-ibiza.jpeg", baseURL), Name: "Montgolfiere", Destination: "Ibiza", Price: 12.4, Address: "Ibiza", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/cala-bassa-ibiza.jpeg", baseURL), Name: "Cala Bassa", Destination: "Ibiza", Price: 13.3, Address: "Ibiza", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/cala-saladeta.jpeg", baseURL), Name: "Cala Saladeta", Destination: "Ibiza", Price: 18.3, Address: "Ibiza", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/es-vedra.jpeg", baseURL), Name: "Es Vedra", Destination: "Ibiza", Price: 14.3, Address: "Ibiza", Longitude: 41.9028, Latitude: 12.4964},
			},
			Hotels: []models2.Hotel{
				{Image: fmt.Sprintf("%s/images/hotels/beachstar-ibiza.jpeg", baseURL), Name: "Beach Star", Destination: "Ibiza", Price: 92.4, Address: "Ibiza", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/hotels/blesshotel-ibiza.jpeg", baseURL), Name: "Bless Hôtel", Destination: "Ibiza", Price: 73.3, Address: "Ibiza", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/hotels/grupotelibizabeach-ibiza.jpeg", baseURL), Name: "Grupotel Ibiza", Destination: "Ibiza", Price: 78.3, Address: "Ibiza", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/hotels/ryanslolas-ibiza.jpeg", baseURL), Name: "Ryanslolas", Destination: "Ibiza", Price: 94.3, Address: "Ibiza", Longitude: 41.9028, Latitude: 12.4964},
			},
		},
		{
			Name:  "Venise",
			Image: fmt.Sprintf("%s/images/venise.jpeg",baseURL),
			Activities: []models2.Activity{
				{Image: fmt.Sprintf("%s/images/activities/saint-marc.jpeg", baseURL), Name: "Saint-Marc", Destination: "Venise", Price: 10.4, Address: "Venise", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/palais-doges-venise.jpeg", baseURL), Name: "Palais Doges", Destination: "Venise", Price: 9.5, Address: "Venise", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/grand-canal-venise.jpeg", baseURL), Name: "Grand Canal", Destination: "Venise", Price: 10.2, Address: "Venise", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/excursion-murano.jpeg", baseURL), Name: "Excursion Murano", Destination: "Venise", Price: 30.2, Address: "Venise", Longitude: 41.9028, Latitude: 12.4964},
			},
			Hotels: []models2.Hotel{
				{Image: fmt.Sprintf("%s/images/hotels/cappellis-venise.jpeg", baseURL), Name: "Cappellis", Destination: "Venise", Price: 60.4, Address: "Venise", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/hotels/palazzinafortuny-venise.jpeg", baseURL), Name: "Palazzina ", Destination: "Venise", Price: 69.5, Address: "Venise", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/hotels/palazzocanova-venise.jpeg", baseURL), Name: "Palazzo Canova", Destination: "Venise", Price: 50.2, Address: "Venise", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/hotels/salutepalace-venise.jpeg", baseURL), Name: "Salute Palace", Destination: "Venise", Price: 30.2, Address: "Venise", Longitude: 41.9028, Latitude: 12.4964},
			},
		},
		{
			Name:  "Mykonos",
			Image: fmt.Sprintf("%s/images/mykonos.jpeg",baseURL),
			Activities: []models2.Activity{
				{Image: fmt.Sprintf("%s/images/activities/petite-venise-mykonos.jpeg", baseURL), Name: "Petit Venise", Destination: "Mykonos", Price: 5.30, Address: "Mykonos", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/super-paradise-beach-mykonos.jpeg", baseURL), Name: "Paradise Beach", Destination: "Mykonos", Price: 5.40, Address: "Mykonos", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/vioma-organic-mykonos.jpeg", baseURL), Name: "Vioma Organic", Destination: "Mykonos", Price: 7.50, Address: "Mykonos", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/kato-mili-mykonos.jpeg", baseURL), Name: "Kato Mili", Destination: "Mykonos", Price: 10.50, Address: "Mykonos", Longitude: 41.9028, Latitude: 12.4964},
			},
			Hotels: []models2.Hotel{
				{Image: fmt.Sprintf("%s/images/hotels/aeolos-resort-mykonos.jpeg", baseURL), Name: "Aeolos Resort", Destination: "Mykonos", Price: 55.30, Address: "Mykonos", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/hotels/agrari-mykonos.jpeg", baseURL), Name: "Agrari", Destination: "Mykonos", Price: 95.40, Address: "Mykonos", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/hotels/fteliablack-mykonos.jpeg", baseURL), Name: "Ftelia Black", Destination: "Mykonos", Price: 97.50, Address: "Mykonos", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/hotels/millionstar-mykonos.jpeg", baseURL), Name: "Million Star", Destination: "Mykonos", Price: 100.50, Address: "Mykonos", Longitude: 41.9028, Latitude: 12.4964},
			},
		},

		{
			Name:  "Paris",
			Image: fmt.Sprintf("%s/images/paris.jpg",baseURL),
			Activities: []models2.Activity{
				{Image: fmt.Sprintf("%s/images/activities/louvre.jpg", baseURL), Name: "Louvre", Destination: "Paris", Price: 12.4, Address: "Paris", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/chaumont.jpg", baseURL), Name: "Chaumont", Destination: "Paris", Price: 13.3, Address: "Paris", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/dame.jpg", baseURL), Name: "Notre Dame", Destination: "Paris", Price: 18.3, Address: "Paris", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/defense.jpg", baseURL), Name: "La défense", Destination: "Paris", Price: 14.3, Address: "Paris", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/effeil.jpg", baseURL), Name: "Tour Eiffel", Destination: "Paris", Price: 15.3, Address: "Paris", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/luxembourg.jpg", baseURL), Name: "Jardin Luxembourg", Destination: "Paris", Price: 20.8, Address: "Paris", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/mitterrand.jpg", baseURL), Name: "Bibliothèque Mitterrand", Destination: "Paris", Price: 20.9, Address: "Paris", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/montmartre.jpg", baseURL), Name: "Montmartre", Destination: "Paris", Price: 23.9, Address: "Paris", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/catacombe.jpg", baseURL), Name: "Catacombes", Destination: "Paris", Price: 10.8, Address: "Paris", Longitude: 41.9028, Latitude: 12.4964},
			},

			Hotels: []models2.Hotel{
				{Image: fmt.Sprintf("%s/images/hotels/bowmann-paris.jpeg", baseURL), Name: "Bowman", Destination: "Paris", Price: 52.4, Address: "Paris", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/hotels/odalys-city-montmarte-paris.jpeg", baseURL), Name: "Odalys City", Destination: "Paris", Price: 83.3, Address: "Paris", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/hotels/saint-petersbourg-paris.jpeg", baseURL), Name: "Saint-Petersbourg", Destination: "Paris", Price: 88.3, Address: "Paris", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/hotels/scarlett-paris.jpeg", baseURL), Name: "Scarlett", Destination: "Paris", Price: 140.3, Address: "Paris", Longitude: 41.9028, Latitude: 12.4964},
			},
		},
		{
			Name:  "Lyon",
			Image: fmt.Sprintf("%s/images/lyon.jpg",baseURL),
			Activities: []models2.Activity{
				{Image: fmt.Sprintf("%s/images/activities/lyon_opera.jpg", baseURL), Name: "Opéra", Destination: "Lyon", Price: 100.4, Address: "Lyon", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/lyon_bellecour.jpg", baseURL), Name: "Place Bellecour", Destination: "Lyon", Price: 30.4, Address: "Lyon", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/lyon_basilique.jpg", baseURL), Name: "Basilique St-Pierre", Destination: "Lyon", Price: 10.2, Address: "Lyon", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/lyon_mairie.jpg", baseURL), Name: "Mairie", Destination: "Lyon", Price: 30.2, Address: "Lyon", Longitude: 41.9028, Latitude: 12.4964},
			},

			Hotels: []models2.Hotel{
				{Image: fmt.Sprintf("%s/images/hotels/boscolo-lyon.jpeg", baseURL), Name: "Boscolo", Destination: "Lyon", Price: 100.4, Address: "Lyon", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/hotels/intercontinental-lyon.jpeg", baseURL), Name: "Intercontinental", Destination: "Lyon", Price: 130.4, Address: "Lyon", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/hotels/sofitel-lyon.jpeg", baseURL), Name: "Sofitel", Destination: "Lyon", Price: 100.2, Address: "Lyon", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/hotels/villa-florentine-lyon.jpeg", baseURL), Name: "Villa Florentine", Destination: "Lyon", Price: 30.2, Address: "Lyon", Longitude: 41.9028, Latitude: 12.4964},
			},
		},
		{
			Name:  "Nice",
			Image: fmt.Sprintf("%s/images/nice.jpg",baseURL),
			Activities: []models2.Activity{
				{Image: fmt.Sprintf("%s/images/activities/nice_orthodox.jpg", baseURL), Name: "Eglise Orthodoxe", Destination: "Nice", Price: 5.30, Address: "Nice", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/nice_riviera.jpg", baseURL), Name: "Riviera", Destination: "Nice", Price: 20.40, Address: "Nice", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/nice_promenade.jpg", baseURL), Name: "Promenade des Anglais", Destination: "Nice", Price: 30.40, Address: "Nice", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/activities/nice_opera.jpg", baseURL), Name: "Opéra", Destination: "Nice", Price: 100.50, Address: "Nice", Longitude: 41.9028, Latitude: 12.4964},
			},
			Hotels: []models2.Hotel{
				{Image: fmt.Sprintf("%s/images/hotels/piedeau-nice.jpeg", baseURL), Name: "Pied D'Eau", Destination: "Nice", Price: 55.30, Address: "Nice", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/hotels/saintgothard-nice.jpg", baseURL), Name: "Saint Gothard", Destination: "Nice", Price: 60.40, Address: "Nice", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/hotels/yelo-mozart-nice.jpg", baseURL), Name: "Yelo Mozart", Destination: "Nice", Price: 50.40, Address: "Nice", Longitude: 41.9028, Latitude: 12.4964},
				{Image: fmt.Sprintf("%s/images/hotels/nicehome-nice.jpg", baseURL), Name: "Nice Home", Destination: "Nice", Price: 40.50, Address: "Nice", Longitude: 41.9028, Latitude: 12.4964},
			},
		},
	}

	// Enregistrer les instances dans la base de données
	for _, destination := range destinations {
		if err := DB.Create(&destination).Error; err != nil {
			log.Fatalf("could not seed user %v: %v", destination, err)
        		}
	}

	fmt.Println("Seeding completed.")

	}

	func getBaseURL(appConfig *config.Config) string {
        if appConfig.WebURL != "" {
            return appConfig.WebURL
        }
        if appConfig.AndroidURL != "" {
            return appConfig.AndroidURL
        }
        if appConfig.IOSURL != "" {
            return appConfig.IOSURL
        }
        return ""
    }