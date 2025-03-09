package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil { // Fixed 'nill' typo to 'nil'
		fmt.Fprintf(w, "ParseForm() error: %v", err)
		return
	}
	fmt.Fprintf(w, "Post REQUEST SUCCESSFUL\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name=%s\n", name)
	fmt.Fprintf(w, "Address=%s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed) // Changed StatusNotFound to StatusMethodNotAllowed
		return
	}
	fmt.Fprintf(w, "Hello")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) // Fixed incorrect variable assignment
	http.Handle("/", fileServer)                        // Fixed incorrect function usage (Handler -> Handle)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8081\n")
	if err := http.ListenAndServe(":8081", nil); err != nil { // Fixed 'nill' typo to 'nil'
		log.Fatal(err)
	}
}
