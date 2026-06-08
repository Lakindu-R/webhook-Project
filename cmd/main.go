package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"

	"webhook-project/routes"
	"webhook-project/config"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file",err)
	}

	config.ConnectDB()
	routes.RegisterRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}