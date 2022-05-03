package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeShortURL(t *testing.T) {
	id := 12345
	id2 := 123

	shortID := computeShortURL(id)
	shortID2 := computeShortURL(id2)

	assert.Equal(t, "hnd", shortID)
	assert.Equal(t, "9b", shortID2)
}

func TestComputeID(t *testing.T) {
	shortID := "hnd"
	shortID2 := "9b"

	id := computeID(shortID)
	id2 := computeID(shortID2)

	assert.Equal(t, 12345, id)
	assert.Equal(t, 123, id2)
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
