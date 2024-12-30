package repository

import (
	"context"
	"ecommerce/internal/core/models"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductStore struct {
	DB *pgxpool.Pool
}

func NewProductStore(db *pgxpool.Pool) *ProductStore {
	return &ProductStore{
		DB: db,
	}
}

func (s *ProductStore) GetProducts() ([]*models.Product, error) {
	products := make([]*models.Product, 0)                                                                     // slice of poiters to product
	query := `SELECT id, name, description, price, stock_quantity, category, status, created_at FROM products` // better this instead of *

	rows, err := s.DB.Query(context.Background(), query)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("no products found")
		}
		return nil, fmt.Errorf("error to query products")
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product

		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Stock,
			&product.Category,
			&product.Status,
			&product.CreateAt,
		); err != nil {
			return nil, fmt.Errorf("error scanning products")
		}

		products = append(products, &product) // add it to the vector products
	}

	// check for errors after the iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration")
	}

	return products, nil
}

func (s *ProductStore) CreateProduct(product *models.Product) (*models.Product, error) {
	query := `INSERT INTO products (name, price, stock_quantity, category) VALUES ($1, $2, $3, $4) RETURNING id, created_at`

	err := s.DB.QueryRow(context.Background(), query,
		product.Name,
		product.Price,
		product.Stock,
		product.Category).Scan(&product.ID, &product.CreateAt)

	if err != nil {
		return nil, fmt.Errorf("unable to insert product")
	}

	return product, nil
}
