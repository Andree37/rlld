package db

import (
	"context"
	"log"
	"time"

	"github.com/andree37/rlld/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var db *mongo.Client

func Init() {
	var err error
	c := config.Getconfig()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err = mongo.Connect(ctx, options.Client().ApplyURI(c.GetString("db.uri")))
	if err != nil {
		log.Fatal("Could not find database")
	}
	err = db.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Could not connect successfully to the database")
	}
}

func GetDB() *mongo.Database {
	return db.Database("rlld")
}

func GetClient() *mongo.Client {
	return db
}
