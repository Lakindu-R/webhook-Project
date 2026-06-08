package handlers

import (
	"encoding/json"
	"net/http"
)

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Access Granted (JWT + IP Verified)",
	})
}