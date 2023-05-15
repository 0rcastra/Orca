package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/0rcastra/Orca/internal/data"
	"github.com/gorilla/mux"
)

// GetHandler handles the GET requests for retrieving a value by key.
func (h *Handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the key from the request path
	vars := mux.Vars(r)
	key := vars["key"]

	// Retrieve the value from the data store
	value, exists := data.Get(h.db, key)
	if !exists {
		// Key not found
		w.WriteHeader(http.StatusNotFound)
		response := ErrorResponse{Message: fmt.Sprintf("Key '%s' not found", key)}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Key found, return the value
	response := GetResponse{Key: key, Value: value}
	json.NewEncoder(w).Encode(response)
}

// GetResponse represents the response for the GET request.
type GetResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// ErrorResponse represents the error response.
type ErrorResponse struct {
	Message string `json:"message"`
}
