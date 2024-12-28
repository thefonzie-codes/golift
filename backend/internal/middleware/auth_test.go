package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/thefonzie-codes/goLift/internal/auth"
	"github.com/thefonzie-codes/goLift/internal/config"
)

func TestAuthMiddleware(t *testing.T) {
	cfg := &config.Config{JWTSecret: "test-secret"}
	userID := uuid.New()

	// Create a test token
	token, err := auth.GenerateToken(userID, "athlete", cfg.JWTSecret)
	if err != nil {
		t.Fatalf("Failed to generate test token: %v", err)
	}

	tests := []struct {
		name       string
		authHeader string
		wantStatus int
	}{
		{
			name:       "Valid Token",
			authHeader: "Bearer " + token,
			wantStatus: http.StatusOK,
		},
		{
			name:       "No Token",
			authHeader: "",
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:       "Invalid Token Format",
			authHeader: "Bearer invalid-token",
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:       "Missing Bearer",
			authHeader: token,
			wantStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test handler that always succeeds
			nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			// Create the middleware handler
			middleware := AuthMiddleware(cfg)(nextHandler)

			// Create test request
			req := httptest.NewRequest("GET", "/", nil)
			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader)
			}

			// Create response recorder
			w := httptest.NewRecorder()

			// Serve the request
			middleware.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("AuthMiddleware() status = %v, want %v", w.Code, tt.wantStatus)
			}
		})
	}
}

func TestRequireRole(t *testing.T) {
	tests := []struct {
		name       string
		roles      []string
		userRole   string
		wantStatus int
	}{
		{
			name:       "Matching Role",
			roles:      []string{"coach"},
			userRole:   "coach",
			wantStatus: http.StatusOK,
		},
		{
			name:       "Non-matching Role",
			roles:      []string{"coach"},
			userRole:   "athlete",
			wantStatus: http.StatusForbidden,
		},
		{
			name:       "Multiple Allowed Roles",
			roles:      []string{"coach", "athlete"},
			userRole:   "athlete",
			wantStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test handler that always succeeds
			nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			// Create the middleware handler
			middleware := RequireRole(tt.roles...)(nextHandler)

			// Create test request with user claims in context
			req := httptest.NewRequest("GET", "/", nil)
			claims := &auth.Claims{
				UserID: uuid.New(),
				Role:   tt.userRole,
			}
			ctx := context.WithValue(req.Context(), "user", claims)
			req = req.WithContext(ctx)

			// Create response recorder
			w := httptest.NewRecorder()

			// Serve the request
			middleware.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("RequireRole() status = %v, want %v", w.Code, tt.wantStatus)
			}
		})
	}
} 