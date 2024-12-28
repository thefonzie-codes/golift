package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/thefonzie-codes/goLift/internal/auth"
	"github.com/thefonzie-codes/goLift/internal/models"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var user models.User
	var passwordHash string
	err := h.db.QueryRow("SELECT id, role, password_hash FROM users WHERE email = $1", req.Email).Scan(
		&user.ID,
		&user.Role,
		&passwordHash,
	)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if !auth.CheckPassword(req.Password, passwordHash) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateToken(user.ID, user.Role, h.config.JWTSecret)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate role
	if req.Role != "coach" && req.Role != "athlete" {
		http.Error(w, "Invalid role", http.StatusBadRequest)
		return
	}

	// Check if email exists
	var exists bool
	err := h.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", req.Email).Scan(&exists)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if exists {
		http.Error(w, "Email already exists", http.StatusConflict)
		return
	}

	// Hash password
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// Insert new user
	var user models.User
	err = h.db.QueryRow(`
		INSERT INTO users (name, email, password_hash, role)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, email, role`,
		req.Name, req.Email, hashedPassword, req.Role,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Role)

	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// Generate token
	token, err := auth.GenerateToken(user.ID, user.Role, h.config.JWTSecret)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": token,
		"user":  user,
	})
} 