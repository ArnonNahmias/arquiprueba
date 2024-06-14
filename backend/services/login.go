package services

import (
	"backend/clients"
	"backend/dao"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	UserID   uint   `json:"userId"`
	jwt.StandardClaims
}

func Login(username, password string) (string, uint, string, error) {
	var user dao.Usuario
	if err := clients.DB.Where("nombre_usuario = ? AND contrasena = ?", username, password).First(&user).Error; err != nil {
		return "", 0, "", errors.New("invalid username or password")
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		UserID:   uint(user.IdUsuario),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", 0, "", err
	}

	return tokenString, uint(user.IdUsuario), user.Tipo, nil
}
