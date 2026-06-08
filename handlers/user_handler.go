package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"webhook-project/config"
	"webhook-project/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	_, err := config.DB.Collection("users").
		InsertOne(context.Background(), user)
	
		if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User created successfully",
	})
}