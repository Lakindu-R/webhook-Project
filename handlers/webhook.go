package handlers

import (
	"encoding/json"
	"net/http"
	"webhook-project/models"
)

var Events []models.Event

func WebhookHandler(w http.ResponseWriter, r *http.Request) {

	var e models.Event

	json.NewDecoder(r.Body).Decode(&e)
	Events = append(Events, e)

	json.NewEncoder(w).Encode(map[string]string{
		"status": "received",
	})
}

func GetEventsHandler(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(Events)
}
