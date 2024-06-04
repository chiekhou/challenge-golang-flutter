package activity

import (
	"example/hello/internal/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	Data models.User `json:"data"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

// UploadImage godoc
//
//	@Summary		Upload image
//	@Description	Upload file
//	@Tags			Activity
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			image	formData	file	true	"account image"
//	@Success		200		{object}	models.Activity
//
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 409 {object} ErrorResponse "Conflict"
// @Failure 500 {object} ErrorResponse "Internal server error"
//
//	@Router			/api/activity/images [post]
func UploadImage(c *gin.Context) {
	file, err := c.FormFile("activity")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
		return
	}

	// Create the directory if it doesn't exist
	uploadPath := "./images/activities"
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		os.MkdirAll(uploadPath, os.ModePerm)
	}

	// Enregistrer le fichier sur le serveur
	if err := c.SaveUploadedFile(file, "./assets/images"+file.Filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Chemin public pour accéder à l'image
	publicPath := "http://localhost:8080/images/" + file.Filename

	c.JSON(http.StatusOK, gin.H{"url": publicPath})
}
