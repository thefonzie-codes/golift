package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/thefonzie-codes/goLift/pkg/testutils"
)

func TestNew(t *testing.T) {
	stats := testutils.NewTestStats("config")
	stats.LogInfo(t, "Starting config tests...")

	// Test with environment variables
	stats.LogInfo(t, "Testing with custom environment variables")
	os.Setenv("DB_HOST", "testhost")
	os.Setenv("DB_PORT", "5433")
	os.Setenv("DB_USER", "testuser")
	os.Setenv("DB_PASSWORD", "testpass")
	os.Setenv("DB_NAME", "testdb")
	os.Setenv("JWT_SECRET", "testsecret")

	cfg := New()

	tests := []struct {
		name     string
		got      string
		expected string
	}{
		{"DBHost", cfg.DBHost, "testhost"},
		{"DBPort", cfg.DBPort, "5433"},
		{"DBUser", cfg.DBUser, "testuser"},
		{"DBPassword", cfg.DBPassword, "testpass"},
		{"DBName", cfg.DBName, "testdb"},
		{"JWTSecret", cfg.JWTSecret, "testsecret"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got == tt.expected {
				stats.LogSuccess(t, fmt.Sprintf("%s: got %s", tt.name, tt.got))
			} else {
				stats.LogError(t, fmt.Sprintf("%s: got %s, want %s", tt.name, tt.got, tt.expected))
			}
		})
	}

	// Test default values
	stats.LogInfo(t, "Testing default values...")
	os.Clearenv()
	cfg = New()

	defaultTests := []struct {
		name     string
		got      string
		expected string
	}{
		{"Default DBHost", cfg.DBHost, "localhost"},
		{"Default DBPort", cfg.DBPort, "5432"},
		{"Default DBUser", cfg.DBUser, "postgres"},
		{"Default DBName", cfg.DBName, "golift"},
	}

	for _, tt := range defaultTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got == tt.expected {
				stats.LogSuccess(t, fmt.Sprintf("%s: got %s", tt.name, tt.got))
			} else {
				stats.LogError(t, fmt.Sprintf("%s: got %s, want %s", tt.name, tt.got, tt.expected))
			}
		})
	}

	stats.PrintSummary(t)
}
