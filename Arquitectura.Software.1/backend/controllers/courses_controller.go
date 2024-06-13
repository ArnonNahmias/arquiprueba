package controllers

import (
	"backend/dao"
	"backend/clients"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCourses(c *gin.Context) {
	var courses []dao.Course
	clients.DB.Find(&courses)
	c.JSON(http.StatusOK, courses)
}

func CreateCourse(c *gin.Context) {
	var course dao.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := clients.DB.Create(&course).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, course)
}

func DeleteCourse(c *gin.Context) {
	var course dao.Course
	id := c.Param("id")
	if err := clients.DB.First(&course, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}
	clients.DB.Delete(&course)
	c.Status(http.StatusNoContent)
}
