package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHomeHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	HomeHandler(res, req)

	assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, `{"message":"Hello"}`, res.Body.String())
}
