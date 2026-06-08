package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"

	"webhook-project/config"
	"webhook-project/models"
	"webhook-project/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	var user models.User

	err := config.DB.Collection("users").
		FindOne(context.TODO(), bson.M{
			"username": req.Username,
			"password": req.Password,
		}).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, _ := utils.GenerateToken(user.Username)

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
