package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/0rcastra/Orca/internal/data"
	"github.com/gorilla/mux"
)

type IncrResponse struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

func (h *Handler) IncrHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, exists := data.Get(h.db, key)
	if !exists {
		// Key not found
		w.WriteHeader(http.StatusNotFound)
		response := ErrorResponse{Message: fmt.Sprintf("Key '%s' not found", key)}
		json.NewEncoder(w).Encode(response)
		return
	}

	_, ok := value.(int)
	if !ok {
		http.Error(w, fmt.Sprintf("value for key '%s' is not an integer", key), http.StatusInternalServerError)
		return
	}

	newValue, err := h.db.Incr(key)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to increment value for key '%s': %v", key, err), http.StatusInternalServerError)
		return
	}

	response := IncrResponse{
		Key:   key,
		Value: newValue,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}
}
