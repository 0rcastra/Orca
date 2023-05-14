package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type DelResponse struct {
	Key     string `json:"key"`
	Deleted bool   `json:"deleted"`
}

func (h *Handler) DelHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	deleted := h.db.Del(key)

	response := DelResponse{
		Key:     key,
		Deleted: deleted,
	}

	w.Header().Set("Content-Type", "application/json")
	if deleted {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
	json.NewEncoder(w).Encode(response)
}
