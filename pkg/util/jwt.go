package util

import (
	"errors"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateToken generate tokens used for auth
func GenerateToken(id, username string, jwtSecret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"username": username,
	})

	tokenString, err := token.SignedString(jwtSecret)

	return tokenString, err
}

// ParseToken parsing token
func ParseToken(tokenString string, jwtSecret []byte) (string, error) {
	token, tokenErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if tokenErr != nil {
		return "", errors.New("noSign")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["id"].(string)

		return userID, nil
	}

	return "", errors.New("noSign")
}
