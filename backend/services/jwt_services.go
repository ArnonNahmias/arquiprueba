package services

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("yourSecretKey")

type JWTService interface {
	GenerateToken(userId uint64) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type authCustomClaims struct {
	UserId uint64 `json:"user_id"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtServices{
		secretKey: string(secretKey),
		issuer:    "yourIssuer",
	}
}

func (service *jwtServices) GenerateToken(userId uint64) (string, error) {
	claims := &authCustomClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(service.secretKey))
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})
}
