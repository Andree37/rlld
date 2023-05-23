package main

import (
	"database/sql"
	"fmt"

	"github.com/andree37/rlld/db"
	"github.com/andree37/rlld/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Something went wrong with the .env")
	}

	db.Init()
	server.Init()

	database := db.GetDB()

	defer func(database *sql.DB) {
		err := database.Close()
		if err != nil {
		}
	}(database)
}
