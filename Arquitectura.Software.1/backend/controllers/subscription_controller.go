package controllers

import (
	"backend/dao"
	"backend/clients"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSubscriptions(c *gin.Context) {
	var subscriptions []dao.Subscription
	clients.DB.Find(&subscriptions)
	c.JSON(http.StatusOK, subscriptions)
}

func CreateSubscription(c *gin.Context) {
	var subscription dao.Subscription
	if err := c.ShouldBindJSON(&subscription); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	clients.DB.Create(&subscription)
	c.JSON(http.StatusCreated, subscription)
}

func DeleteSubscription(c *gin.Context) {
	var subscription dao.Subscription
	id := c.Param("id")
	if err := clients.DB.First(&subscription, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subscription not found"})
		return
	}
	clients.DB.Delete(&subscription)
	c.Status(http.StatusNoContent)
}
