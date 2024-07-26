package main

import (
	"fmt"
	"log"
	"net/http"
)

func notifyHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Handle the webhook payload
	// Here you can process the payload received from the webhook
	// For example, you might parse JSON or form data from r.Body

	// Example: Read and log the request body
	// Replace this with your actual processing logic
	log.Println("Webhook received:")
	// Example reading the body
	buf := make([]byte, 1024)
	n, err := r.Body.Read(buf)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	log.Println(string(buf[:n]))

	// Respond to the webhook
	fmt.Fprintf(w, "Webhook received successfully\n")
}

func main() {
	// Register handler for /notify endpoint
	http.HandleFunc("/notify", notifyHandler)

	// Start server on port 8082
	log.Println("Starting server on port 8082...")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
