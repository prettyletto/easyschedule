package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/prettyletto/easyschedule/internal/models"
)

type AppointmentRepository interface {
	SaveAppointment(appointment *models.Appointment) error
	FindAllAppointments() ([]models.Appointment, error)
}

type appointmentRepository struct {
	database *sqlx.DB
}

func NewAppointMentRepo(database *sqlx.DB) AppointmentRepository {
	return &appointmentRepository{database: database}
}

func (r *appointmentRepository) SaveAppointment(appointment *models.Appointment) error {

	_, err := r.database.NamedExec(`INSERT INTO appointment(
		id,user_id,service_id,team_id,start_time,
		status,canceled_reason,reschedule_count,priority) 
		VALUES(:id,:user_id,:service_id,
		:team_id,:start_time,:sstatus,
		:scanceled_reason,:sreschedule_count,:spriority 
		);`, appointment)
	if err != nil {
		return fmt.Errorf("Failed to insert appointment into database: %v", err)
	}

	return nil
}

func (r *appointmentRepository) FindAllAppointments() ([]models.Appointment, error) {
	var appointments []models.Appointment

	err := r.database.Select(&appointments, `SELECT id,user_id,team_id,start_time,status,canceled_reason,reschedule_count,priority FROM appointment`)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch appointments from database: %v", err)
	}

	return appointments, nil
}
