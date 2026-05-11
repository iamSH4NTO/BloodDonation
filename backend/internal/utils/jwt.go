package utils

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func getSecretKey() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Println("WARNING: JWT_SECRET is not set. Using default insecure key for development.")
		return []byte("dev_secret_key_change_me_in_production")
	}
	return []byte(secret)
}

type Claims struct {
	UserID     string `json:"user_id"`
	Role       string `json:"role"`
	IsVerified bool   `json:"is_verified"`
	jwt.RegisteredClaims
}

func GenerateToken(userID string, role string, isVerified bool) (string, error) {
	jwtKey := getSecretKey()

	expirationTime := time.Now().Add(15 * time.Hour) // Increased for dev convenience
	claims := &Claims{
		UserID:     userID,
		Role:       role,
		IsVerified: isVerified,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (*Claims, error) {
	jwtKey := getSecretKey()

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
