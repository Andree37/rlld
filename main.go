package main

import (
	"database/sql"
	"github.com/andree37/rlld/db"
	"github.com/andree37/rlld/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Something went wrong with the .env")
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
