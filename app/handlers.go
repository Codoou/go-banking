package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Cody", City: "Wausau", Zipcode: "54455"},
		{Name: "Frank", City: "Plover", Zipcode: "54401"},
		{Name: "Bob", City: "Stevens Point", Zipcode: "90210"},
	}

	// How to get header keys
	// r.Header.Get("Content-Type")

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
