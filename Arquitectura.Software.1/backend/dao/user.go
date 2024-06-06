package dao

import (
	"backend/clients"
)

func SelectUserByEmail(email string) (*clients.Usuario, error) {
	var user clients.Usuario
	result := clients.DB.Where("email = ?", email).First(&user)
	return &user, result.Error
}
