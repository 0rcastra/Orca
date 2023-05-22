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

func TestDecrHandler(t *testing.T) {
	db := data.NewDatabase()
	db.Set("age", "23")

	req := httptest.NewRequest(http.MethodPost, "/decr/age", nil)
	res := httptest.NewRecorder()

	h := handler.NewHandler(db)

	router := mux.NewRouter()
	router.HandleFunc("/decr/{key}", h.DecrHandler).Methods(http.MethodPost)

	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %d, want %d", res.Code, http.StatusOK)
	}

	var response handler.DecrResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		t.Errorf("failed to decode response body: %s", err.Error())
	}

	expectedKey := "age"
	expectedValue := 22
	if response.Key != expectedKey {
		t.Errorf("unexpected response key: got %s, want %s", response.Key, expectedKey)
	}
	if response.Value != expectedValue {
		t.Errorf("unexpected response value: got %d, want %d", response.Value, expectedValue)
	}
}

func TestDecrHandler_Error(t *testing.T) {
	db := data.NewDatabase()
	db.Set("name", "Seongbin")

	req := httptest.NewRequest(http.MethodPost, "/decr/name", nil)
	res := httptest.NewRecorder()

	h := handler.NewHandler(db)

	router := mux.NewRouter()
	router.HandleFunc("/decr/{key}", h.DecrHandler).Methods(http.MethodPost)

	router.ServeHTTP(res, req)

	if res.Code != http.StatusInternalServerError {
		t.Errorf("unexpected status code: got %d, want %d", res.Code, http.StatusInternalServerError)
	}

	expectedErrorMessage := "failed to decrement value for key 'name': value for key 'name' is not a valid integer: strconv.Atoi: parsing \"Seongbin\": invalid syntax\n"
	if res.Body.String() != expectedErrorMessage {
		t.Errorf("unexpected response body: got %q, want %q", res.Body.String(), expectedErrorMessage)
	}
}
