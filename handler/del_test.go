package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/0rcastra/Orca/handler"
	"github.com/0rcastra/Orca/internal/data"
	"github.com/gorilla/mux"
)

func TestDelHandler(t *testing.T) {
	db := data.NewDatabase()
	db.Set("name", "Seongbin")

	req := httptest.NewRequest(http.MethodDelete, "/del/name", nil)

	res := httptest.NewRecorder()

	h := handler.NewHandler(db)

	router := mux.NewRouter()
	router.HandleFunc("/del/{key}", h.DelHandler).Methods(http.MethodDelete)

	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %d, want %d", res.Code, http.StatusOK)
	}

	var response handler.DelResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		t.Errorf("failed to decode response body: %s", err.Error())
	}

	expectedKey := "name"
	expectedDeleted := true
	if response.Key != expectedKey {
		t.Errorf("unexpected response key: got %s, want %s", response.Key, expectedKey)
	}
	if response.Deleted != expectedDeleted {
		t.Errorf("unexpected response deleted value: got %v, want %v", response.Deleted, expectedDeleted)
	}
}
