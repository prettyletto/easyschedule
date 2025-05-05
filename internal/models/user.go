package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type User struct {
	ID      string `json:"id" db:"id"`
	Name    string `json:"name" db:"name" validate:"required"`
	Phone   string `json:"phone" db:"phone" validate:"required"`
	Address string `json:"address" db:"address"`
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

func (u *User) GenerateId() {
	u.ID = uuid.New().String()
}
