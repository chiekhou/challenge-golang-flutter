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

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"example/hello/handlers"
	"example/hello/internal/dao"
	"example/hello/internal/models"
)

func setupGroupeVoyageTestDB() *gorm.DB {
	// Charger les variables d'environnement à partir du fichier .env
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	// Construire le DSN pour se connecter à PostgreSQL
	dsn := "user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=5432 sslmode=disable"

	// Ouvrir la connexion à la base de données
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// Auto migrer les modèles nécessaires
	db.AutoMigrate(&models.GroupeVoyage{}, &models.GroupeMembers{})

	return db
}

func TestListGroupeVoyages(t *testing.T) {
	db := setupGroupeVoyageTestDB()
	groupeVoyageDAO := dao.NewGroupeVoyageDAO(db)
	groupeVoyageHandler := handlers.NewGroupeVoyageHandler(groupeVoyageDAO)
	server := gin.Default()
	server.GET("/api/groupes", groupeVoyageHandler.GetGroupeVoyages)

	// Créer des groupes de voyage de test
	groupe1 := models.GroupeVoyage{
		Budget:        1500.0,
		Roadmap:       "Roadmap for Groupe 1",
		UserID:        1,
		NbPersonnes:   4,
		DateDepart:    time.Now(),
		DateRetour:    time.Now().AddDate(0, 0, 7),
		Nom:           "Groupe 1",
		DestinationID: 1,
	}
	groupe2 := models.GroupeVoyage{
		Budget:        2000.0,
		Roadmap:       "Roadmap for Groupe 2",
		UserID:        2,
		NbPersonnes:   6,
		DateDepart:    time.Now(),
		DateRetour:    time.Now().AddDate(0, 0, 10),
		Nom:           "Groupe 2",
		DestinationID: 2,
	}
	db.Create(&groupe1)
	db.Create(&groupe2)

	req, err := http.NewRequest(http.MethodGet, "/api/groupes", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Groupe 1")
	assert.Contains(t, w.Body.String(), "Groupe 2")
}

func TestAddGroupeVoyage(t *testing.T) {
	db := setupGroupeVoyageTestDB()
	groupeVoyageDAO := dao.NewGroupeVoyageDAO(db)
	groupeVoyageHandler := handlers.NewGroupeVoyageHandler(groupeVoyageDAO)
	server := gin.Default()
	server.POST("/api/groupes", groupeVoyageHandler.CreateGroupeVoyage)

	groupe := models.GroupeVoyage{
		Budget:        2500.0,
		Roadmap:       "Roadmap for New Group",
		UserID:        3,
		NbPersonnes:   8,
		DateDepart:    time.Now(),
		DateRetour:    time.Now().AddDate(0, 0, 14),
		Nom:           "New Group",
		DestinationID: 3,
	}
	body, _ := json.Marshal(groupe)
	req, err := http.NewRequest(http.MethodPost, "/api/groupes", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "New Group")
}

func TestDeleteGroupeVoyage(t *testing.T) {
	db := setupGroupeVoyageTestDB()
	groupeVoyageDAO := dao.NewGroupeVoyageDAO(db)
	groupeVoyageHandler := handlers.NewGroupeVoyageHandler(groupeVoyageDAO)
	server := gin.Default()
	server.DELETE("/api/groupes/delete/:id", groupeVoyageHandler.DeleteGroupeVoyage)

	// Créer un groupe de voyage de test
	groupe := models.GroupeVoyage{
		Budget:        1800.0,
		Roadmap:       "Roadmap for Group to Delete",
		UserID:        4,
		NbPersonnes:   5,
		DateDepart:    time.Now(),
		DateRetour:    time.Now().AddDate(0, 0, 5),
		Nom:           "Group to Delete",
		DestinationID: 4,
	}
	db.Create(&groupe)

	req, err := http.NewRequest(http.MethodDelete, "/api/groupes/delete/"+strconv.Itoa(int(groupe.ID)), nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
