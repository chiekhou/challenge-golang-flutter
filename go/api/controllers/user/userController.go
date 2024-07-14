package user

import (
	"example/hello/internal/initializers"
	"example/hello/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetGroupsByUser - Récupère tous les groupes de voyage pour un utilisateur spécifique
// @Summary Récupère tous les groupes de voyage pour un utilisateur spécifique
// @Description Récupère la liste de tous les groupes de voyage auxquels un utilisateur est associé
// @Tags Utilisateurs
// @Produce json
// @Param userID path string true "ID de l'utilisateur"
// @Success 200 {object} gin.H "Liste des groupes de voyage"
// @Failure 404 {object} gin.H "Utilisateur non trouvé"
// @Failure 500 {object} gin.H "Erreur serveur interne"
// @Router /user/{userID}/groups [get]
func GetGroupsByUser(c *gin.Context) {
	userID := c.Param("userID")
	var groups []models.GroupeVoyage

	if err := initializers.DB.Where("user_id = ?", userID).Find(&groups).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur serveur interne"})
		return
	}

	if len(groups) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aucun groupe de voyage trouvé pour cet utilisateur"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": groups})
}

// UpdateUserFields - Met à jour les champs d'un utilisateur
// @Summary Met à jour les champs d'un utilisateur
// @Description Met à jour les informations d'un utilisateur existant
// @Tags Utilisateurs
// @Accept json
// @Produce json
// @Param userID path string true "ID de l'utilisateur"
// @Param user body models.User true "Détails de l'utilisateur"
// @Success 200 {object} gin.H "Utilisateur mis à jour avec succès"
// @Failure 400 {object} gin.H "Requête invalide"
// @Failure 404 {object} gin.H "Utilisateur non trouvé"
// @Failure 500 {object} gin.H "Erreur serveur interne"
// @Router /user/{userID} [put]
func UpdateUserFields(c *gin.Context) {
	userID := c.Param("userID")
	var user models.User

	if err := initializers.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Utilisateur non trouvé"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := initializers.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de mettre à jour l'utilisateur"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Utilisateur mis à jour avec succès", "data": user})
}
