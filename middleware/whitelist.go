package middleware

import (
	"fmt"
	"net/http"

	"webhook-project/config"
	"webhook-project/models"
	"webhook-project/services"

	"go.mongodb.org/mongo-driver/bson"
)

func WhitelistMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		ip := GetClientIP(r)
		fmt.Println("Extract IP",ip)

		var result models.WhitelistIP
		fmt.Println("CLIENT IP =", ip)

		err := config.DB.Collection("whitelisted_ips").
			FindOne(r.Context(), bson.M{
				"ip_address": ip,
				"enabled":    true,
			}).Decode(&result)

		if err != nil {

			services.SaveLog(ip, r.URL.Path, "DENIED", "IP not whitelisted")

			http.Error(w, "IP Not Allowed", http.StatusForbidden)
			return
		}

		services.SaveLog(ip, r.URL.Path, "ALLOWED", "")

		next(w, r)
	}
}
