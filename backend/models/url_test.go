package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase10ToBase62(t *testing.T) {
	id := int64(12345)
	id2 := int64(123)
	id3 := int64(1231231233213123)

	shortID := base10ToBase62(id)
	shortID2 := base10ToBase62(id2)
	shortID3 := base10ToBase62(id3)

	assert.Equal(t, "DNH", shortID)
	assert.Equal(t, "B9", shortID2)
	assert.Equal(t, "FnmgAry3B", shortID3)
}

func TestBase62ToBase10(t *testing.T) {
	shortID := "DNH"
	shortID2 := "B9"
	shortID3 := "FnmgAry3B"

	id, _ := base62ToBase10(shortID)
	id2, _ := base62ToBase10(shortID2)
	id3, _ := base62ToBase10(shortID3)

	assert.Equal(t, int64(12345), id)
	assert.Equal(t, int64(123), id2)
	assert.Equal(t, int64(1231231233213123), id3)
}

func TestIsValidURL(t *testing.T) {
	url := new(URL)
	url.OriginalUrl = "https://www.google.com"

	valid, _ := url.IsValidURL()
	assert.Equal(t, true, valid)

	url.OriginalUrl = "www.google.com"

	valid, _ = url.IsValidURL()
	assert.Equal(t, true, valid)

	url.OriginalUrl = "google.com"

	valid, _ = url.IsValidURL()
	assert.Equal(t, true, valid)

	url.OriginalUrl = "potato"

	valid, _ = url.IsValidURL()
	assert.Equal(t, false, valid)

	url.OriginalUrl = ""

	valid, _ = url.IsValidURL()
	assert.Equal(t, false, valid)

	url.OriginalUrl = "https://google.com"

	valid, _ = url.IsValidURL()
	assert.Equal(t, true, valid)
}
