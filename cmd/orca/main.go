package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/0rcastra/Orca/handler"
	"github.com/gorilla/mux"
)

func main() {
	// Parse command-line flags
	port := flag.Int("port", 8080, "the port on which to listen for incoming connections")
	flag.Parse()

	// Create the router
	r := mux.NewRouter()

	// Defining the commands
	r.HandleFunc("/set/{key}/{value}", handler.SetHandler).Methods("POST")

	// Create the HTTP server
	server := &http.Server{
		Addr:           ":" + strconv.Itoa(*port),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Start the server
	log.Printf("Starting Orca server on port %d...", *port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
