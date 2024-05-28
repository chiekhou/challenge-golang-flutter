package voyage

import (
	"example/hello/api/controllers/requests"
	"example/hello/bin/utils"
	"example/hello/internal/initializers"
	"example/hello/internal/models"
	mailer2 "example/hello/pkg/mailer"
	"net/http"

	"github.com/gin-gonic/gin"
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
			"./pkg/mailer/templates/forgottenpass.html",
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
func JoinGroupr(c *gin.Context) {

}

func CreateGroup(c *gin.Context) {
	// Extraire les données de la requête
	var groupData requests.GroupRequest
	err := c.ShouldBindJSON(&groupData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	group := models.Group{

		Budget:        groupData.Budget,
		Roadmap:       groupData.Roadmap,
		UserID:        groupData.UserID,
		NbPersonnes:   groupData.NbPersonnes,
		DateDepart:    groupData.DateDepart,
		DateRetour:    groupData.DateRetour,
		Nom:           groupData.Nom,
		DestinationID: groupData.DestinationID,
	}
	err = initializers.DB.Create(&group).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de créer le groupe de voyage"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Groupe de voyage créé avec succès"})
}

func UpdateBudget(c *gin.Context) {
	var budgetRequest requests.UpdateBudgetRequest
	if err := c.ShouldBindJSON(&budgetRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	groupID := c.Param("id")
	var group models.GroupeVoyage
	if err := initializers.DB.First(&group, groupID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Groupe de voyage non trouvé"})
		return
	}

	group.Budget = budgetRequest.Budget
	if err := initializers.DB.Save(&group).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de mettre à jour le budget"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Budget mis à jour avec succès"})
}

func JoinGroup(c *gin.Context) {
	groupID := c.Param("group_id")

	var group models.GroupeVoyage
	if err := initializers.DB.First(&group, groupID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Groupe de voyage non trouvé"})
		return
	}

	userID := c.MustGet("user_id").(uint)
	group.UserID = userID

	if err := initializers.DB.Save(&group).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de rejoindre le groupe de voyage"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vous avez rejoint le groupe de voyage avec succès"})
}
