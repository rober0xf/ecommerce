package repository

import (
	"context"
	"ecommerce/internal/core/models"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	DB *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{DB: db}
}

func (s *Store) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := `SELECT id, name, email, password FROM users WHERE email = $1`

	err := s.DB.QueryRow(context.Background(), query, email).Scan(&user.ID, &user.Name, &user.LastName, &user.Email, &user.Password, &user.CreateAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to query user") // other error
	}

	return &user, nil
}

func (s *Store) GetUserByID(id int) (*models.User, error) {
	return nil, nil
}

func (s *Store) CreateUser(user *models.User) error {
	return nil
}
