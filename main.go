package main

import (
	"encoding/json"
	"net/http"
)

// Define a simple structure for the response
type Response struct {
	Message string `json:"message"`
}

// Handler function for the root endpoint
func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello, World!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", helloHandler) // Set up the route
	http.ListenAndServe(":3000", nil)  // Start the server on port 3000
}
