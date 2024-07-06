package groupeVoyage

import (
	"example/hello/api/controllers/requests"
	"example/hello/bin/utils"
	"example/hello/internal/initializers"
	"example/hello/internal/models"
	mailer2 "example/hello/pkg/mailer"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Créé un groupe de groupeVoyage
// @Description Permet aux user de créé un groupe de groupeVoyage
// @Tags Groupe Voyage
// @Accept json
// @Produce json
// @Security Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer Add access token here)
// @Param budget body requests.GroupRequest true "Mise à jour du budget"
// @Success 200 {object} gin.H "Groupe créé"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 404 {object} gin.H "Bad request"
// @Failure 409 {object} gin.H "Conflict"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /create_group [post]
func CreateGroup(c *gin.Context) {
	user, exists := c.Get("currentUser")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var groupData requests.GroupRequest
	err := c.ShouldBindJSON(&groupData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	group := models.GroupeVoyage{
		Budget: groupData.Budget,
		UserID: user.(models.User).ID,
		Nom:    groupData.Nom,
	}
	err = initializers.DB.Create(&group).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de créer le groupe de groupeVoyage"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Groupe de groupeVoyage créé avec succès"})
}

// @Summary Voir un groupe de voyage
// @Description Permet de voir le contenu du groupe de voyage pour celui qui l'a créé mais également les membres du groupe
// @Tags Groupe Voyage
// @Accept json
// @Produce json
// @Security Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer Add access token here)
// @Success 200 {object} []models.GroupeVoyage "Liste groupe de voyage"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 404 {object} gin.H "Bad request"
// @Failure 409 {object} gin.H "Conflict"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /groupes/my_groups [get]
func GetMyGroups(c *gin.Context) {
	user, exists := c.Get("currentUser")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	currentUser := user.(models.User)

	var groups []models.GroupeVoyage

	// Obtenir les groupes par l'id du user et précharger les membres
	if err := initializers.DB.Where("user_id = ?", currentUser.ID).Preload("Members").Find(&groups).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Ajouter les groupes où l'utilisateur est membre
	var memberGroups []models.GroupeVoyage
	subQuery := initializers.DB.Table("groupe_members").Select("groupe_voyage_id").Where("user_id = ?", currentUser.ID)
	if err := initializers.DB.Where("id IN (?)", subQuery).Preload("Members").Find(&memberGroups).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	groups = append(groups, memberGroups...)
	c.JSON(http.StatusOK, groups)
}

// @Summary Récupérer un groupe de voyage par ID
// @Description Récupère un groupe de voyage par son ID si l'utilisateur est le créateur ou un membre du groupe
// @Tags Groupe Voyage
// @Accept json
// @Produce json
// @Security Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param group_id path uint true "ID du groupe de voyage"
// @Success 200 {object} models.GroupeVoyage "Détails du groupe de voyage"
// @Failure 401 {object} gin.H "Unauthorized"
// @Failure 403 {object} gin.H "Forbidden"
// @Failure 404 {object} gin.H "Not found"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /groupes/{group_id} [get]
func GetGroupById(c *gin.Context) {
	user, exists := c.Get("currentUser")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	currentUser := user.(models.User)

	groupID, err := strconv.ParseUint(c.Param("group_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID du groupe invalide"})
		return
	}

	var group models.GroupeVoyage
	if err := initializers.DB.Preload("Members").First(&group, groupID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Groupe non trouvé"})
		return
	}

	// Vérifiez si l'utilisateur est le créateur ou un membre du groupe
	if group.UserID != currentUser.ID {
		var member models.GroupeMembers
		if err := initializers.DB.Where("groupe_voyage_id = ? AND user_id = ?", groupID, currentUser.ID).First(&member).Error; err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "Accès interdit"})
			return
		}
	}

	c.JSON(http.StatusOK, group)
}

