package routes

import (
	"net/http"
	"webhook-project/handlers"
	"webhook-project/middleware"
)

func RegisterRoutes() {
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/whitelist", middleware.JWTMiddleware(http.HandlerFunc(handlers.AddIP)))
	http.HandleFunc("/whitelist/get", middleware.JWTMiddleware(http.HandlerFunc(handlers.GetIPs)))
	http.HandleFunc("/whitelist/single", middleware.JWTMiddleware(http.HandlerFunc(handlers.GetSingleIP)))
	http.HandleFunc("/whitelist/update", middleware.JWTMiddleware(http.HandlerFunc(handlers.UpdateIP)))
	// http.Handle("/whitelist/update", middleware.JWTMiddleware(http.HandlerFunc(handlers.UpdateIPStauts)))
	http.Handle("/whitelist/delete", middleware.JWTMiddleware(http.HandlerFunc(handlers.DeleteIP)))
	http.HandleFunc("/logs", middleware.JWTMiddleware(http.HandlerFunc(handlers.GetLogs)))
	http.HandleFunc("/register", middleware.JWTMiddleware(http.HandlerFunc(handlers.CreateUser)))

	http.HandleFunc("/protected",
		middleware.JWTMiddleware(
			middleware.WhitelistMiddleware(http.HandlerFunc(handlers.ProtectedHandler)),
		),
	)

	http.HandleFunc("/webhook", handlers.WebhookHandler)
	http.HandleFunc("/events", handlers.GetEventsHandler)

}
