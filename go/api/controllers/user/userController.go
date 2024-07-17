package user

import (
	"example/hello/internal/initializers"
	"example/hello/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers - Récupère tous les utilisateurs
// @Summary Récupère tous les utilisateurs
// @Description Récupère la liste de tous les utilisateurs
// @Tags User
// @Produce json
// @Success 200 {object} []models.User
// @Failure 500 {object} gin.H "Erreur serveur interne"
// @Security Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer Add access token here)
// @Router /api/users [get]
func GetUsers(c *gin.Context) {
	var users []models.User
	if err := initializers.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur serveur interne"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// CreateUser - Crée un nouvel utilisateur
// @Summary Crée un nouvel utilisateur
// @Description Crée un nouvel utilisateur avec les informations fournies
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.User true "Utilisateur à créer"
// @Success 200 {object} models.User
// @Failure 500 {object} gin.H "Erreur serveur interne"
// @Security Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer Add access token here)
// @Router /api/users  [post]
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := initializers.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur serveur interne"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUser - Récupère un utilisateur par ID
// @Summary Récupère un utilisateur par ID
// @Description Récupère les informations d'un utilisateur spécifique
// @Tags User
// @Produce json
// @Param id path int true "ID de l'utilisateur"
// @Success 200 {object} models.User
// @Failure 404 {object} gin.H "Utilisateur non trouvé"
// @Failure 500 {object} gin.H "Erreur serveur interne"
// @Security Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer Add access token here)
// @Router /api/users/{id} [get]
func GetUser(c *gin.Context) {
	var user models.User
	if err := initializers.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Utilisateur non trouvé"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser - Met à jour un utilisateur par ID
// @Summary Met à jour un utilisateur par ID
// @Description Met à jour les informations d'un utilisateur spécifique
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "ID de l'utilisateur"
// @Param user body models.User true "Utilisateur à mettre à jour"
// @Success 200 {object} models.User
// @Failure 404 {object} gin.H "Utilisateur non trouvé"
// @Failure 500 {object} gin.H "Erreur serveur interne"
// @Security Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer Add access token here)
// @Router /api/users/{id} [put]
func UpdateUser(c *gin.Context) {
	var user models.User
	if err := initializers.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Utilisateur non trouvé"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

// DeleteUser - Supprime un utilisateur par ID
// @Summary Supprime un utilisateur par ID
// @Description Supprime un utilisateur spécifique
// @Tags User
// @Produce json
// @Param id path int true "ID de l'utilisateur"
// @Success 204 {object} nil
// @Failure 404 {object} gin.H "Utilisateur non trouvé"
// @Failure 500 {object} gin.H "Erreur serveur interne"
// @Security Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer Add access token here)
// @Router /api/users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	// Commencer une transaction
	tx := initializers.DB.Begin()

	// Mettre à jour les groupes de voyages pour qu'ils ne référencent plus cet utilisateur
	if err := tx.Model(&models.GroupeVoyage{}).Where("user_id = ?", id).Update("user_id", nil).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la mise à jour des groupes de voyages"})
		return
	}

	// Supprimer l'utilisateur
	if err := tx.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la suppression de l'utilisateur"})
		return
	}

	// Valider la transaction
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la validation de la transaction"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
