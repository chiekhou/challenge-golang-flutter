package root

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Deploiement godoc
// @Summary Affiche un message de déploiement réussi
// @Description Renvoie un message JSON indiquant que le déploiement a réussi
// @Tags déploiement
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router / [get]
func Deploiement(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Déploiement réussi inshaàllah!",
	})
}
