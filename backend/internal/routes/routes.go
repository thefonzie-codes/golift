package routes

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/thefonzie-codes/goLift/internal/handlers"
)

func SetupRoutes(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	// Initialize handlers
	h := handlers.NewHandler(db)

	// Routes
	r.Get("/health", h.HealthCheck)
	r.Get("/users", h.GetUsers)
	r.Get("/users/{id}", h.GetUser)

	// API routes will go here...

	return r
}
