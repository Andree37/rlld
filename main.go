package main

import (
	"flag"
	"fmt"
	"os"

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

	database := db.GetDB()

	defer database.Close()
}
