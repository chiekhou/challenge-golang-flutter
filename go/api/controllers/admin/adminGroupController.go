package admin

import (
	"example/hello/internal/initializers"
	"example/hello/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllGroups - Récupère tous les groupes de voyage
func GetAllGroups(c *gin.Context) {
	var groups []models.GroupeVoyage
	if err := initializers.DB.Find(&groups).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de récupérer les groupes de voyage"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": groups})
}

// GetGroup - Récupère un groupe de voyage par ID
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
func DeleteGroup(c *gin.Context) {
	id := c.Param("id")
	if err := initializers.DB.Delete(&models.GroupeVoyage{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de supprimer le groupe de voyage"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Groupe de voyage supprimé avec succès"})
}
