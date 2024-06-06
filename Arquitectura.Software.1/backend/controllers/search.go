package controllers

import (
	"backend/domain"
	"backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Search(c *gin.Context) {
	query := strings.TrimSpace(c.Query("query"))
	results, err := services.Search(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Result{
			Message: "Error in search: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, domain.SearchResponse{
		Results: results,
	})
}
