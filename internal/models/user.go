package models

import "github.com/google/uuid"

type User struct {
	ID      string `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Phone   string `json:"phone" db:"phone"`
	Address string `json:"address" db:"address"`
}

func (u *User) GenerateId() {
	u.ID = uuid.New().String()
}
