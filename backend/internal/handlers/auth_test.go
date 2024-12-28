package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/thefonzie-codes/goLift/internal/config"
	"github.com/thefonzie-codes/goLift/internal/testutils"
	"github.com/thefonzie-codes/goLift/internal/auth"
)

func TestRegister(t *testing.T) {
	db := testutils.SetupTestDB(t)
	defer testutils.TeardownTestDB(t, db)

	cfg := &config.Config{JWTSecret: "test-secret"}
	h := NewHandler(db, cfg)

	tests := []struct {
		name       string
		payload    RegisterRequest
		wantStatus int
	}{
		{
			name: "Valid Registration",
			payload: RegisterRequest{
				Name:     "Test User",
				Email:    "test@example.com",
				Password: "password123",
				Role:     "athlete",
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "Invalid Role",
			payload: RegisterRequest{
				Name:     "Test User",
				Email:    "test@example.com",
				Password: "password123",
				Role:     "invalid",
			},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			h.Register(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Register() status = %v, want %v", w.Code, tt.wantStatus)
			}

			if w.Code == http.StatusOK {
				var response map[string]interface{}
				json.NewDecoder(w.Body).Decode(&response)

				if response["token"] == nil {
					t.Error("Register() response missing token")
				}
				if response["user"] == nil {
					t.Error("Register() response missing user")
				}
			}
		})
	}
}

func TestLogin(t *testing.T) {
	db := testutils.SetupTestDB(t)
	defer testutils.TeardownTestDB(t, db)

	cfg := &config.Config{JWTSecret: "test-secret"}
	h := NewHandler(db, cfg)

	// First register a test user
	hashedPassword, _ := auth.HashPassword("password123")
	_, err := db.Exec(`
		INSERT INTO users (name, email, password_hash, role)
		VALUES ($1, $2, $3, $4)`,
		"Test User", "test@example.com", hashedPassword, "athlete",
	)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	tests := []struct {
		name       string
		payload    LoginRequest
		wantStatus int
	}{
		{
			name: "Valid Login",
			payload: LoginRequest{
				Email:    "test@example.com",
				Password: "password123",
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "Invalid Password",
			payload: LoginRequest{
				Email:    "test@example.com",
				Password: "wrongpassword",
			},
			wantStatus: http.StatusUnauthorized,
		},
		{
			name: "User Not Found",
			payload: LoginRequest{
				Email:    "nonexistent@example.com",
				Password: "password123",
			},
			wantStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			h.Login(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Login() status = %v, want %v", w.Code, tt.wantStatus)
			}

			if w.Code == http.StatusOK {
				var response map[string]string
				json.NewDecoder(w.Body).Decode(&response)

				if response["token"] == "" {
					t.Error("Login() response missing token")
				}
			}
		})
	}
} 