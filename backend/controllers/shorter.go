package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/andree37/rlld/db"
	"github.com/andree37/rlld/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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

	database := db.GetDB()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	res, err := database.Collection("short").InsertOne(ctx, bson.M{"long_url": shorter.Url})
	if err != nil {
		log.Printf("something went wrong: %v", err)
		c.JSON(400, gin.H{"error": err})
		return
	}

	id := res.InsertedID

	c.JSON(200, gin.H{"message": short_url, "id": id})
}
