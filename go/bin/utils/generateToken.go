package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("secret")

func GenerateToken(email string) (string, error) {
	expiration := time.Now().Add(time.Hour * 2).Unix()

	claims := &jwt.StandardClaims{
		Subject:   email,
		ExpiresAt: expiration,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
