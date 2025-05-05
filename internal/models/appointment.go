package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Appointment struct {
	ID              string    `json:"id" db:"id"`
	UserID          string    `json:"user_id" db:"user_id" validate:"required"`
	ServiceID       string    `json:"service_id" db:"service_id" validate:"required"`
	TeamID          string    `json:"team_id" db:"team_id" validate:"required"`
	StartTime       time.Time `json:"start_time" db:"start_time" validate:"required"`
	Status          string    `json:"status" db:"status"`
	CanceledReason  *string   `json:"canceled_reason,omitempty" db:"canceled_reason,omitempty"`
	RescheduleCount int       `json:"reschedule_count" db:"reschedule_count"`
	Priority        bool      `json:"priority" db:"priority"`
}

func (a *Appointment) Validate() error {
	validate := validator.New()
	return validate.Struct(a)
}

func (a *Appointment) GenerateId() {
	a.ID = uuid.New().String()
}
