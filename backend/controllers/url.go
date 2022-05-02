package controllers

import (
	"net/http"

	"github.com/andree37/rlld/models"
	"github.com/gin-gonic/gin"
)

type URLController struct{}

func (s URLController) Tinify(c *gin.Context) {
	var URLModel models.URL
	err := c.ShouldBindJSON(&URLModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// compute the short url
	err = URLModel.TranslateToShortID()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"short_url": URLModel.ShortUrl})
}

func (s URLController) GetURLFromID(c *gin.Context) {
	var URLModel models.URL

	URLModel.ShortID = c.Param("short_id")
	err := URLModel.GetURL()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusMovedPermanently, URLModel.OriginalUrl)

}
