package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/prettyletto/easyschedule/internal/models"
)

type UserRepository interface {
	SaveUser(user *models.User) error
	FindAllUsers() ([]models.User, error)
}

type userRepository struct {
	database *sqlx.DB
}

func New(database *sqlx.DB) UserRepository {
	return &userRepository{database: database}
}

func (r *userRepository) SaveUser(user *models.User) error {
	_, err := r.database.NamedExec(`INSERT INTO users(id,name,phone,address)
	VALUES(:id,:name,:phone,:address)`, user)
	if err != nil {
		return fmt.Errorf("Failed to insert user into database: %v", err)
	}

	return nil
}
func (r *userRepository) FindAllUsers() ([]models.User, error) {
	var users []models.User

	err := r.database.Select(&users, `SELECT id,name,phone,address FROM users`)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch users from database: %v", err)
	}

	return users, nil
}
