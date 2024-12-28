package config

import "os"

type Config struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	JWTSecret  string
}

func New() *Config {
	return &Config{
		DBHost:     getEnvOrDefault("DB_HOST", "localhost"),
		DBPort:     getEnvOrDefault("DB_PORT", "5434"),
		DBName:     getEnvOrDefault("DB_NAME", "golift_dev"),
		DBUser:     getEnvOrDefault("DB_USER", "thefonziecodes"),
		DBPassword: getEnvOrDefault("DB_PASSWORD", "Alfie@3046"),
		JWTSecret:  getEnvOrDefault("JWT_SECRET", "your_jwt_secret_key"),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
