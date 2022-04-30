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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShorterController struct{}

func (s ShorterController) DoShort(c *gin.Context) {
	var shorter models.Shorter
	err := c.ShouldBindJSON(&shorter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database := db.GetDB()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	res, err := database.Collection("short").InsertOne(ctx, bson.M{"long_url": shorter.Url})
	if err != nil {
		log.Printf("something went wrong: %v", err)
		c.JSON(400, gin.H{"error": err})
		return
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {

		short_url, err := shorter.Shortens(oid.Hex())
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		}

		c.JSON(200, gin.H{"short_url": short_url, "id": oid})
	} else {
		c.JSON(400, gin.H{"error": "could not create id"})
	}
}

// func (s ShorterController) GetUrlFromId(c *gin.Context) {
// 	var shorter models.Shorter
// 	err := c.ShouldBindJSON(&shorter)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	database := db.GetDB()
// 	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
// 	defer cancel()

// 	err = database.Collection("short").FindOne(ctx, filter).Decode(&result)

// }
