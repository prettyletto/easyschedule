package server

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/prettyletto/easyschedule/internal/handlers"
	"github.com/prettyletto/easyschedule/internal/repository"
	"github.com/prettyletto/easyschedule/internal/service"
)

func RegisterRoutes(mux *http.ServeMux, db *sqlx.DB) {

	userRepository := repository.New(db)
	userService := service.New(userRepository)
	userHandler := handlers.New(userService)

	mux.HandleFunc("POST /users", userHandler.CreateUserHandler)
	mux.HandleFunc("GET /users", userHandler.GetAllUsers)
}
