package main

import (
	"github.com/karthiknarayan07/IAM-System/http"
	"github.com/karthiknarayan07/IAM-System/http/handlers"
	"github.com/karthiknarayan07/IAM-System/repository"
	"github.com/karthiknarayan07/IAM-System/service"

	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Fetch database connection string from environment variables
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	// Initialize the database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize user-related layers
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Set up the router
	router := http.NewRouter()
	router.RegisterHandlers(userHandler)

	// Start the HTTP server
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080"
	}
	http.StartServer(router, port)
}
