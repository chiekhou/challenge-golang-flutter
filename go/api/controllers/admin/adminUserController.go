package admin

import (
	"example/hello/internal/initializers"
	"example/hello/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllUsers - Récupère tous les utilisateurs
// @Summary Récupère tous les utilisateurs
// @Description Récupère la liste de tous les utilisateurs
// @Tags Utilisateurs
// @Produce json
// @Success 200 {object} gin.H "Liste des utilisateurs"
// @Router /admin/users [get]
func GetAllUsers(c *gin.Context) {
	var users []models.User
	initializers.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

// DeleteUser - Supprime un utilisateur
// @Summary Supprime un utilisateur
// @Description Supprime un utilisateur de la base de données
// @Tags Utilisateurs
// @Param id path string true "ID de l'utilisateur"
// @Success 200 {object} gin.H "Utilisateur supprimé avec succès"
// @Failure 500 {object} gin.H "Échec de la suppression de l'utilisateur"
// @Router /admin/users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := initializers.DB.Delete(&models.User{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// UpdateUser - Met à jour un utilisateur
// @Summary Met à jour un utilisateur
// @Description Met à jour les informations d'un utilisateur existant
// @Tags Utilisateurs
// @Accept json
// @Produce json
// @Param id path string true "ID de l'utilisateur"
// @Param user body models.User true "Détails de l'utilisateur"
// @Success 200 {object} gin.H "Utilisateur mis à jour avec succès"
// @Failure 400 {object} gin.H "Requête invalide"
// @Failure 404 {object} gin.H "Utilisateur non trouvé"
// @Failure 500 {object} gin.H "Erreur serveur interne"
// @Router /admin/users/{id} [put]
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := initializers.DB.First(&user, id).Error; err != nil {
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
