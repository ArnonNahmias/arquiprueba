package controllers

import (
	"backend/dao"
	"backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetCourses(c *gin.Context) {
	courses, err := services.GetCourses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, courses)
}

func CreateCourse(c *gin.Context) {
	var course dao.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newCourse, err := services.CreateCourse(course)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newCourse)
}

func DeleteCourse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}
	if err := services.DeleteCourse(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
