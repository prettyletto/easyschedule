package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/prettyletto/easyschedule/internal/models"
	"github.com/prettyletto/easyschedule/internal/service"
)

type AppointmentHandler struct {
	service service.AppointmentService
}

func NewAppointmentHandler(appointmentService service.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{service: appointmentService}
}

func (h *AppointmentHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	var appointment models.Appointment
	if err := json.NewDecoder(r.Body).Decode(&appointment); err != nil {
		http.Error(w, "Error in the payload for appointment", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateAppointment(&appointment); err != nil {
		http.Error(w, fmt.Sprintf("Failed to create a new appointment: %v", err), http.StatusBadRequest)
		return
	}
}

func (h *AppointmentHandler) GetAllAppointmentsHandler(w http.ResponseWriter, r *http.Request) {
	appointments, err := h.service.ListAllAppointments()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get all apppointments: %v", err), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(appointments); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
