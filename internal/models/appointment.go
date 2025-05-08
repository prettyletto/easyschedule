package models

import "time"

type Appointment struct {
	ID          string
	ClientToken string
	Time        time.Time
	ServiceID   string
	Status      string
}
