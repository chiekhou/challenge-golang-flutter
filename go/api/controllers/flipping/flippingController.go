package flipping

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FeatureToggle struct {
	ActiveVoyage bool `json:"active_voyage"`
}

var featureToggles = map[string]bool{
	"active_voyage": false,
}

// getFeatureToggle godoc
// @Summary Récupère l'état d'une fonctionnalité
// @Description Récupère l'état (activé ou désactivé) d'une fonctionnalité donnée.
// @Tags Features
// @Produce json
// @Param feature query string true "Nom de la fonctionnalité"
// @Success 200 {object} map[string]bool
// @Failure 400 {string} string "Feature query parameter is missing"
// @Failure 404 {string} string "Feature not found"
// @Router /api/flipping/feature [get]
func GetFeatureToggle(c *gin.Context) {
	feature := c.Query("feature")
	fmt.Println("Feature requested:", feature)

	if feature == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Feature query parameter is missing"})
		return
	}

	enabled, exists := featureToggles[feature]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Feature not found"})
		return
	}

	response := map[string]bool{
		"enabled": enabled,
	}
	c.JSON(http.StatusOK, response)
}

// UpdateFeatureToggle godoc
// @Summary Met à jour l'état d'une fonctionnalité
// @Description Permet de mettre à jour l'état (activé ou désactivé) d'une fonctionnalité spécifique.
// @Tags Features
// @Accept json
// @Produce json
// @Param feature query string true "Nom de la fonctionnalité" Example("create_voyage")
// @Param state body map[string]bool true "Nouvel état de la fonctionnalité"
// @Success 200 {object} map[string]bool
// @Failure 400 {string} string "Feature query parameter is missing"
// @Failure 404 {string} string "Feature not found"
// @Router /api/flipping/feature [put]
func UpdateFeatureToggle(c *gin.Context) {
	feature := c.Query("feature")
	if feature == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Feature query parameter is missing"})
		return
	}

	var state map[string]bool
	if err := c.ShouldBindJSON(&state); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	/* enabled, exists := featureToggles[feature]
	   if !exists {
	       c.JSON(http.StatusNotFound, gin.H{"error": "Feature not found"})
	       return
	   }*/

	newState, ok := state["enabled"]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "State parameter 'enabled' is missing"})
		return
	}

	featureToggles[feature] = newState
	c.JSON(http.StatusOK, gin.H{"enabled": newState})
}
