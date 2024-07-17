package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"example/hello/api/controllers/destinations"
	"example/hello/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// setupRouter initialise le routeur Gin avec les routes nécessaires pour les tests.
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/api/destinations", destinations.GetDestinations)
	r.GET("/api/destinations/:id", destinations.GetDestination)
	r.POST("/api/destinations", destinations.CreateDestination)
	r.PATCH("/api/destinations/update/:id", destinations.UpdateDestination)
	r.DELETE("/api/destinations/delete/:id", destinations.DeleteDestination)
	r.POST("/api/destinations/:id/activity", destinations.CreateActivityDestination)

	return r
}

// setupDatabase initialise une base de données SQLite en mémoire pour les tests.
func setupDatabase() *gorm.DB {
	db, _ = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.Destination{}, &models.Activity{}, &models.DestinationActivity{})
	return db
}

func TestGetDestinations(t *testing.T) {
	db := setupDatabase()
	defer db.Migrator().DropTable(&models.Destination{}, &models.Activity{}, &models.DestinationActivity{})

	// Créer quelques données de test
	db.Create(&models.Destination{Name: "Test Destination 1"})
	db.Create(&models.Destination{Name: "Test Destination 2"})

	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/destinations", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response destinations.DestinationResponse
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Len(t, response.Data.([]models.Destination), 2)
}

func TestGetDestination(t *testing.T) {
	db := setupDatabase()
	defer db.Migrator().DropTable(&models.Destination{}, &models.Activity{}, &models.DestinationActivity{})

	// Créer une donnée de test
	destination := models.Destination{Name: "Test Destination"}
	db.Create(&destination)

	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/destinations/"+strconv.Itoa(int(destination.ID)), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.Destination
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, destination.Name, response.Name)
}

func TestCreateDestination(t *testing.T) {
	db := setupDatabase()
	defer db.Migrator().DropTable(&models.Destination{}, &models.Activity{}, &models.DestinationActivity{})

	router := setupRouter()

	destination := map[string]interface{}{
		"name":  "New Destination",
		"image": "image_url",
	}

	jsonValue, _ := json.Marshal(destination)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/destinations", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response destinations.DestinationResponse
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, destination["name"], response.Data.(models.Destination).Name)
}

func TestUpdateDestination(t *testing.T) {
	db := setupDatabase()
	defer db.Migrator().DropTable(&models.Destination{}, &models.Activity{}, &models.DestinationActivity{})

	// Créer une donnée de test
	destination := models.Destination{Name: "Old Destination"}
	db.Create(&destination)

	router := setupRouter()

	updatedDestination := map[string]interface{}{
		"name":  "Updated Destination",
		"image": "updated_image_url",
	}

	jsonValue, _ := json.Marshal(updatedDestination)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/api/destinations/update/"+strconv.Itoa(int(destination.ID)), bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response destinations.DestinationResponse
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, updatedDestination["name"], response.Data.(models.Destination).Name)
}

func TestDeleteDestination(t *testing.T) {
	db := setupDatabase()
	defer db.Migrator().DropTable(&models.Destination{}, &models.Activity{}, &models.DestinationActivity{})

	// Créer une donnée de test
	destination := models.Destination{Name: "Test Destination"}
	db.Create(&destination)

	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/destinations/delete/"+strconv.Itoa(int(destination.ID)), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response destinations.SuccessResponse
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.True(t, response.Data)
}

func TestCreateActivityDestination(t *testing.T) {
	db := setupDatabase()
	defer db.Migrator().DropTable(&models.Destination{}, &models.Activity{}, &models.DestinationActivity{})

	// Créer une destination de test
	destination := models.Destination{Name: "Test Destination"}
	db.Create(&destination)

	router := setupRouter()

	activity := map[string]interface{}{
		"name": "Test Activity",
	}

	jsonValue, _ := json.Marshal(activity)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/destinations/"+strconv.Itoa(int(destination.ID))+"/activity", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response destinations.DestinationResponse
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, destination.Name, response.Data.(models.Destination).Name)
}

func TestVerifyActivityName(t *testing.T) {
	db := setupDatabase()
	defer db.Migrator().DropTable(&models.Destination{}, &models.Activity{}, &models.DestinationActivity{})

	// Créer une destination et une activité de test
	destination := models.Destination{Name: "Test Destination"}
	activity := models.Activity{Name: "Test Activity"}
	db.Create(&destination)
	db.Create(&activity)
	db.Create(&models.DestinationActivity{DestinationID: destination.ID, ActivityID: activity.ID})

	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/destination/"+strconv.Itoa(int(destination.ID))+"/activities/verify/"+activity.Name, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]models.DestinationActivity
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, destination.ID, response["destinationActivity"].DestinationID)
}
