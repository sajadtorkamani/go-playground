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
	Quote{Quote: "We're all going to die", Author: "Sajad"},
	Quote{Quote: "Life is absurd", Author: "Sajad"},
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage!!!!!!!!")
}

func listQuotes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(quotes)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/quotes", listQuotes)
}

func main() {
	PORT := 8080
	fmt.Printf("Server listening on http://localhost:%v\n", PORT)
	handleRequests()
	http.ListenAndServe(fmt.Sprintf(":%v", PORT), nil)

}
