package controllers

import (
	"backend/dao"
	"backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var loginRequest dao.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, userID, userType, err := services.Login(loginRequest.NombreUsuario, loginRequest.Contrasena)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":  token,
		"userId": userID,
		"type":   userType,
	})
}
