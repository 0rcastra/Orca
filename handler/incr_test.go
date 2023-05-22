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

func TestIncrHandler(t *testing.T) {
	db := data.NewDatabase()
	db.Set("count", "5")

	req := httptest.NewRequest(http.MethodPost, "/incr/count", nil)
	res := httptest.NewRecorder()

	h := handler.NewHandler(db)

	router := mux.NewRouter()
	router.HandleFunc("/incr/{key}", h.IncrHandler).Methods(http.MethodPost)

	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %d, want %d", res.Code, http.StatusOK)
	}

	var response handler.IncrResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		t.Errorf("failed to decode response body: %s", err.Error())
	}

	expectedKey := "count"
	expectedValue := 6
	if response.Key != expectedKey {
		t.Errorf("unexpected response key: got %s, want %s", response.Key, expectedKey)
	}
	if response.Value != expectedValue {
		t.Errorf("unexpected response value: got %d, want %d", response.Value, expectedValue)
	}
}

func TestIncrHandler_Error(t *testing.T) {
	db := data.NewDatabase()
	db.Set("count", "invalid")

	req := httptest.NewRequest(http.MethodPost, "/incr/count", nil)
	res := httptest.NewRecorder()

	h := handler.NewHandler(db)

	router := mux.NewRouter()
	router.HandleFunc("/incr/{key}", h.IncrHandler).Methods(http.MethodPost)

	router.ServeHTTP(res, req)

	if res.Code != http.StatusInternalServerError {
		t.Errorf("unexpected status code: got %d, want %d", res.Code, http.StatusInternalServerError)
	}

	expectedErrorMessage := "failed to increment value for key 'count': value for key 'count' is not a valid integer: strconv.Atoi: parsing \"invalid\": invalid syntax\n"
	if res.Body.String() != expectedErrorMessage {
		t.Errorf("unexpected response body: got %q, want %q", res.Body.String(), expectedErrorMessage)
	}
}
