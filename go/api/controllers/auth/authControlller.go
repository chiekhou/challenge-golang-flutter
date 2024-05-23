package auth

import (
	"example/hello/bin/utils"
	"example/hello/internal/initializers"
	"example/hello/internal/models"
	mailer2 "example/hello/pkg/mailer"
	jwt2 "github.com/dgrijalva/jwt-go"
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

type EmailRequest struct {
	Email string `form:"email" json:"email" binding:"required"`
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
	mailer2.SendGoMail(user.Email, "Inscription", "./pkg/mailer/templates/registry.html", user)
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

// @Summary Envoie un mail à un user existant afin de réinitialiser son mot de passe
// @Description Envoie un mail de reset de mot de passe
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body EmailRequest true "User data"
// @Success 200 {object} gin.H "Connexion réussie"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 404 {object} gin.H "Bad request"
// @Failure 409 {object} gin.H "Conflict"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /forgotten_password [post]
func MailRecovery(c *gin.Context) {
	var EmailReq EmailRequest

	err := c.ShouldBindJSON(&EmailReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	initializers.DB.Where("email=?", EmailReq.Email).Find(&userFound)

	token, err := utils.GenerateToken(userFound.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resetUrl := "localhost:8080/reset?token=" + token

	mailer2.SendGoMail(userFound.Email,
		"Réinitialiser votre mot de passe",
		"./pkg/mailer/templates/forgottenpass.html",
		userFound)

	c.JSON(200, gin.H{
		"message": "mail envoyé",
		"url":     resetUrl,
	})
}

type ResetPasswordRequest struct {
	Token       string `json:"token" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

// @Summary Réinitialiser le mot de passe
// @Description Permet à l'utilisateur de réinitialiser son mot de passe en utilisant un jeton de réinitialisation valide.
// @Tags Auth
// @Accept json
// @Produce json
// @Param resetPasswordRequest body ResetPasswordRequest true "Données pour réinitialiser le mot de passe"
// @Success 204 "Mot de passe réinitialisé avec succès"
// @Failure 400 {object} gin.H "Token invalide ou expiré"
// @Failure 500 {object} gin.H "Erreur interne du serveur"
// @Router /reset_password [put]
func ResetPassword(c *gin.Context) {
	var ResetPassReq ResetPasswordRequest
	var jwtKey = []byte("votre_clé_secrète")

	if err := c.ShouldBindJSON(&ResetPassReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims := &jwt2.StandardClaims{}
	token, err := jwt2.ParseWithClaims(ResetPassReq.Token, claims, func(token *jwt2.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token invalide ou expiré"})
		return
	}

	var user models.User
	initializers.DB.Where("email = ?", claims.Subject).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Utilisateur non trouvé"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(ResetPassReq.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur de génération de hash de mot de passe"})
		return
	}

	user.Password = string(passwordHash)
	initializers.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"message": "Mot de passe réinitialisé avec succès"})
}
