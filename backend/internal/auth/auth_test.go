package auth

import (
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func TestGenerateAndParseToken(t *testing.T) {
	userID := uuid.New()
	role := "athlete"
	secret := "test-secret"

	// Generate token
	token, err := GenerateToken(userID, role, secret)
	if err != nil {
		t.Fatalf("GenerateToken() error = %v", err)
	}

	// Parse and validate token
	parsedToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		t.Fatalf("Failed to parse token: %v", err)
	}

	claims, ok := parsedToken.Claims.(*Claims)
	if !ok || !parsedToken.Valid {
		t.Fatal("Invalid token claims")
	}

	if claims.UserID != userID {
		t.Errorf("UserID = %v, want %v", claims.UserID, userID)
	}
	if claims.Role != role {
		t.Errorf("Role = %v, want %v", claims.Role, role)
	}
}

func TestPasswordHashing(t *testing.T) {
	password := "testpassword123"

	// Test hashing
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword() error = %v", err)
	}

	// Test valid password
	if !CheckPassword(password, hash) {
		t.Error("CheckPassword() failed for valid password")
	}

	// Test invalid password
	if CheckPassword("wrongpassword", hash) {
		t.Error("CheckPassword() succeeded for invalid password")
	}
} 