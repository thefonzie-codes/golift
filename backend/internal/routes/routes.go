package routes

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/thefonzie-codes/goLift/internal/config"
	"github.com/thefonzie-codes/goLift/internal/handlers"
	"github.com/thefonzie-codes/goLift/internal/middleware"
)

func SetupRoutes(db *sql.DB, cfg *config.Config) http.Handler {
	r := chi.NewRouter()

	// Middleware
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)
	r.Use(chimiddleware.RequestID)
	r.Use(chimiddleware.RealIP)

	h := handlers.NewHandler(db, cfg)

	// Public routes
	r.Get("/health", h.HealthCheck)
	r.Post("/login", h.Login)
	r.Post("/register", h.Register)

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(cfg))
		
		r.Get("/users", h.GetUsers)
		r.Get("/users/{id}", h.GetUser)

		// Coach-only routes
		r.Group(func(r chi.Router) {
			r.Use(middleware.RequireRole("coach"))
			// Add coach-specific routes here
		})

		// Athlete-only routes
		r.Group(func(r chi.Router) {
			r.Use(middleware.RequireRole("athlete"))
			// Add athlete-specific routes here
		})
	})

	return r
}
