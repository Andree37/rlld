package controllers

import (
	"net/http"

	"github.com/andree37/rlld/models"
	"github.com/gin-gonic/gin"
)

type PingController struct{}

func (p PingController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func (p PingController) DoPing(c *gin.Context) {
	var json models.Ping
	err := c.ShouldBindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "pong", "written": json.Command})
}
