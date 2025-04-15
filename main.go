package main

import (
	"library/config"
	"library/src/routes"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database connection
	if err := config.ConnectDB(); err != nil {
		log.Fatal("Failed to connect to database")
	}

	// Initialize server
	router := routes.SetupRoutes()
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
    }
