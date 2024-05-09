package auth

import (
	"example/hello/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 201 {object} models.User "User created"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 409 {object} gin.H "Conflict"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /Signup [post]
func Signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBind(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

//TODO Connexion
//TODO mot de passe oublié
//TODO
