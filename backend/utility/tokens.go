package utility

import (
	"time"
"fmt"
	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken generates a JWT token
func GenerateToken(email, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

// DecodingToken decodes a JWT token
func DecodingToken(tokenString, secretKey string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}