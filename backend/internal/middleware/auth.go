package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/thefonzie-codes/goLift/internal/auth"
	"github.com/thefonzie-codes/goLift/internal/config"
)

func AuthMiddleware(cfg *config.Config) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			log.Printf("Auth header: %s", authHeader)

			if authHeader == "" {
				http.Error(w, "Authorization header required", http.StatusUnauthorized)
				return
			}

			// Extract the token from the Authorization header
			// Format: "Bearer <token>"
			bearerToken := strings.Split(authHeader, " ")
			log.Printf("Split token: %v", bearerToken)

			if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
				http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
				return
			}

			// Parse and validate the token
			token, err := jwt.ParseWithClaims(bearerToken[1], &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(cfg.JWTSecret), nil
			})

			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			claims, ok := token.Claims.(*auth.Claims)
			if !ok || !token.Valid {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}

			// Add claims to request context
			ctx := context.WithValue(r.Context(), "user", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// Optional: Role-based middleware
func RequireRole(roles ...string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, ok := r.Context().Value("user").(*auth.Claims)
			if !ok {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			for _, role := range roles {
				if claims.Role == role {
					next.ServeHTTP(w, r)
					return
				}
			}

			http.Error(w, "Forbidden", http.StatusForbidden)
		})
	}
} 