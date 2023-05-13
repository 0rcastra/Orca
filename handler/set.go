package handler

import (
	"encoding/json"
	"net/http"

	"github.com/0rcastra/Orca/internal/data"
	"github.com/gorilla/mux"
)

type SetResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func SetHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the key and value from the request URL path parameters
	vars := mux.Vars(r)
	key := vars["key"]
	value := vars["value"]

	// Create a new instance of the Database
	db := data.NewDatabase()

	// Set the key-value pair in the data store
	db.Set(key, value)

	// Send the response back to the client
	response := SetResponse{
		Key:   key,
		Value: value,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
