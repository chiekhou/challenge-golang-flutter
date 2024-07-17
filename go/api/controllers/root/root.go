package root

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Deploiement(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Déploiement réussi inshaàllah!",
	})
}
