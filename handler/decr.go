package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type DecrResponse struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

func (h *Handler) DecrHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	newValue, err := h.db.Decr(key)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to decrement value for key '%s': %v", key, err), http.StatusInternalServerError)
		return
	}

	response := DecrResponse{
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
