package auth

import (
	"example/hello/internal/initializers"
	"example/hello/internal/models"
	mailer2 "example/hello/pkg/mailer"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

// Requête personnalisée que l'on va binder sur le model de "User" enregistré en base de données
type SignupRequest struct {
	FirstName string `form:"first_name" json:"first_name" binding:"required"`
	LastName  string `form:"last_name" json:"last_name" binding:"required"`
	Address   string `form:"address" json:"address" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

// Requête pour pouvoir se loguer
type LoginRequest struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// @Summary Allow you to register as a new User
// @Description Create a new user with the provided information
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body SignupRequest true "User data"
// @Success 201 {object} SignupRequest "User created"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 404 {object} gin.H "Bad request"
// @Failure 409 {object} gin.H "Conflict"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /Signup [post]
func Signup(c *gin.Context) {
	var signupReq SignupRequest

	if err := c.ShouldBindJSON(&signupReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	initializers.DB.Where("email = ?", signupReq.Email).Find(&userFound)
	if userFound.ID != 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "email already used"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(signupReq.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		FirstName: signupReq.FirstName,
		LastName:  signupReq.LastName,
		Username:  signupReq.Username,
		Password:  string(passwordHash),
		Email:     signupReq.Email,
		Address:   signupReq.Address,
	}
	mailer2.SendGoMail(user.Email, "Inscription", "./pkg/mailer/templates/registry.html")
	initializers.DB.Create(&user)
	c.JSON(http.StatusCreated, gin.H{"data": user})
}

// @Summary Allow you to log and have an JWT Token
// @Description login to the app
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body LoginRequest true "User data"
// @Success 200 {object} gin.H "Connexion réussie"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 404 {object} gin.H "Bad request"
// @Failure 409 {object} gin.H "Conflict"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /login [post]
func Login(c *gin.Context) {
	var loginReq LoginRequest

	err := c.ShouldBindJSON(&loginReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	initializers.DB.Where("username=?", loginReq.Email).Find(&userFound)

	if userFound.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(loginReq.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userFound.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

//func Logout(c *gin.Context) {}

// @Summary Récupère le profil de l'utilisateur actuellement connecté
// @Description Retourne les informations du profil de l'utilisateur connecté
// @Tags Auth
// @Accept json
// @Produce json
// @Security Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} gin.H "Success"
// @Failure 401 {object} gin.H "Unauthorized"
// @Router /profile [get]
func UserProfile(c *gin.Context) {
	user, _ := c.Get("currentUser")
	c.JSON(http.StatusOK, gin.H{"user": user})
}

//TODO mot de passe oublié
