package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockDb struct{}

func (db MockDb) GetAll() []Note {
	return nil
}

func TestHomeHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	HomeHandler(res, req)

	assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, `{"message":"Hello"}`, res.Body.String())
}

func TestNotesHandler(t *testing.T) {
	mockDb := MockDb{}
	notesHandle := NotesHandler(mockDb)
	req, _ := http.NewRequest("GET", "/api/v1/notes", nil)
	res := httptest.NewRecorder()
	notesHandle.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)
}
