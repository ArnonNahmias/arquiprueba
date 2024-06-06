package controllers

import (
	"backend/domain"
	"backend/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Subscribe(c *gin.Context) {
	var request domain.SubscribeRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.Result{
			Message: "Invalid request: " + err.Error(),
		})
		return
	}

	if err := services.Subscribe(request.UserID, request.CourseID); err != nil {
		c.JSON(http.StatusConflict, domain.Result{
			Message: "Error in subscribe: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, domain.Result{
		Message: fmt.Sprintf("User %d successfully subscribed to course %d", request.UserID, request.CourseID),
	})
}
