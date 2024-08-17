package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Quote struct {
	Quote  string
	Author string
}

var quotes = []Quote{
	Quote{Quote: "Quote 1", Author: "Sajad"},
	Quote{Quote: "Quote 2", Author: "Sajad"},
	Quote{Quote: "Quote 3", Author: "Sajad"},
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage!.")
}

func listQuotes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(quotes)
}

func createQuote(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Message string
	}{
		Message: "Ok!",
	}

	json.NewEncoder(w).Encode(response)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("GET /quotes", listQuotes)
	http.HandleFunc("POST /quotes", createQuote)
}

func main() {
	PORT := 8080
	fmt.Printf("Server listening on http://localhost:%v\n", PORT)
	handleRequests()

	err := http.ListenAndServe(fmt.Sprintf(":%v", PORT), nil)
	if err != nil {
		panic(err)
	}
}
