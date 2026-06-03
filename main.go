package main

import (
	"fmt"
	"net/http"
	"webhook-project/handlers"
)

func main() {
	http.HandleFunc("/webhook", handlers.WebhookHandler)
	http.HandleFunc("/events", handlers.GetEventsHandler)

	fmt.Println("Server running on :8080")

	http.ListenAndServe(":8080", nil)
}
