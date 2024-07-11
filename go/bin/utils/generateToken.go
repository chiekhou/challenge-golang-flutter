package utils

import (
<<<<<<< HEAD
<<<<<<< HEAD
=======
	"encoding/base64"
	"fmt"
>>>>>>> origin/feature/merge_voyage
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func GenerateToken(email string, ID uint) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"email": email,
		"ID":    ID,
		"exp":   now.Add(time.Hour * 72).Unix(),
		"iat":   now.Unix(),
	}

	secret := []byte(os.Getenv("SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

<<<<<<< HEAD
	tokenString, err := token.SignedString([]byte("SECRET"))
=======
	"encoding/base64"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func GenerateToken(email string, ID uint) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"email": email,
		"ID":    ID,
		"exp":   now.Add(time.Hour * 72).Unix(),
		"iat":   now.Unix(),
	}

	secret := []byte(os.Getenv("SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret)
>>>>>>> origin/feature/merge_voyage
=======
	tokenString, err := token.SignedString(secret)
>>>>>>> origin/feature/merge_voyage
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
<<<<<<< HEAD
<<<<<<< HEAD
=======
=======
>>>>>>> origin/feature/merge_voyage

var jwtKey = []byte(os.Getenv("SECRET"))

type InvitationClaims struct {
	GroupID uint `json:"group_id"`
	UserID  uint `json:"user_id"`
	jwt.StandardClaims
}

// Générer un token d'invitation
func GenerateInvitationToken(groupID, userID uint) (string, error) {
	data := fmt.Sprintf("%d:%d:%d", groupID, userID, time.Now().Unix())
	return base64.URLEncoding.EncodeToString([]byte(data)), nil
}

// Valider un token d'invitation
func ValidateInvitationToken(token string) (uint, error) {
	decoded, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return 0, err
	}

	var groupID, userID, timestamp uint
	_, err = fmt.Sscanf(string(decoded), "%d:%d:%d", &groupID, &userID, &timestamp)
	if err != nil {
		return 0, err
	}

	// Ajoutez ici toute validation supplémentaire, comme vérifier si le token a expiré

	return userID, nil
}
<<<<<<< HEAD
>>>>>>> origin/feature/merge_voyage
=======
>>>>>>> origin/feature/merge_voyage
