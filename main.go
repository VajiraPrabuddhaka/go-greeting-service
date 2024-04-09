package main

import (
	"fmt"
	"log"
	"net/http"
)

// GreetingHandler handles incoming HTTP requests for greeting.
func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	// Set the response header
	w.Header().Set("Content-Type", "text/plain")

	// Check the request method
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse query parameter "name"
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	// Write the response
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	// Register the handler function with the default serve mux ("/greet" route)
	http.HandleFunc("/greet", GreetingHandler)
	// Start the HTTP server on port 8080
	fmt.Println("Server is listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
