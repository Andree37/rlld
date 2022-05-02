package models

import (
	"github.com/andree37/rlld/db"
)

type Meme struct {
	Url string `json:"url"`
}

func (m *Meme) GetRandomMeme() error {
	database := db.GetDB()
	query := `SELECT "url" FROM "memes" TABLESAMPLE SYSTEM_ROWS(1);`

	stmt, err := database.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow().Scan(&m.Url)
	if err != nil {
		return err
	}

	return nil
}
