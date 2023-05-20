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

	req1 := httptest.NewRequest(http.MethodDelete, "/del/name", nil)
	res1 := httptest.NewRecorder()
	h1 := handler.NewHandler(db)
	router1 := mux.NewRouter()
	router1.HandleFunc("/del/{key}", h1.DelHandler).Methods(http.MethodDelete)
	router1.ServeHTTP(res1, req1)

	if res1.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %d, want %d", res1.Code, http.StatusOK)
	}

	var response1 handler.DelResponse
	if err := json.NewDecoder(res1.Body).Decode(&response1); err != nil {
		t.Errorf("failed to decode response body: %s", err.Error())
	}

	expectedKey1 := "name"
	expectedDeleted1 := true
	if response1.Key != expectedKey1 {
		t.Errorf("unexpected response key: got %s, want %s", response1.Key, expectedKey1)
	}
	if response1.Deleted != expectedDeleted1 {
		t.Errorf("unexpected response deleted value: got %v, want %v", response1.Deleted, expectedDeleted1)
	}

	req2 := httptest.NewRequest(http.MethodDelete, "/del/nonexistent", nil)
	res2 := httptest.NewRecorder()
	h2 := handler.NewHandler(db)
	router2 := mux.NewRouter()
	router2.HandleFunc("/del/{key}", h2.DelHandler).Methods(http.MethodDelete)
	router2.ServeHTTP(res2, req2)

	if res2.Code != http.StatusNotFound {
		t.Errorf("unexpected status code: got %d, want %d", res2.Code, http.StatusNotFound)
	}

	var response2 handler.DelResponse
	if err := json.NewDecoder(res2.Body).Decode(&response2); err != nil {
		t.Errorf("failed to decode response body: %s", err.Error())
	}

	expectedKey2 := "nonexistent"
	expectedDeleted2 := false
	if response2.Key != expectedKey2 {
		t.Errorf("unexpected response key: got %s, want %s", response2.Key, expectedKey2)
	}
	if response2.Deleted != expectedDeleted2 {
		t.Errorf("unexpected response deleted value: got %v, want %v", response2.Deleted, expectedDeleted2)
	}
}
