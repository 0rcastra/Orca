package handler

import (
	"net/http"

	"github.com/0rcastra/Orca/internal/data"
	"github.com/gorilla/mux"
)

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
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
