package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

func ParseToken(tokenString string) (string, uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-secret-key"), nil
	})

	if err != nil {
		return "", 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := claims["email"].(string)
		var groupID uint
		if gid, ok := claims["groupID"]; ok {
			groupID = uint(gid.(float64))
		}
		return email, groupID, nil
	}

	return "", 0, errors.New("token invalide")
}
