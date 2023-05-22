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

func TestGetHandler(t *testing.T) {
	db := data.NewDatabase()
	db.Set("name", "Seongbin")

	// Test case 1: Key exists
	req := httptest.NewRequest(http.MethodGet, "/get/name", nil)
	res := httptest.NewRecorder()

	h := handler.NewHandler(db)

	router := mux.NewRouter()
	router.HandleFunc("/get/{key}", h.GetHandler).Methods(http.MethodGet)

	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %d, want %d", res.Code, http.StatusOK)
	}

	var response handler.GetResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		t.Errorf("failed to decode response body: %s", err.Error())
	}

	expectedKey := "name"
	expectedValue := "Seongbin"
	if response.Key != expectedKey {
		t.Errorf("unexpected response key: got %s, want %s", response.Key, expectedKey)
	}
	if response.Value != expectedValue {
		t.Errorf("unexpected response value: got %s, want %s", response.Value, expectedValue)
	}

	// Test case 2: Key does not exist
	req = httptest.NewRequest(http.MethodGet, "/get/nonexistent", nil)
	res = httptest.NewRecorder()

	router.ServeHTTP(res, req)

	if res.Code != http.StatusNotFound {
		t.Errorf("unexpected status code: got %d, want %d", res.Code, http.StatusNotFound)
	}

	var errorResponse handler.ErrorResponse
	if err := json.NewDecoder(res.Body).Decode(&errorResponse); err != nil {
		t.Errorf("failed to decode error response body: %s", err.Error())
	}

	expectedErrorMessage := "Key 'nonexistent' not found"
	if errorResponse.Message != expectedErrorMessage {
		t.Errorf("unexpected error response message: got %s, want %s", errorResponse.Message, expectedErrorMessage)
	}
}
