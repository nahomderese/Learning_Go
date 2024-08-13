package infrastructure

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(username, id string) (string, error) {

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"id":       id,
		"exp":      time.Now().AddDate(0, 0, 1).Unix(),
	}).SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return token, nil
}
