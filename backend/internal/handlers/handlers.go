package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/thefonzie-codes/goLift/internal/models"
	"github.com/thefonzie-codes/goLift/internal/config"
)

type Handler struct {
	db     *sql.DB
	config *config.Config
}

func NewHandler(db *sql.DB, cfg *config.Config) *Handler {
	return &Handler{db: db, config: cfg}
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := h.db.Query("SELECT id, name, email, password_hash, role, specializations, self_guided FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		var specializations sql.NullString
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.PasswordHash,
			&user.Role,
			&specializations,
			&user.SelfGuided,
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if specializations.Valid {
			user.Specializations = &specializations.String
		}
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid UUID format", http.StatusBadRequest)
		return
	}

	var user models.User
	var specializations sql.NullString
	err = h.db.QueryRow("SELECT id, name, email, password_hash, role, specializations, self_guided FROM users WHERE id = $1", id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&specializations,
		&user.SelfGuided,
	)
	if err == sql.ErrNoRows {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if specializations.Valid {
		user.Specializations = &specializations.String
	}

	json.NewEncoder(w).Encode(user)
}

// Add more handlers as needed...
