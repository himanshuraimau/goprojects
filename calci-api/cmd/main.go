package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/himanshuraimau/goprojects/calci-api/api/handlers"
)

func main() {
	r := mux.NewRouter()

	// Define the API routes
	r.HandleFunc("/add", handlers.Add).Methods("POST")
	r.HandleFunc("/subtract", handlers.Subtract).Methods("POST")
	r.HandleFunc("/multiply", handlers.Multiply).Methods("POST")
	r.HandleFunc("/divide", handlers.Divide).Methods("POST")
	r.HandleFunc("/sum", handlers.Sum).Methods("POST")

	// Start the server
	log.Println("Starting server on :3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
