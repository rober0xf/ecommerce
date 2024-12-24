package models

import (
	"time"
)

type User struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	LastName string    `json:"lastname"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
	CreateAt time.Time `json:"created_at"`
}

// easier to test
type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(user *User) error
}

type Payload struct {
	Name     string `json:"name" validate:"required"`
	LastName string `json:"lastname" validate:"required"`
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required,min=5,max=50"`
}
