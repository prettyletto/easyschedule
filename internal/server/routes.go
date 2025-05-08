package server

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/prettyletto/easyschedule/internal/handlers"
	"github.com/prettyletto/easyschedule/internal/repository"
	"github.com/prettyletto/easyschedule/internal/service"
)

func RegisterRoutes(mux *http.ServeMux, db *sqlx.DB) {

	userRepository := repository.NewUserRepo(db)
	userService := service.New(userRepository)
	userHandler := handlers.New(userService)

	appointmentRepository := repository.NewAppointMentRepo(db)
	appointmentService := service.NewAppointmentService(userRepository, appointmentRepository)
	appointmentHandler := handlers.NewAppointmentHandler(appointmentService)

	mux.HandleFunc("POST /users", userHandler.CreateUserHandler)
	mux.HandleFunc("GET /users", userHandler.GetAllUsers)
	mux.HandleFunc("PUT /users/{id}", userHandler.UpdateUserHandler)
	mux.HandleFunc("DELETE /users/{id}", userHandler.DeleteUserHandler)

	mux.HandleFunc("GET /appointments", appointmentHandler.GetAllAppointmentsHandler)
	mux.HandleFunc("POST /appointments", appointmentHandler.CreateUserHandler)
}
