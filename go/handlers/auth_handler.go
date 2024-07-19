// Package handlers définit les gestionnaires HTTP pour l'authentification
package handlers

import (
	"example/hello/internal/dao"
	"example/hello/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// AuthHandler représente le gestionnaire d'authentification
type AuthHandler struct {
	userDAO *dao.UserDAO
}

// NewAuthHandler crée une nouvelle instance de AuthHandler
func NewAuthHandler(userDAO *dao.UserDAO) *AuthHandler {
	return &AuthHandler{
		userDAO: userDAO,
	}
}

// Signup gère l'inscription d'un nouvel utilisateur
func (h *AuthHandler) Signup(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user.Password = string(hashedPassword)

	if err := h.userDAO.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}

// Login gère la connexion d'un utilisateur existant
func (h *AuthHandler) Login(c *gin.Context) {
	var loginCredentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&loginCredentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input"})
		return
	}

	user, err := h.userDAO.GetUserByEmail(loginCredentials.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginCredentials.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// TODO: Generate and return JWT token

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
}

func (h *AuthHandler) ForgotPassword(c *gin.Context) {
	var input struct {
		Email string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userDAO.GetUserByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	token, err := h.userDAO.GeneratePasswordResetToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate reset token"})
		return
	}

	// Send the reset token via email (email sending logic not shown here)
	// e.g., sendPasswordResetEmail(user.Email, token)

	c.JSON(http.StatusOK, gin.H{"message": "Password reset email sent", "token": token})
}

func (h *AuthHandler) ResetPassword(c *gin.Context) {
	type ResetPasswordRequest struct {
		Email           string `json:"email" binding:"required,email"`
		ResetToken      string `json:"reset_token" binding:"required"`
		NewPassword     string `json:"new_password" binding:"required,min=8"`
		ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=NewPassword"`
	}

	var req ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userDAO.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Verify the reset token
	validToken, err := h.userDAO.VerifyPasswordResetToken(user.ID, req.ResetToken)
	if err != nil || !validToken {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or expired reset token"})
		return
	}

	// Hash the new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Update the user's password in the database
	if err := h.userDAO.UpdatePassword(user.ID, string(hashedPassword)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	// Clear the session, JWT token, or any authentication mechanism you're using
	// For example, if you're using JWT, you may need to expire the token
	// or remove it from the client's cookies

	// Logic for clearing session or JWT token (implement according to your setup)
	// For JWT token, you might implement a blacklist or expiration mechanism

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

func (h *AuthHandler) Profile(c *gin.Context) {
	// Extract user information from the request context or JWT token
	// Assuming you have a middleware or authentication logic that verifies the user
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Cast the user to your user model
	u, ok := user.(*models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user information"})
		return
	}

	// Prepare the response JSON with user details
	response := gin.H{
		"id":         u.ID,
		"first_name": u.FirstName,
		"last_name":  u.LastName,
		"photo":      u.Photo,
		"username":   u.Username,
		"email":      u.Email,
		"address":    u.Address,
		"role_id":    u.RoleID,
		"groupe_voyage": func() []uint {
			var ids []uint
			for _, gv := range u.GroupeVoyage {
				ids = append(ids, gv.ID)
			}
			return ids
		}(),
		// Add more fields as per your user model
	}

	c.JSON(http.StatusOK, response)
}
