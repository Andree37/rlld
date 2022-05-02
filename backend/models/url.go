package models

import (
	"strings"

	"github.com/andree37/rlld/db"
)

type URL struct {
	ID          int     `json:"id"`
	OriginalUrl string  `json:"original_url"`
	ShortID     string  `json:"short_id"`
	MemePrctg   float64 `json:"meme_prctg"`
}

func computeShortURL(id int) string {
	m := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var shortURL []string

	for id > 0 {
		shortURL = append(shortURL, string(m[id%62]))
		id /= 62
	}

	return strings.Join(shortURL, "")
}

func computeID(shortID string) int {
	id := 0
	for i := len(shortID) - 1; i >= 0; i-- {
		val_i := int(shortID[i])
		if val_i >= int('a') && val_i <= int('z') {
			id = id*62 + val_i - int('a')
		} else if val_i >= int('A') && val_i <= int('Z') {
			id = id*62 + val_i - int('Z') + 26
		} else {
			id = id*62 + val_i - int('0') + 52
		}
	}

	return id
}

func (u *URL) TranslateToShortID() error {
	database := db.GetDB()
	query := `INSERT INTO tiny_urls ("original_url", "meme_percentage") values ($1, $2) RETURNING id`

	stmt, err := database.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var insertedID int
	err = stmt.QueryRow(u.OriginalUrl, u.MemePrctg).Scan(&insertedID)
	if err != nil {
		return err
	}

	computedID := computeShortURL(insertedID)

	u.ShortID = computedID

	return nil
}

func (u *URL) GetURL() error {
	database := db.GetDB()
	query := `SELECT "original_url", "meme_percentage" FROM "tiny_urls" WHERE "id" = $1`

	// get the databaseID
	id := computeID(u.ShortID)

	stmt, err := database.Prepare(query)
	if err != nil {
		return err
	}
	err = stmt.QueryRow(id).Scan(&u.OriginalUrl, &u.MemePrctg)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return nil
}
