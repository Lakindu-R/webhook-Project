package handlers

import (
	"encoding/json"
	"net/http"

	"webhook-project/models"
)

var Events []models.Event

func WebhookHandler(w http.ResponseWriter, r *http.Request) {

	var event models.Event

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	Events = append(Events, event)

	json.NewEncoder(w).Encode(map[string]string{
		"status": "received",
	})
}

func GetEventsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(Events)
}