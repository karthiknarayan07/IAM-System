package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		SSLMode  string
		URL      string // Added to support DATABASE_URL directly
	}
	Server struct {
		HTTPPort string
		GRPCPort string
		Debug    bool
	}
	JWT struct {
		Secret string
		TTL    int // in hours
	}
}

func LoadConfig() (*Config, error) {
	var config Config

	// Load environment variables from .env file if present
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		// Ignore error if .env file doesn't exist
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error reading .env file: %v", err)
		}
	}

	// Bind environment variables
	viper.BindEnv("DB_HOST")
	viper.BindEnv("DB_PORT")
	viper.BindEnv("DB_USER")
	viper.BindEnv("DB_PASSWORD")
	viper.BindEnv("DB_NAME")
	viper.BindEnv("DB_SSLMODE")
	viper.BindEnv("DATABASE_URL") // Support DATABASE_URL

	viper.BindEnv("HTTP_PORT")
	viper.BindEnv("GRPC_PORT")
	viper.BindEnv("DEBUG")

	viper.BindEnv("JWT_SECRET")
	viper.BindEnv("JWT_TTL")

	// Set defaults for missing config values
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", 5432)
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_NAME", "postgres")
	viper.SetDefault("DB_SSLMODE", "disable")
	viper.SetDefault("HTTP_PORT", "8080")
	viper.SetDefault("GRPC_PORT", "50051")
	viper.SetDefault("JWT_TTL", 24)

	// Map environment variables to struct fields
	config.Database.Host = viper.GetString("DB_HOST")
	config.Database.Port = viper.GetInt("DB_PORT")
	config.Database.User = viper.GetString("DB_USER")
	config.Database.Password = viper.GetString("DB_PASSWORD")
	config.Database.DBName = viper.GetString("DB_NAME")
	config.Database.SSLMode = viper.GetString("DB_SSLMODE")
	config.Database.URL = viper.GetString("DATABASE_URL") // Use DATABASE_URL if available

	config.Server.HTTPPort = viper.GetString("HTTP_PORT")
	config.Server.GRPCPort = viper.GetString("GRPC_PORT")
	config.Server.Debug = viper.GetBool("DEBUG")

	config.JWT.Secret = viper.GetString("JWT_SECRET")
	config.JWT.TTL = viper.GetInt("JWT_TTL")

	// Validate required configurations
	if err := validateConfig(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func validateConfig(config *Config) error {
	if config.Database.Password == "" {
		return fmt.Errorf("database password is required")
	}
	if config.JWT.Secret == "" {
		return fmt.Errorf("JWT secret is required")
	}
	return nil
}
