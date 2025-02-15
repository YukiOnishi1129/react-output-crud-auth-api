package auth

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const JWT_TOKEN_EXPIRATION = 60 * 60 * 24

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func CreateToken(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": email,
		"exp": expirationTime.Unix(),
	})
	log.Printf("%s", os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		log.Printf("failed to create token: %v", err)
		return "", err
	}
	return tokenString, nil

}