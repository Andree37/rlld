package controllers

import (
	"math/rand"
	"net/http"

	"github.com/andree37/rlld/models"
	"github.com/gin-gonic/gin"
)

const GAG_SHUFFLE_URL = "https://www.9gag.com/shuffle"
const RICK_ROLL = "https://www.youtube.com/watch?v=dQw4w9WgXcQ"

type URLController struct{}

func (s URLController) Tinify(c *gin.Context) {
	var URLModel models.URL
	err := c.ShouldBindJSON(&URLModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect json"})
		return
	}
	// check if the URL is valid
	valid, err := URLModel.IsValidURL()
	if err != nil || !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect URL type"})
		return
	}

	// compute the short url
	err = URLModel.TranslateToShortID()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not insert URL"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"short_id": URLModel.ShortID})
}

func (s URLController) GetURLFromID(c *gin.Context) {
	var URLModel models.URL

	URLModel.ShortID = c.Param("short_id")
	err := URLModel.GetURL()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r := rand.Float64()
	if r > URLModel.MemePrctg {
		c.Redirect(http.StatusSeeOther, URLModel.OriginalUrl)
	} else {
		// fetch a meme from the 9gag randomizer, ty 9gag for making life easy :)
		//c.Redirect(http.StatusSeeOther, GAG_SHUFFLE_URL)
		c.Redirect(http.StatusSeeOther, RICK_ROLL)
	}
}
