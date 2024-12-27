package repository

import (
	"context"
	"ecommerce/internal/core/models"
	"errors"
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
	query := `SELECT id, name, lastname, email, password, created_at FROM users WHERE email = $1`

	err := s.DB.QueryRow(context.Background(), query, email).Scan(&user.ID, &user.Name, &user.LastName, &user.Email, &user.Password, &user.CreateAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, pgx.ErrNoRows
		}
		return nil, fmt.Errorf("failed to query user") // other error
	}

	return &user, nil
}

func (s *Store) GetUserByID(id int) (*models.User, error) {
	var user models.User
	query := `SELECT id, name, lastname, email, password, created_at FROM users WHERE id = $1`

	err := s.DB.QueryRow(context.Background(), query, id).Scan(&user.ID, &user.Name, &user.LastName, &user.Email, &user.Password, &user.CreateAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to query user") // other error
	}

	return &user, nil
}

func (s *Store) CreateUser(user *models.User) error {
	query := `INSERT INTO users (name, lastname, email, password) VALUES ($1, $2, $3, $4) RETURNING id, created_at`
	err := s.DB.QueryRow(context.Background(), query,
		user.Name,
		user.LastName,
		user.Email,
		user.Password,
	).Scan(&user.ID, &user.CreateAt)

	if err != nil {
        return fmt.Errorf("error creating user: %w", err)
	}

	return nil
}
