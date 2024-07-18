package tests

import (
	"bytes"
	"encoding/json"
	"example/hello/handlers"
	"example/hello/internal/dao"
	"example/hello/internal/models"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	dsn := os.Getenv("DB_URL")
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

func TestForgotPassword(t *testing.T) {
	db := setupTestDB()
	userDAO := dao.NewUserDAO(db)
	authHandler := handlers.NewAuthHandler(userDAO)
	server := gin.Default()
	server.POST("/forgot_password", authHandler.ForgotPassword)

	// Creating a test user
	user := models.User{Email: "test@example.com", Password: "password123"}
	db.Create(&user)

	forgotPasswordRequest := map[string]string{"email": "test@example.com"}
	body, _ := json.Marshal(forgotPasswordRequest)
	req, err := http.NewRequest(http.MethodPost, "/forgot_password", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "password reset email sent")
}

func TestLogout(t *testing.T) {
	db := setupTestDB()
	userDAO := dao.NewUserDAO(db)
	authHandler := handlers.NewAuthHandler(userDAO)
	server := gin.Default()
	server.POST("/logout", authHandler.Logout)

	// Assuming authentication is handled with tokens, simulate a logged-in user
	token := "mock_valid_token" // Replace with actual token generation logic if needed
	req, err := http.NewRequest(http.MethodPost, "/logout", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "logged out successfully")
}

func TestProfile(t *testing.T) {
	db := setupTestDB()
	userDAO := dao.NewUserDAO(db)
	authHandler := handlers.NewAuthHandler(userDAO)
	server := gin.Default()
	server.GET("/profile", authHandler.Profile)

	// Creating a test user
	user := models.User{Email: "test@example.com", Password: "password123"}
	db.Create(&user)

	// Assuming authentication is handled with tokens, simulate a logged-in user
	token := "mock_valid_token" // Replace with actual token generation logic if needed
	req, err := http.NewRequest(http.MethodGet, "/profile", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "test@example.com")
}

func TestResetPassword(t *testing.T) {
	db := setupTestDB()
	userDAO := dao.NewUserDAO(db)
	authHandler := handlers.NewAuthHandler(userDAO)
	server := gin.Default()
	server.POST("/reset_password", authHandler.ResetPassword)

	// Creating a test user
	user := models.User{Email: "test@example.com", Password: "password123"}
	db.Create(&user)

	resetPasswordRequest := map[string]string{"email": "test@example.com", "new_password": "new_password123"}
	body, _ := json.Marshal(resetPasswordRequest)
	req, err := http.NewRequest(http.MethodPost, "/reset_password", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "password reset successful")
}
