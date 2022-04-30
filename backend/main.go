package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/andree37/rlld/config"
	"github.com/andree37/rlld/db"
	"github.com/andree37/rlld/server"
)

func main() {

	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.Init()
	server.Init()

	dbClient := db.GetClient()

	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		if err := dbClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}
