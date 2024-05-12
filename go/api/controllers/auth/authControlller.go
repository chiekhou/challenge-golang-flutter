package auth

import (
	"example/hello/internal/initializers"
	"example/hello/internal/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
// @Failure 404 {object} gin.H "Bad request"
// @Failure 409 {object} gin.H "Conflict"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /Signup [post]
func Signup(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	initializers.DB.Where("email = ?", user.Email).Find(&userFound)
	if userFound.ID != 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "email already used"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user = models.User{
		Username: user.Username,
		Password: string(passwordHash),
	}

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

//TODO Connexion
//TODO mot de passe oublié
//TODO
