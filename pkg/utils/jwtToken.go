package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id int, userName string) (string, error) {
	claims := jwt.MapClaims{
		"id":       id,
		"userName": userName,
		"exp":      time.Now().Add(10 * time.Minute).Unix(), // convert to Unix timestamp
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	

	// Use a real secret key in production!
	secretKey := []byte("privateKey")

	return token.SignedString(secretKey)
}
