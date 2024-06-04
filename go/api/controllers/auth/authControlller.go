package auth

import (
	"example/hello/internal/initializers"
	"example/hello/internal/models"
	mailer2 "example/hello/pkg/mailer"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	Data models.User `json:"data"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

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
// @Success 201 {object} UserResponse "User created"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 409 {object} ErrorResponse "Conflict"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /Signup [post]
func Signup(c *gin.Context) {
	var signupReq SignupRequest

	if err := c.ShouldBindJSON(&signupReq); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	var userFound models.User
	initializers.DB.Where("email = ?", signupReq.Email).Find(&userFound)
	if userFound.ID != 0 {
		c.JSON(http.StatusConflict, ErrorResponse{Error: "email already used"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(signupReq.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
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
	initializers.DB.Create(&user)
	mailer2.SendGoMail(user.Email, "Inscription", "./pkg/mailer/templates/registry.html", user)

	c.JSON(http.StatusCreated, UserResponse{Data: user})
}

// @Summary Allow you to log in and get a JWT Token
// @Description Login to the app
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body LoginRequest true "User data"
// @Success 200 {object} TokenResponse "Successful login"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /login [post]
func Login(c *gin.Context) {
	var loginReq LoginRequest

	err := c.ShouldBindJSON(&loginReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	var userFound models.User
	initializers.DB.Where("email = ?", loginReq.Email).Find(&userFound)

	if userFound.ID == 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(loginReq.Password)); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid password"})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userFound.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, TokenResponse{Token: token})
}

// @Summary Get the profile of the currently logged-in user
// @Description Returns the profile information of the logged-in user
// @Tags Auth
// @Accept json
// @Produce json
// @Security Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} UserResponse "Success"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Router /profile [get]
func UserProfile(c *gin.Context) {
	user, _ := c.Get("currentUser")
	c.JSON(http.StatusOK, UserResponse{Data: user.(models.User)})
}

// @Summary Send an email to an existing user to reset their password
// @Description Sends a password reset email
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body EmailRequest true "User data"
// @Success 200 {object} MessageResponse "Email sent"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 404 {object} ErrorResponse "User not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /forgotten_password [post]
func MailRecovery(c *gin.Context) {
	var emailReq EmailRequest

	err := c.ShouldBindJSON(&emailReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	var userFound models.User
	initializers.DB.Where("email = ?", emailReq.Email).Find(&userFound)
	if userFound.ID == 0 {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "user not found"})
		return
	}

	mailer2.SendGoMail(userFound.Email,
		"Réinitialiser votre mot de passe",
		"./pkg/mailer/templates/forgottenpass.html",
		userFound)

	c.JSON(http.StatusOK, MessageResponse{Message: "mail envoyé"})
}
