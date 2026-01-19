package main

import (
	"fmt"
	"net/http"
)

// The handler function
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my Go Server! 🐹")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	// 1. Map paths to handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server starting on :8080...")
	
	// 2. Start the server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed:", err)
	}
}