package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID              uuid.UUID  `json:"id"`
	Name            string     `json:"name"`
	Email           string     `json:"email"`
	PasswordHash    string     `json:"-"`
	Role            string     `json:"role"`
	Specializations *string    `json:"specializations,omitempty"`
	SelfGuided      bool       `json:"self_guided"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	LastLogin       time.Time  `json:"last_login,omitempty"`
	IsActive        bool       `json:"is_active"`
	ResetPasswordToken string    `json:"-"`
}

type Program struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CoachID     uuid.UUID `json:"coach_id"`
	AthleteID   uuid.UUID `json:"athlete_id"`
	DaysPerWeek int       `json:"days_per_week"`
	TemplateID  int       `json:"template_id,omitempty"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Exercise struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Equipment   string    `json:"equipment"`
	Difficulty  string    `json:"difficulty"`
}

type AthleteMax struct {
	ID         uuid.UUID  `json:"id"`
	AthleteID  uuid.UUID  `json:"athlete_id"`
	ExerciseID uuid.UUID  `json:"exercise_id"`
	MaxWeight  float64    `json:"max_weight"`
	Date       time.Time  `json:"date"`
}

// Add more models as needed...
