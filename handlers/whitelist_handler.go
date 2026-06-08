package handlers

import (
	"encoding/json"
	"net/http"
	"webhook-project/config"
	"webhook-project/models"
	"webhook-project/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddIP(w http.ResponseWriter, r *http.Request) {

	var ip models.WhitelistIP
	json.NewDecoder(r.Body).Decode(&ip)

	// VALIDATION
	if !utils.IsValidIP(ip.IPAddress) {
		http.Error(w, "Invalid IP format", http.StatusBadRequest)
		return
	}

	// DUPLICATE CHECK
	var existing models.WhitelistIP
	err := config.DB.Collection("whitelisted_ips").
		FindOne(r.Context(), bson.M{"ip_address": ip.IPAddress}).Decode(&existing)

	if err == nil {
		http.Error(w, "Duplicate IP not allowed", http.StatusBadRequest)
		return
	}

	ip.Enabled = true

	config.DB.Collection("whitelisted_ips").InsertOne(r.Context(), ip)

	json.NewEncoder(w).Encode(ip)
}

func GetSingleIP(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	objID, _ := primitive.ObjectIDFromHex(id)

	var ip models.WhitelistIP

	err := config.DB.Collection("whitelisted_ips").
		FindOne(r.Context(), bson.M{"_id": objID}).Decode(&ip)

	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(ip)
}

func GetIPs(w http.ResponseWriter, r *http.Request) {
	cur, _ := config.DB.Collection("whitelisted_ips").Find(r.Context(), bson.M{})
	var ips []models.WhitelistIP
	cur.All(r.Context(), &ips)
	json.NewEncoder(w).Encode(ips)
}

func UpdateIPStauts(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	objID, _ := primitive.ObjectIDFromHex(id)
	var body struct {
		Enabled bool `json:"enabled"`
	}
	json.NewDecoder(r.Body).Decode(&body)
	config.DB.Collection("whitelisted_ips").UpdateOne(
		r.Context(),
		bson.M{"_id": objID},
		bson.M{"$set": bson.M{"enabled": body.Enabled}},
	)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "IP status updated successfully",
	})
}

func UpdateIP(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	objID, _ := primitive.ObjectIDFromHex(id)

	var body models.WhitelistIP
	json.NewDecoder(r.Body).Decode(&body)

	if !utils.IsValidIP(body.IPAddress) {
		http.Error(w, "Invalid IP", http.StatusBadRequest)
		return
	}

	config.DB.Collection("whitelisted_ips").UpdateOne(
		r.Context(),
		bson.M{"_id": objID},
		bson.M{"$set": bson.M{
			"ip_address": body.IPAddress,
			"enabled":    body.Enabled,
		}},
	)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "updated",
	})
}

func DeleteIP(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	objID, _ := primitive.ObjectIDFromHex(id)
	config.DB.Collection("whitelisted_ips").DeleteOne(
		r.Context(),
		bson.M{"_id": objID},
	)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "IP deleted successfully",
	})
}
