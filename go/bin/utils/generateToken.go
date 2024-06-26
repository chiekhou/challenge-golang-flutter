package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(email string, groupID ...uint) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	if len(groupID) > 0 && groupID[0] != 0 {
		claims["groupID"] = groupID[0]
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("SECRET"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
