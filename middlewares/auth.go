package middleware

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// GenerateToken membuat token JWT dengan klaim userID
func GenerateToken(userID int) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token berlaku 24 jam
		Subject:   strconv.Itoa(userID),                               // UserID disimpan sebagai Subject
	}

	// Buat token baru dengan klaim yang sudah ditetapkan
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Tandatangani token dengan kunci rahasia dari environment
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

func ValidateToken(tokenString string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return 0, err
	}

	// Cek apakah klaim token valid dan berisi klaim yang benar
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return 0, err
	}

	// Konversi Subject (userID) kembali ke integer
	userID, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return 0, err
	}

	// Kembalikan userID jika token valid
	return userID, nil
}
