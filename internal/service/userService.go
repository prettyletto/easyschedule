package service

import (
	"github.com/prettyletto/easyschedule/internal/models"
	"github.com/prettyletto/easyschedule/internal/repository"
)

type UserService interface {
	CreateUser(user *models.User) error
	ListAllUsers() ([]models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func New(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *models.User) error {
	user.GenerateId()

	return s.repo.SaveUser(user)
}

func (s *userService) ListAllUsers() ([]models.User, error) {

	return s.repo.FindAllUsers()
}
