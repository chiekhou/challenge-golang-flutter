package voyage

import (
	"example/hello/api/controllers/requests"
	"example/hello/bin/utils"
	"example/hello/internal/initializers"
	"example/hello/internal/models"
	mailer2 "example/hello/pkg/mailer"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Invitation groupe de voyage
// @Description Envoie un mail d'invitation afin de de rejoindre un groupen de voyage
// @Tags Voyage
// @Accept json
// @Produce json
// @Security Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} gin.H "Invitation envoyée"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 404 {object} gin.H "Bad request"
// @Failure 409 {object} gin.H "Conflict"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /send_invitation [post]
func SendInvitation(c *gin.Context) {
	var emailRequest requests.EmailRequest

	err := c.ShouldBindJSON(&emailRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(emailRequest.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	type emailData struct {
		userFound models.User
		token     string
	}
	var email emailData

	initializers.DB.Where("email=?", emailRequest.Email).First(&userFound)
	if userFound.ID == 0 {
		mailer2.SendGoMail(userFound.Email,
			"Inscription",
			"./pkg/mailer/templates/registry.html",
			email)
	} else {
		mailer2.SendGoMail(userFound.Email,
			"Invitation dans un groupen e voyage",
			"./pkg/mailer/templates/invite.html",
			email)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Email envoyé",
		"token":   token,
	})
}

// @Summary Rejoindre un groupe de voyage
// @Description Envoie un mail d'invitation afin de de rejoindre un groupen de voyage
// @Tags Voyage
// @Accept json
// @Produce json
// @Security Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} gin.H "Invitation envoyée"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 404 {object} gin.H "Bad request"
// @Failure 409 {object} gin.H "Conflict"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /send_invitation [post]
func JoinGroup(c *gin.Context) {

}
