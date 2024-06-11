package utils

import (
	"github.com/golang-jwt/jwt/v4"
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
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
