package app

import (
	"log"
	"net/http"
)

func Start() {

	// Routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	// Starts
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
