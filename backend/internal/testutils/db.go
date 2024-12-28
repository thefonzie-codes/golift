package testutils

import (
	"database/sql"
	"log"
	"testing"

	"github.com/thefonzie-codes/goLift/internal/config"
	"github.com/thefonzie-codes/goLift/internal/database"
)

func SetupTestDB(t *testing.T) *sql.DB {
	cfg := &config.Config{
		DBHost:     "localhost",
		DBPort:     "5434",
		DBName:     "golift_test",
		DBUser:     "thefonziecodes",
		DBPassword: "Alfie@3046",
	}

	db, err := database.Initialize(cfg)
	if err != nil {
		t.Fatalf("Failed to initialize test database: %v", err)
	}

	// Clear all tables before each test
	tables := []string{"progress", "workout_exercises", "athlete_maxes", "exercises", "workouts", "programs", "users"}
	for _, table := range tables {
		_, err := db.Exec("DELETE FROM " + table)
		if err != nil {
			log.Fatalf("Failed to clear table %s: %v", table, err)
		}
	}

	return db
}

func TeardownTestDB(t *testing.T, db *sql.DB) {
	if err := db.Close(); err != nil {
		t.Errorf("Failed to close test database: %v", err)
	}
} 