package middlewares

import (
	"example/hello/internal/initializers"
	"example/hello/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CheckAdmin - Vérifie si l'utilisateur est un administrateur
// @Summary Vérifie si l'utilisateur est un administrateur
// @Description Vérifie le rôle de l'utilisateur pour s'assurer qu'il est administrateur
// @Tags Middleware
// @Security Bearer
// @Param Authorization header string true "Insérez votre jeton d'accès" default(Bearer <Add access token here>)
// @Success 200 {object} gin.H "Autorisé"
// @Failure 401 {object} gin.H "Non autorisé"
// @Failure 403 {object} gin.H "Interdit"
// @Router /middleware/checkadmin [get]
func CheckAdmin(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	var user models.User
	initializers.DB.First(&user, userID)
	if user.RoleID != 1 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		c.Abort()
		return
	}

	c.Next()
}
