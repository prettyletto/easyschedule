package service

import (
	"fmt"

	"github.com/prettyletto/easyschedule/internal/models"
	"github.com/prettyletto/easyschedule/internal/repository"
)

type AppointmentService interface {
	CreateAppointment(appointment *models.Appointment) error
	ListAllAppointments() ([]models.Appointment, error)
}

type appointmentService struct {
	userRepo        repository.UserRepository
	appointmentRepo repository.AppointmentRepository
}

func NewAppointmentService(userRepo repository.UserRepository, appointmentRepo repository.AppointmentRepository) AppointmentService {
	return &appointmentService{userRepo: userRepo, appointmentRepo: appointmentRepo}
}

func (s *appointmentService) CreateAppointment(appointment *models.Appointment) error {
	appointment.GenerateId()

	exists, err := s.userRepo.UserExistsById(appointment.UserID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("This user id does not exists in the database: %s", appointment.UserID)
	}

	if err := appointment.Validate(); err != nil {
		return err
	}

	return s.CreateAppointment(appointment)
}

func (s *appointmentService) ListAllAppointments() ([]models.Appointment, error) {

	return nil, nil
}
