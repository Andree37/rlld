package db

import (
	"database/sql"
	"fmt"

	"github.com/andree37/rlld/config"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Init() {

	c := config.Getconfig()

	host := c.GetString("db.host")
	port := c.GetInt("db.port")
	user := c.GetString("db.user")
	password := c.GetString("db.password")
	dbname := c.GetString("db.dbname")

	var err error
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to the database!")
}

func GetDB() *sql.DB {
	return db
}