// @Summary Met à jour le budget d'un groupe de groupeVoyage
// @Description Met à jour le budget d'un groupe de groupeVoyage spécifique en utilisant son ID
// @Tags Groupe Voyage
// @Accept json
// @Produce json
// @Security Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param group_id path string true "ID du groupe de groupeVoyage"
// @Param budget body requests.UpdateBudgetRequest true "Mise à jour du budget"
// @Success 200 {object} gin.H "Budget mis à jour avec succès"
// @Failure 400 {object} gin.H "Bad Request"
// @Failure 404 {object} gin.H "Groupe de groupeVoyage non trouvé"
// @Failure 500 {object} gin.H"Impossible de mettre à jour le budget"
// @Router /groupes/{group_id}/update_budget [put]
func UpdateBudget(c *gin.Context) {
	var budgetRequest requests.UpdateBudgetRequest
	if err := c.ShouldBindJSON(&budgetRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	groupID := c.Param("id")
	var group models.GroupeVoyage
	if err := initializers.DB.First(&group, groupID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Groupe de Voyage non trouvé"})
		return
	}

	group.Budget = budgetRequest.Budget
	if err := initializers.DB.Save(&group).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de mettre à jour le budget"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Budget mis à jour avec succès"})
}

// @Summary Invitation groupe de groupeVoyage
// @Description Envoie un mail d'invitation afin de de rejoindre un groupen de groupeVoyage
// @Tags Groupe Voyage
// @Accept json
// @Produce json
// @Security Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param group_id path string true "ID du groupe de groupeVoyage"
// @Param budget body requests.EmailRequest true "Mise à jour du budget"
// @Success 200 {object} gin.H "Invitation envoyée"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 404 {object} gin.H "Bad request"
// @Failure 409 {object} gin.H "Conflict"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /groupes/{group_id}/send_invitation [post]
func SendInvitation(c *gin.Context) {
	var emailRequest requests.EmailRequest
	groupID, err := strconv.ParseUint(c.Param("group_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID du groupe invalide"})
		return
	}

	if err := c.ShouldBindJSON(&emailRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Requête invalide"})
		return
	}

	// Vérifiez si le groupe existe
	var group models.GroupeVoyage
	if err := initializers.DB.First(&group, groupID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Groupe non trouvé"})
		return
	}

	// Vérifiez si l'utilisateur existe
	var userFound models.User
	if err := initializers.DB.Where("email = ?", emailRequest.Email).First(&userFound).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Utilisateur non trouvé"})
		return
	}

	token, err := utils.GenerateInvitationToken(uint(groupID), userFound.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de générer le token d'invitation"})
		return
	}

	initializers.DB.Where("email=?", emailRequest.Email).First(&userFound)
	if userFound.ID == 0 {
		mailer2.SendGoMail(userFound.Email,
			"Inscription",
			"./pkg/mailer/templates/forgottenpass.html",
			token)
	} else {
		mailer2.SendGoMail(userFound.Email,
			"Invitation dans un groupen e groupeVoyage",
			"./pkg/mailer/templates/invite.html",
			token)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Email envoyé",
		"token":   token,
	})

}

// @Summary Rejoindre un groupe de groupeVoyage
// @Description Permet à un utilisateur de rejoindre un groupe de groupeVoyage en utilisant un token d'invitation
// @Tags Groupe Voyage
// @Accept json
// @Produce json
// @Security Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param group_id path string true "ID du groupe de groupeVoyage"
// @Param token query string true "Token d'invitation"
// @Success 200 {object} gin.H "Vous avez rejoint le groupe avec succès"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 404 {object} gin.H "Not found"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /groupes/{group_id}/join [post]
func Join(c *gin.Context) {
	groupID := c.Param("group_id")
	token := c.Query("token")

	// Valider le token et obtenir l'ID de l'utilisateur
	userID, err := utils.ValidateInvitationToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token invalide"})
		return
	}

	// Vérifiez si le groupe existe
	var group models.GroupeVoyage
	if err := initializers.DB.First(&group, groupID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Groupe de Voyage non trouvé"})
		return
	}

	// Vérifiez si l'utilisateur existe
	var user models.User
	if err := initializers.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Utilisateur non trouvé"})
		return
	}

	// Ajouter l'utilisateur au groupe en utilisant le modèle GroupMember
	groupMember := models.GroupeMembers{
		GroupeVoyageID: group.ID,
		UserID:         user.ID,
	}
	if err := initializers.DB.Create(&groupMember).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de rejoindre le groupe de voyage"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vous avez rejoint le groupe de voyage avec succès"})
}
