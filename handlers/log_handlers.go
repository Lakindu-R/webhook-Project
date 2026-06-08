package handlers

import (
	"encoding/json"
	"net/http"

	"webhook-project/config"
	"webhook-project/models"
)

func GetLogs(w http.ResponseWriter, r *http.Request) {

	cur, _ := config.DB.Collection("access_logs").Find(r.Context(), map[string]interface{}{})

	var logs []models.AccessLog
	cur.All(r.Context(), &logs)

	json.NewEncoder(w).Encode(logs)
}