package groupeVoyage

import (
	"errors"
	"example/hello/api/controllers/requests"
	"example/hello/bin/utils"
	"example/hello/internal/initializers"
	"example/hello/internal/models"
	mailer2 "example/hello/pkg/mailer"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// @Summary Créé un groupe de voyage
// @Description Permet aux utilisateurs de créer un groupe de voyage
// @Tags Groupe Voyage
// @Accept json
// @Produce json
// @Security Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer Add access token here)
// @Param group body requests.GroupRequest true "Données du groupe"
// @Success 200 {object} gin.H "Groupe créé"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 401 {object} gin.H "Unauthorized"
// @Failure 404 {object} gin.H "Voyage non trouvé"
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

	var voyage models.Voyage
	if groupData.VoyageID != 0 {
		if err := initializers.DB.Preload("Activities").Preload("Hotels").First(&voyage, groupData.VoyageID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Voyage non trouvé"})
			return
		} else if voyage.UserId != user.(models.User).ID {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Vous n'êtes pas autorisé à utiliser ce voyage"})
			return
		}
	}

	group := models.GroupeVoyage{
		Budget: groupData.Budget,
		UserID: user.(models.User).ID,
		Nom:    groupData.Nom,
		Voyage: voyage,
	}

	if err := initializers.DB.Create(&group).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de créer le groupe de voyage"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Groupe de voyage créé avec succès",
		"groupe":  group,
	})
}

// @Summary Suppression d'un groupe de voyage
// @Description Permet de supprimer un groupe de voyage
// @Tags Groupe Voyage
// @Accept json
// @Produce json
// @Security Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer Add access token here)
// @Param group_id path int true "ID du groupe de voyage"
// @Success 200 {object} gin.H "Groupe de voyage supprimé avec succès"
// @Success 204 "Groupe de voyage supprimé avec succès, aucune réponse"
// @Failure 400 {object} gin.H "Requête incorrecte"
// @Failure 404 {object} gin.H "Groupe de voyage non trouvé"
// @Failure 409 {object} gin.H "Conflit lors de la suppression du groupe de voyage"
// @Failure 500 {object} gin.H "Erreur interne du serveur"
// @Router /groupes/{group_id}/delete_group [delete]
func DeleteGroup(c *gin.Context) {
	user, exists := c.Get("currentUser")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	currentUser := user.(models.User)
	groupID := c.Param("group_id")

	var group models.GroupeVoyage
	if err := initializers.DB.Where("id = ?", groupID).First(&group).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Groupe de voyage non trouvé"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	if group.UserID != currentUser.ID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Vous n'êtes pas autorisé à supprimer ce groupe"})
		return
	}

	// Supprimer tous les membres associés au groupe de voyage
	if err := initializers.DB.Model(&group).Association("Members").Clear(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la suppression des membres du groupe"})
		return
	}

	var voyage models.Voyage
	if err := initializers.DB.Model(&models.Voyage{}).Where("id = ?", voyage.ID).Update("groupe_voyage_id", nil).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de mettre à jour le groupe de voyage dans le voyage"})
		return
	}

	if err := initializers.DB.Delete(&group).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de supprimer le groupe de voyage"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Groupe supprimé avec succès",
	})
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

	if err := initializers.DB.Where("user_id = ?", currentUser.ID).Preload("Members").Find(&groups).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

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
	user, exists := c.Get("currentUser")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}

	currentUser := user.(models.User)

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

	if currentUser.ID != group.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Vous n'êtes pas autorisé à effectuer cette opération"})
		return
	}

	// Vérifiez si l'utilisateur existe
	var userFound models.User
	if err := initializers.DB.Where("email = ?", emailRequest.Email).First(&userFound).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Utilisateur non trouvé"})
		return
	}

	//Vérifier si l'user ne fait pas déjà partie du groupe
	var memberFound models.GroupeMembers
	if err := initializers.DB.Where("groupe_voyage_id = ? AND user_id = ?", groupID, userFound.ID).First(&memberFound).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "L'utilisateur fait déjà partie du groupe"})
		return
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur interne du serveur"})
		return
	}

	token, err := utils.GenerateInvitationToken(uint(groupID), userFound.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de générer le token d'invitation"})
		return
	}

	invitationURL := fmt.Sprintf("http://10.0.2.2:8080/groupes/%d/join?token=%s", groupID, token)

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
			map[string]interface{}{"invitationURL": invitationURL})
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
// @Router /groupes/{group_id}/join [get]
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

	// Vérifiez si l'utilisateur est déjà membre du groupe
	var existingGroupMember models.GroupeMembers
	if err := initializers.DB.Where("groupe_voyage_id = ? AND user_id = ?", group.ID, user.ID).First(&existingGroupMember).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Vous êtes déjà membre de ce groupe de voyage"})
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

// @Summary Suppression d'un membre d'un groupe de voyage
// @Description Permet de supprimer un membre d'un groupe de voyage
// @Tags Groupe Voyage
// @Accept json
// @Produce json
// @Security Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer Add access token here)
// @Param group_id path int true "ID du groupe de voyage"
// @Param member_id path int true "ID du membre à supprimer"
// @Success 200 {object} gin.H "Membre supprimé du groupe avec succès"
// @Failure 400 {object} gin.H "Requête incorrecte"
// @Failure 401 {object} gin.H "Non autorisé"
// @Failure 404 {object} gin.H "Groupe de voyage ou membre non trouvé"
// @Failure 409 {object} gin.H "Conflit lors de la suppression du membre"
// @Failure 500 {object} gin.H "Erreur interne du serveur"
// @Router /groupes/{group_id}/member/{member_id}/delete_member [delete]
func DeleteGroupMember(c *gin.Context) {
	user, exists := c.Get("currentUser")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	currentUser := user.(models.User)
	groupID := c.Param("group_id")

	var group models.GroupeVoyage
	if err := initializers.DB.Where("id = ?", groupID).First(&group).Error; err != nil {
		// Vérifier si l'erreur est de type record not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Groupe de voyage non trouvé"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Vérifier si l'utilisateur courant est le propriétaire du groupe
	if group.UserID != currentUser.ID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Vous n'êtes pas autorisé à supprimer ce groupe"})
		return
	}

	//Si le membre existe dans le groupe
	memberID := c.Param("member_id")
	var member models.GroupeMembers
	if err := initializers.DB.Where("groupe_voyage_id = ? AND user_id = ?", groupID, memberID).First(&member).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Membre non trouvé"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := initializers.DB.Delete(&member).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Membre supprimé avec succés"})
}
