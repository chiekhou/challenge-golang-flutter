package admin

import (
	"example/hello/internal/initializers"
	"example/hello/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllGroups - Récupère tous les groupes de voyage
// @Summary Récupère tous les groupes de voyage
// @Description Récupère la liste de tous les groupes de voyage
// @Tags Groupes
// @Produce json
// @Success 200 {object} gin.H "Liste des groupes de voyage"
// @Failure 500 {object} gin.H "Erreur serveur interne"
// @Router /admin/groups [get]
func GetAllGroups(c *gin.Context) {
	var groups []models.GroupeVoyage
	if err := initializers.DB.Find(&groups).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de récupérer les groupes de voyage"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": groups})
}

// GetGroup - Récupère un groupe de voyage par ID
// @Summary Récupère un groupe de voyage par ID
// @Description Récupère les détails d'un groupe de voyage spécifique
// @Tags Groupes
// @Produce json
// @Param id path string true "ID du groupe de voyage"
// @Success 200 {object} gin.H "Détails du groupe de voyage"
// @Failure 404 {object} gin.H "Groupe de voyage non trouvé"
// @Router /admin/groups/{id} [get]
func GetGroup(c *gin.Context) {
	id := c.Param("id")
	var group models.GroupeVoyage
	if err := initializers.DB.First(&group, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Groupe de voyage non trouvé"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": group})
}

// CreateGroup - Crée un nouveau groupe de voyage
// @Summary Crée un nouveau groupe de voyage
// @Description Ajoute un nouveau groupe de voyage à la base de données
// @Tags Groupes
// @Accept json
// @Produce json
// @Param group body models.GroupeVoyage true "Détails du groupe de voyage"
// @Success 200 {object} gin.H "Groupe de voyage créé"
// @Failure 400 {object} gin.H "Requête invalide"
// @Failure 500 {object} gin.H "Erreur serveur interne"
// @Router /admin/groups [post]
func CreateGroup(c *gin.Context) {
	var group models.GroupeVoyage
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := initializers.DB.Create(&group).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de créer le groupe de voyage"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": group})
}

// UpdateGroup - Met à jour un groupe de voyage
// @Summary Met à jour un groupe de voyage
// @Description Met à jour les détails d'un groupe de voyage existant
// @Tags Groupes
// @Accept json
// @Produce json
// @Param id path string true "ID du groupe de voyage"
// @Param group body models.GroupeVoyage true "Détails du groupe de voyage"
// @Success 200 {object} gin.H "Groupe de voyage mis à jour"
// @Failure 400 {object} gin.H "Requête invalide"
// @Failure 404 {object} gin.H "Groupe de voyage non trouvé"
// @Failure 500 {object} gin.H "Erreur serveur interne"
// @Router /admin/groups/{id} [put]
func UpdateGroup(c *gin.Context) {
	id := c.Param("id")
	var group models.GroupeVoyage
	if err := initializers.DB.First(&group, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Groupe de voyage non trouvé"})
		return
	}
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := initializers.DB.Save(&group).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de mettre à jour le groupe de voyage"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": group})
}

// DeleteGroup - Supprime un groupe de voyage
// @Summary Supprime un groupe de voyage
// @Description Supprime un groupe de voyage de la base de données
// @Tags Groupes
// @Param id path string true "ID du groupe de voyage"
// @Success 200 {object} gin.H "Groupe de voyage supprimé avec succès"
// @Failure 500 {object} gin.H "Erreur serveur interne"
// @Router /admin/groups/{id} [delete]
func DeleteGroup(c *gin.Context) {
	id := c.Param("id")
	if err := initializers.DB.Delete(&models.GroupeVoyage{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de supprimer le groupe de voyage"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Groupe de voyage supprimé avec succès"})
}
