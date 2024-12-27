package testutils

import (
	"database/sql"
	"os"
	"testing"

	"github.com/thefonzie-codes/goLift/internal/config"
)

// TestDB returns a test database connection
func TestDB(t *testing.T) *sql.DB {
	cfg := &config.Config{
		DBHost:     os.Getenv("TEST_DB_HOST"),
		DBPort:     os.Getenv("TEST_DB_PORT"),
		DBUser:     os.Getenv("TEST_DB_USER"),
		DBPassword: os.Getenv("TEST_DB_PASSWORD"),
		DBName:     os.Getenv("TEST_DB_NAME"),
	}

	db, err := sql.Open("postgres", cfg.DBHost)
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	return db
}

// CleanupDB cleans up the test database after tests
func CleanupDB(t *testing.T, db *sql.DB) {
	if err := db.Close(); err != nil {
		t.Errorf("Error closing test database: %v", err)
	}
}
