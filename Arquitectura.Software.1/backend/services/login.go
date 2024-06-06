package services

import (
	"backend/clients"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

var jwtSecret = []byte("your_secret_key")

func generateJWT(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func Login(email string, password string) (string, error) {
	if strings.TrimSpace(email) == "" {
		return "", errors.New("email is required")
	}

	if strings.TrimSpace(password) == "" {
		return "", errors.New("password is required")
	}

	hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))

	userDAO, err := SelectUserByEmail(email)
	if err != nil {
		return "", fmt.Errorf("error getting user from DB: %w", err)
	}

	if hash != userDAO.Password {
		return "", fmt.Errorf("invalid credentials")
	}

	token, err := generateJWT(email)
	if err != nil {
		return "", fmt.Errorf("error generating JWT token: %w", err)
	}

	return token, nil
}

func SelectUserByEmail(email string) (*clients.Usuario, error) {
	var user clients.Usuario
	result := clients.DB.Where("email = ?", email).First(&user)
	return &user, result.Error
}
