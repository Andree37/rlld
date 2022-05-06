package models

import (
	"regexp"
	"strings"

	"github.com/andree37/rlld/db"
	"github.com/jxskiss/base62"
)

var blacklisted = []string{"ADD HERE THE URL"}

type URL struct {
	ID          int     `json:"id"`
	OriginalUrl string  `json:"original_url"`
	ShortID     string  `json:"short_id"`
	MemePrctg   float64 `json:"meme_prctg"`
}

func base10ToBase62(num int64) string {
	return string(base62.FormatInt(num))
}

func base62ToBase10(s string) (int64, error) {
	return base62.ParseInt([]byte(s))
}

func (u *URL) IsValidURL() (bool, error) {
	//check if it has a https, if not add it and then test

	v := len(strings.Split(u.OriginalUrl, "https://www")) == 2

	if !v {
		u.OriginalUrl = "https://www." + u.OriginalUrl
	}

	// check for blacklisted URLs
	for _, v := range blacklisted {
		if v == u.OriginalUrl {
			return false, nil
		}
	}

	r := `https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`

	match, err := regexp.MatchString(r, u.OriginalUrl)
	if err != nil {
		return false, err
	}

	return match, nil
}

func (u *URL) TranslateToShortID() error {
	database := db.GetDB()
	query := `INSERT INTO tiny_urls ("original_url", "meme_percentage") values ($1, $2) RETURNING id`

	stmt, err := database.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var insertedID int64
	err = stmt.QueryRow(u.OriginalUrl, u.MemePrctg).Scan(&insertedID)
	if err != nil {
		return err
	}

	computedID := base10ToBase62(insertedID)

	u.ShortID = computedID

	return nil
}

func (u *URL) GetURL() error {
	database := db.GetDB()
	query := `SELECT "original_url", "meme_percentage" FROM "tiny_urls" WHERE "id" = $1`

	// get the databaseID
	id, err := base62ToBase10(u.ShortID)
	if err != nil {
		return err
	}

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
