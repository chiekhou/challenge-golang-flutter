package tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"

	// "strconv"
	"testing"

	"example/hello/handlers"
	"example/hello/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupUserTestDB() (*gorm.DB, *sql.DB) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get *sql.DB from *gorm.DB: %v", err)
	}
	return db, sqlDB
}

func TestCreateUser(t *testing.T) {
	db, sqlDB := setupUserTestDB()
	defer sqlDB.Close()

	userHandler := handlers.NewUserHandler(db)
	server := gin.Default()
	server.POST("/api/users", userHandler.CreateUser)

	newUser := models.User{
		Username: "johndoe",
		Email:    "johndoe@example.com",
		Password: "password123",
	}

	body, _ := json.Marshal(newUser)
	req, err := http.NewRequest(http.MethodPost, "/api/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetUsers(t *testing.T) {
	db, sqlDB := setupUserTestDB()
	defer sqlDB.Close()

	userHandler := handlers.NewUserHandler(db)
	server := gin.Default()
	server.GET("/api/users", userHandler.GetUsers)

	req, err := http.NewRequest(http.MethodGet, "/api/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetUserByID(t *testing.T) {
	db, sqlDB := setupUserTestDB()
	defer sqlDB.Close()

	userHandler := handlers.NewUserHandler(db)
	server := gin.Default()
	server.GET("/api/users/:id", userHandler.GetUserByID)

	req, err := http.NewRequest(http.MethodGet, "/api/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateUser(t *testing.T) {
	db, sqlDB := setupUserTestDB()
	defer sqlDB.Close()

	userHandler := handlers.NewUserHandler(db)
	server := gin.Default()
	server.PUT("/api/users/:id", userHandler.UpdateUser)

	updatedUser := models.User{
		Username: "updatedjohndoe",
		Email:    "updatedjohndoe@example.com",
		Password: "updatedpassword123",
	}

	body, _ := json.Marshal(updatedUser)
	req, err := http.NewRequest(http.MethodPut, "/api/users/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteUser(t *testing.T) {
	db, sqlDB := setupUserTestDB()
	defer sqlDB.Close()

	userHandler := handlers.NewUserHandler(db)
	server := gin.Default()
	server.DELETE("/api/users/:id", userHandler.DeleteUser)

	req, err := http.NewRequest(http.MethodDelete, "/api/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
