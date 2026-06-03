package handlers

import (
	"encoding/json"
	"net/http"

	"webhook-project/models"
	"webhook-project/storage"
)

// POST /webhook
func WebhookHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	var event models.Event

	err := json.NewDecoder(r.Body).Decode(&event)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	storage.Events = append(storage.Events, event)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Webhook received",
	})
}

// GET /events
func GetEventsHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(storage.Events)
}