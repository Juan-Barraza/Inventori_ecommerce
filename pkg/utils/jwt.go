package utils

import (
	"fmt"
	modelsgorm "inventory/internal/fiber/infrastructure/persistence/modelsGORM"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("JWT-TOKEN-GO")

func CreateToken(user modelsgorm.UserGormJson, cantDuration int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": string(user.ID),
			"exp":      time.Now().Add(time.Hour * time.Duration(cantDuration)).Unix(),
			"iat":      time.Now().Unix(),
			"email":    user.Email,
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
