package controllers

import (
	"fmt"
	"net/http"

	"github.com/andree37/rlld/db"
	"github.com/andree37/rlld/models"
	"github.com/gin-gonic/gin"
)

type ShorterController struct{}

func (s ShorterController) URLToShortURL(c *gin.Context) {
	var shorter models.Shorter
	err := c.ShouldBindJSON(&shorter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database := db.GetDB()
	query := `INSERT INTO tiny_urls ("long_url") values ($1) RETURNING id`

	stmt, err := database.Prepare(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	var insertedID int
	err = stmt.QueryRow(shorter.Url).Scan(&insertedID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// compute the short url
	fmt.Printf("inserted id: %v", insertedID)
	short_id := models.IDToShortID(insertedID)

	c.JSON(http.StatusOK, gin.H{"short_id": short_id})
}

func (s ShorterController) ShortURLToURL(c *gin.Context) {
	short_id := c.Param("short_id")

	database := db.GetDB()
	query := `SELECT "long_url" FROM "tiny_urls" WHERE "id" = $1`

	stmt, err := database.Prepare(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get the databaseID
	id := models.ShortIDToID(short_id)

	var long_url string
	err = stmt.QueryRow(id).Scan(&long_url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	//c.JSON(http.StatusOK, gin.H{"long_url": long_url})

	fmt.Printf("long url: %v\n", long_url)

	c.Redirect(http.StatusMovedPermanently, long_url)

}
