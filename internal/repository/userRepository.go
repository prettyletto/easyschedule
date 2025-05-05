package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/prettyletto/easyschedule/internal/models"
)

type UserRepository interface {
	SaveUser(user *models.User) error
	FindAllUsers() ([]models.User, error)
	UpdateUser(newUser *models.User) (*models.User, error)
	RemoveUser(id string) error
	UserExistsById(id string) (bool, error)
}

type userRepository struct {
	database *sqlx.DB
}

func NewUserRepo(database *sqlx.DB) UserRepository {
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

func (r *userRepository) UpdateUser(newUser *models.User) (*models.User, error) {
	query := `UPDATE users 
	SET name = :name,
	phone = :phone,
	address = :address
	WHERE ID = :id`

	res, err := r.database.NamedExec(query, newUser)
	if err != nil {
		return nil, fmt.Errorf("Failed to update user with %s in the database: %v", newUser.ID, err)
	}

	updatedRows, err := res.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve rows: %w", err)
	}
	if updatedRows == 0 {
		return nil, fmt.Errorf("No user found with id: %s", newUser.ID)
	}

	return newUser, nil
}

func (r *userRepository) RemoveUser(id string) error {
	query := `DELETE FROM users WHERE id = ?`
	res, err := r.database.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Failed to delete user in the database: %v", err)
	}

	deletedRows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("Failed to delete the user: %v", err)
	}
	if deletedRows == 0 {
		return fmt.Errorf("No user was found with the id: %s", id)
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

func (r *userRepository) UserExistsById(id string) (bool, error) {
	var exists bool
	err := r.database.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)", id).Scan(&exists)
	return exists, err
}
