package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type URLShortener struct {
	mu   sync.RWMutex
	urls map[string]string
}

// Shorten handles the POST request to create a short URL
func (s *URLShortener) Shorten(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests for shortening
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		LongURL string `json:"long_url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Simple key generation: hex of the string length
	// In a real app, you'd use a hash like MD5 or a random string
	key := fmt.Sprintf("%x", len(input.LongURL))

	s.mu.Lock()
	s.urls[key] = input.LongURL
	s.mu.Unlock()

	response := map[string]string{
		"short_url": "http://localhost:8080/" + key,
		"key":       key,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Redirect handles the GET request for the root and the short keys
func (s *URLShortener) Redirect(w http.ResponseWriter, r *http.Request) {
	// 1. Handle the actual root path (Home Page)
	if r.URL.Path == "/" {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, `
			<h1>Go URL Shortener 🐹</h1>
			<p>To shorten a URL, send a POST request to <code>/shorten</code></p>
			<p>Example: <code>{"long_url": "https://google.com"}</code></p>
		`)
		return
	}

	// 2. Handle the keys (anything after the /)
	key := r.URL.Path[1:] // Trims the leading slash

	s.mu.RLock()
	longURL, exists := s.urls[key]
	s.mu.RUnlock()

	if !exists {
		http.NotFound(w, r)
		return
	}

	// 3. Perform the redirect
	http.Redirect(w, r, longURL, http.StatusFound)
}

func main() {
	// Initialize the map
	shortener := &URLShortener{
		urls: make(map[string]string),
	}

	// Routes
	http.HandleFunc("/shorten", shortener.Shorten)
	http.HandleFunc("/", shortener.Redirect)

	fmt.Println("🚀 URL Shortener running on :8080")
	fmt.Println("   - POST to /shorten to create links")
	fmt.Println("   - GET / to see this message")

	// Start the server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server: %s\n", err)
	}
}