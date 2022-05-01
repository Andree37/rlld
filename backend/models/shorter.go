package models

import (
	"sort"
	"strings"
)

type Shorter struct {
	Url      string `json:"url"`
	ShortUrl string `json:"short_url"`
}

// Shorter logic here
func IDToShortID(id int) string {
	m := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var shortURL []string

	for id > 0 {
		shortURL = append(shortURL, string(m[id%62]))
		id /= 62
	}

	sort.Sort(sort.Reverse(sort.StringSlice(shortURL)))

	url := strings.Join(shortURL, "")

	return url
}

func ShortIDToID(shortURL string) int {
	id := 0
	for _, v := range shortURL {
		val_i := int(v)
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
