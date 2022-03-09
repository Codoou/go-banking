package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	//mux := http.NewServeMux()
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/greet", greet)
	router.HandleFunc("/customers", getAllCustomers)

	// Starts
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
