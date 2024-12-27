package database

import (
	"testing"

	"github.com/thefonzie-codes/goLift/internal/config"
)

func TestInitialize(t *testing.T) {
	t.Log("Starting database initialization tests...")

	tests := []struct {
		name    string
		cfg     *config.Config
		wantErr bool
	}{
		{
			name: "Invalid credentials",
			cfg: &config.Config{
				DBHost:     "localhost",
				DBPort:     "5434",
				DBUser:     "invaliduser",
				DBPassword: "invalidpass",
				DBName:     "invaliddb",
			},
			wantErr: true,
		},
		{
			name: "Valid credentials",
			cfg: &config.Config{
				DBHost:     "localhost",
				DBPort:     "5434",
				DBUser:     "thefonziecodes",
				DBPassword: "Alfie@3046",
				DBName:     "golift_test",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Testing %s configuration", tt.name)
			t.Logf("Attempting to connect to: %s:%s as %s", tt.cfg.DBHost, tt.cfg.DBPort, tt.cfg.DBUser)

			db, err := Initialize(tt.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Initialize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && db == nil {
				t.Error("Initialize() returned nil db without error")
			}
			if db != nil {
				t.Log("Successfully connected to database")
				db.Close()
			}
		})
	}
}
