package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Define Article struct
type Article struct {
	Title string `json:"Title"`
	Desc  string `json:"content"`
}

// Define Articles as a slice of Article
type Articles []Article

// Handler function for "/articles" endpoint
func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Description"},
	}

	// Convert struct to JSON and send as a response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

// Handler function for "/"
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

// Function to handle requests and route them
func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequest()
}
