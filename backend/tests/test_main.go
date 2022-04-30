package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andree37/rlld/server"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := server.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.Bytes()), &response)

	if err != nil {
		t.Errorf("something went wrong with the json unmarshal: %v", err)
	}

	// get the value from the response
	// this should be a struct
	value, exists := response["message"]

	assert.Equal(t, true, exists)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", value)

}
