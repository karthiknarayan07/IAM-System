package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/karthiknarayan07/IAM-System/config"
	"github.com/karthiknarayan07/IAM-System/http/handlers"
	"github.com/karthiknarayan07/IAM-System/repository"
	"github.com/karthiknarayan07/IAM-System/service"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load the application configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Set Gin mode (debug, release, test)
	if cfg.Server.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Use the DATABASE_URL if it's available, otherwise fall back to individual DB configurations
	var dsn string
	if cfg.Database.URL != "" {
		dsn = cfg.Database.URL // DATABASE_URL is prioritized if available
	} else {
		dsn = "host=" + cfg.Database.Host + " port=" + strconv.Itoa(cfg.Database.Port) +
			" user=" + cfg.Database.User + " password=" + cfg.Database.Password +
			" dbname=" + cfg.Database.DBName + " sslmode=" + cfg.Database.SSLMode
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

	// Set up Gin router
	router := gin.Default() // Includes logger and recovery middleware

	// Register user handlers
	router.POST("/users", userHandler.RegisterUser)   // Maps to RegisterUser
	router.GET("/users/:id", userHandler.GetUserByID) // Maps to GetUserByID

	// Start the HTTP server
	port := cfg.Server.HTTPPort
	log.Printf("Starting server on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
