package middleware

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("123abc")

func GenerateToken(userID uint) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		Subject:   strconv.Itoa(int(userID)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return 0, err
	}

	userID, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return 0, err
	}

	return uint(userID), nil
}
