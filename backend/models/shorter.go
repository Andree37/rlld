package models

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
)

type Shorter struct {
	Url string `json:"url" binding:"required"`
}

// Shorter logic here
func (s *Shorter) Shortens(id string) (string, error) {

	return idToShortUrl(id)
}

func idToShortUrl(id string) (string, error) {
	m := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortURL := ""

	// convert id to an integer[]
	var ints []string
	for _, c := range id {
		ints = append(ints, fmt.Sprintf("%v", int(c)))
	}
	x := strings.Join(ints, "")
	integer_id := new(big.Int)
	integer_id, ok := integer_id.SetString(x, 10)
	if !ok {
		return "", errors.New("could not turn string into integer")
	}

	mod_int := new(big.Int)

	for integer_id.Cmp(big.NewInt(0)) > 0 {
		shortURL += string(m[mod_int.Mod(integer_id, big.NewInt(62)).Int64()])
		integer_id.Div(integer_id, big.NewInt(62))
	}

	//reverse the shortURL
	result := ""
	for _, v := range shortURL {
		result = string(v) + result
	}

	return result, nil
}

func shortURLToID(shortURL string) int {
	id := 0
	for _, v := range shortURL {
		val_i := int(v)
		if val_i >= int('a') && val_i <= int('Z') {
			id = id*62 + val_i - int('a')
		} else if val_i >= int('A') && val_i <= int('Z') {
			id = id*62 + val_i - int('Z') + 26
		} else {
			id = id*62 + val_i - int('0') + 52
		}
	}

	return id
}
