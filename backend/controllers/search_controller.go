package controllers

import (
	"backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Search(c *gin.Context) {
	query := c.Query("query")
	courses, err := services.Search(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, courses)
}

func SearchByID(c *gin.Context) {
	id := c.Param("id")
	course, err := services.SearchByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, course)
}
