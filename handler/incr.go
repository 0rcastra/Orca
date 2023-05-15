package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type IncrResponse struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

func (h *Handler) IncrHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

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
