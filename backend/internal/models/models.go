package models

import "time"

type User struct {
	ID                 int       `json:"id"`
	Name               string    `json:"name"`
	Email              string    `json:"email"`
	PasswordHash       string    `json:"-"`
	Role               string    `json:"role"`
	Specializations    string    `json:"specializations,omitempty"`
	SelfGuided         bool      `json:"self_guided"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	LastLogin          time.Time `json:"last_login,omitempty"`
	IsActive           bool      `json:"is_active"`
	ResetPasswordToken string    `json:"-"`
}

type Program struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CoachID     int       `json:"coach_id"`
	AthleteID   int       `json:"athlete_id"`
	DaysPerWeek int       `json:"days_per_week"`
	TemplateID  int       `json:"template_id,omitempty"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Add more models as needed...
