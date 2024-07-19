package handlers

import (
	"net/http"
	"strconv"

	"example/hello/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// VoyageHandler représente le gestionnaire de voyages
type VoyageHandler struct {
	db *gorm.DB
}

// NewVoyageHandler initialise un nouveau gestionnaire de voyages avec une instance de DB
func NewVoyageHandler(db *gorm.DB) *VoyageHandler {
	return &VoyageHandler{db}
}

// ListVoyages récupère la liste des voyages
func (h *VoyageHandler) ListVoyages(c *gin.Context) {
	var voyages []models.Voyage
	h.db.Find(&voyages)
	c.JSON(http.StatusOK, voyages)
}

// AddVoyage ajoute un nouveau voyage
func (h *VoyageHandler) AddVoyage(c *gin.Context) {
	var voyage models.Voyage
	if err := c.BindJSON(&voyage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.db.Create(&voyage)
	c.JSON(http.StatusCreated, voyage)
}

// DeleteVoyage supprime un voyage par ID
func (h *VoyageHandler) DeleteVoyage(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid voyage ID"})
		return
	}

	var voyage models.Voyage
	result := h.db.First(&voyage, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Voyage not found"})
		return
	}

	h.db.Delete(&voyage)
	c.JSON(http.StatusOK, gin.H{"message": "Voyage deleted"})
}
