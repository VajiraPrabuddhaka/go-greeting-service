package main

import (
	"encoding/json"
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

	fmt.Printf("Making a http request to private service name:%s", name)

	requestURL := fmt.Sprintf("http://tsproxy-service:%d?name=%s", 8080, name)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
	}

	bookStatus := new(LibraryServiceResponse)

	// Read the response body
	if err := json.NewDecoder(res.Body).Decode(bookStatus); err != nil {
		fmt.Printf("error parsing response: %s\n", err)
	}

	// Write the response
	fmt.Fprintf(w, "Hello, %s..!! %s", bookStatus.Name, bookStatus.Message)
}

type LibraryServiceResponse struct {
	Name    string `json:"name"`
	Message string `json:"message"`
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
