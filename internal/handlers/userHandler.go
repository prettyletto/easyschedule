package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/prettyletto/easyschedule/internal/models"
	"github.com/prettyletto/easyschedule/internal/service"
)

type UserHandler struct {
	service service.UserService
}

func New(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Error in the payload for User", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateUser(&user); err != nil {
		http.Error(w, fmt.Sprintf("Failed to create user: %v", err), http.StatusBadRequest)
		return
	}

}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.ListAllUsers()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get all users: %v", err), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application-json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
