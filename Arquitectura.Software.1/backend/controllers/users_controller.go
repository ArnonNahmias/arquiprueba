package controllers

import (
	"backend/dao"
	"backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var request dao.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	token, userId, role, err := services.Login(request.NombreUsuario, request.Contrasena)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "userId": userId, "role": role})
}
