package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Controller) HealthCheck(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"message": "health check",
	})
	return
}
