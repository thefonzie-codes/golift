package config

import (
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	// Test with environment variables
	os.Setenv("DB_HOST", "testhost")
	os.Setenv("DB_PORT", "5555")
	os.Setenv("DB_NAME", "testdb")
	os.Setenv("DB_USER", "testuser")
	os.Setenv("DB_PASSWORD", "testpass")
	os.Setenv("JWT_SECRET", "testsecret")

	cfg := New()

	tests := []struct {
		name     string
		got      string
		want     string
		envKey   string
		defaultVal string
	}{
		{"DBHost", cfg.DBHost, "testhost", "DB_HOST", "localhost"},
		{"DBPort", cfg.DBPort, "5555", "DB_PORT", "5434"},
		{"DBName", cfg.DBName, "testdb", "DB_NAME", "golift_dev"},
		{"DBUser", cfg.DBUser, "testuser", "DB_USER", "thefonziecodes"},
		{"DBPassword", cfg.DBPassword, "testpass", "DB_PASSWORD", "Alfie@3046"},
		{"JWTSecret", cfg.JWTSecret, "testsecret", "JWT_SECRET", "your_jwt_secret_key"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("New() %s = %v, want %v", tt.name, tt.got, tt.want)
			}

			// Test default values
			os.Unsetenv(tt.envKey)
			cfg = New()
			if got := getEnvOrDefault(tt.envKey, tt.defaultVal); got != tt.defaultVal {
				t.Errorf("getEnvOrDefault() = %v, want %v", got, tt.defaultVal)
			}
		})
	}
}
