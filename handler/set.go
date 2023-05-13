package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) SetHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the key and value from the request URL path parameters
	vars := mux.Vars(r)
	key := vars["key"]
	value := vars["value"]

	// Set the key-value pair in the data store
	h.db.Set(key, value)

	// Send the response back to the client
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))

}
