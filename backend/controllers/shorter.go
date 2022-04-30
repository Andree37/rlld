package controllers

import (
	"net/http"

	"github.com/andree37/rlld/models"
	"github.com/gin-gonic/gin"
)

type ShorterController struct{}

func (s ShorterController) DoShort(c *gin.Context) {

	var shorter models.Shorter
	err := c.ShouldBindJSON(&shorter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	short_url := shorter.Shortens()

	c.JSON(200, gin.H{"message": short_url})
}
