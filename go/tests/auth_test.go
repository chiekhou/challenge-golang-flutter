package tests

import (
	"bytes"
	"encoding/json"
	"example/hello/internal/dao"
	"example/hello/internal/handlers"
	"example/hello/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	dsn := "user=yourusername password=yourpassword dbname=yourdbname port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.User{})
	return db
}

func TestSignup(t *testing.T) {
	db := setupTestDB()
	userDAO := dao.NewUserDAO(db)
	authHandler := handlers.NewAuthHandler(userDAO)
	server := gin.Default()
	server.POST("/signup", authHandler.Signup)

	user := models.User{Email: "test@example.com", Password: "password123"}
	body, _ := json.Marshal(user)
	req, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "test@example.com")
}

func TestLogin(t *testing.T) {
	db := setupTestDB()
	userDAO := dao.NewUserDAO(db)
	authHandler := handlers.NewAuthHandler(userDAO)
	server := gin.Default()
	server.POST("/login", authHandler.Login)

	// Creating a test user
	user := models.User{Email: "test@example.com", Password: "password123"}
	db.Create(&user)

	loginCredentials := map[string]string{"email": "test@example.com", "password": "password123"}
	body, _ := json.Marshal(loginCredentials)
	req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "token")
}

// Other tests for forgotten_password, logout, profile, reset_password...
