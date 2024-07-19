package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"example/hello/handlers"
	"example/hello/internal/models"

	"github.com/joho/godotenv"
)

func setupVoyageTestDB() *gorm.DB {
	// Charger les variables d'environnement depuis le fichier .env
	err := godotenv.Load("../.env")
	if err != nil {
		panic("Error loading .env file")
	}

	// Récupérer les données de connexion à la base de données depuis les variables d'environnement
	dsn := os.Getenv("DB_URL")

	// Connexion à la base de données PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto migration des modèles nécessaires
	db.AutoMigrate(&models.Voyage{})

	return db
}
func TestListVoyages(t *testing.T) {
	db := setupVoyageTestDB()
	voyageHandler := handlers.NewVoyageHandler(db)
	server := gin.Default()
	server.GET("/api/voyages", voyageHandler.ListVoyages)

	// Create test voyages
	voyage1 := models.Voyage{
		Activities:  []models.Activity{{Name: "Hiking"}, {Name: "Swimming"}},
		DateAller:   time.Now(),
		DateRetour:  time.Now().AddDate(0, 0, 10),
		Destination: "Paris",
		Hotels:      []models.Hotel{{Name: "Hotel 1"}, {Name: "Hotel 2"}},
	}
	voyage2 := models.Voyage{
		Activities:  []models.Activity{{Name: "Sightseeing"}, {Name: "Shopping"}},
		DateAller:   time.Now(),
		DateRetour:  time.Now().AddDate(0, 0, 5),
		Destination: "New York",
		Hotels:      []models.Hotel{{Name: "Hotel A"}, {Name: "Hotel B"}},
	}
	db.Create(&voyage1)
	db.Create(&voyage2)

	req, err := http.NewRequest(http.MethodGet, "/api/voyages", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Paris")
	assert.Contains(t, w.Body.String(), "New York")
}

func TestAddVoyage(t *testing.T) {
	db := setupVoyageTestDB()
	voyageHandler := handlers.NewVoyageHandler(db)
	server := gin.Default()
	server.POST("/api/voyages", voyageHandler.AddVoyage)

	voyage := models.Voyage{
		Activities:  []models.Activity{{Name: "Diving"}, {Name: "Surfing"}},
		DateAller:   time.Now(),
		DateRetour:  time.Now().AddDate(0, 0, 7),
		Destination: "Phuket",
		Hotels:      []models.Hotel{{Name: "Resort 1"}, {Name: "Resort 2"}},
	}
	body, _ := json.Marshal(voyage)
	req, err := http.NewRequest(http.MethodPost, "/api/voyages", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Phuket")
}

func TestDeleteVoyage(t *testing.T) {
	db := setupVoyageTestDB()
	voyageHandler := handlers.NewVoyageHandler(db)
	server := gin.Default()
	server.DELETE("/api/voyages/delete/:id", voyageHandler.DeleteVoyage)

	// Create a test voyage
	voyage := models.Voyage{
		Activities:  []models.Activity{{Name: "Kayaking"}, {Name: "Fishing"}},
		DateAller:   time.Now(),
		DateRetour:  time.Now().AddDate(0, 0, 3),
		Destination: "Alaska",
		Hotels:      []models.Hotel{{Name: "Cabin 1"}},
	}
	db.Create(&voyage)

	req, err := http.NewRequest(http.MethodDelete, "/api/voyages/delete/"+strconv.Itoa(int(voyage.ID)), nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "deleted")
}
