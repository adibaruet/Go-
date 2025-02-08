package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequest() {
	http.HandleFunc("/", homePage) // Fixed function name
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequest()
}
