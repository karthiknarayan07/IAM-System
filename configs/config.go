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
	}
	Server struct {
		HTTPPort string
		GRPCPort string
	}
	JWT struct {
		Secret string
		TTL    int // in hours
	}
}

func LoadConfig() (*Config, error) {
	var config Config

	// Database configs
	viper.BindEnv("DB_HOST")
	viper.BindEnv("DB_PORT")
	viper.BindEnv("DB_USER")
	viper.BindEnv("DB_PASSWORD")
	viper.BindEnv("DB_NAME")
	viper.BindEnv("DB_SSLMODE")

	// Server configs
	viper.BindEnv("HTTP_PORT")
	viper.BindEnv("GRPC_PORT")

	// JWT configs
	viper.BindEnv("JWT_SECRET")
	viper.BindEnv("JWT_TTL")

	// Set defaults
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", 5432)
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_NAME", "postgres")
	viper.SetDefault("DB_SSLMODE", "disable")
	viper.SetDefault("HTTP_PORT", "8080")
	viper.SetDefault("GRPC_PORT", "50051")
	viper.SetDefault("JWT_TTL", 24)

	// Map environment variables to struct
	config.Database.Host = viper.GetString("DB_HOST")
	config.Database.Port = viper.GetInt("DB_PORT")
	config.Database.User = viper.GetString("DB_USER")
	config.Database.Password = viper.GetString("DB_PASSWORD")
	config.Database.DBName = viper.GetString("DB_NAME")
	config.Database.SSLMode = viper.GetString("DB_SSLMODE")

	config.Server.HTTPPort = viper.GetString("HTTP_PORT")
	config.Server.GRPCPort = viper.GetString("GRPC_PORT")

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
