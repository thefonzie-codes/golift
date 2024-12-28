package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/thefonzie-codes/goLift/internal/config"
	"github.com/thefonzie-codes/goLift/internal/database"
	"github.com/thefonzie-codes/goLift/internal/routes"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found: %v", err)
	}

	// Initialize config
	cfg := config.New()

	// Initialize database
	db, err := database.Initialize(cfg)
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}
	defer db.Close()

	// Initialize router with config
	router := routes.SetupRoutes(db, cfg)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
