package store

import (
	"context"
	"database/sql"
)

type Product struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       int64  `json:"stock"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	IsActive    string `json:"is_active"`
}

type ProductStore struct {
	db *sql.DB
}

func (p *ProductStore) Create(ctx context.Context, product *Product) error {
	query := `
	INSERT INTO products (name, description, stock, is_active)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at, updated_at
	`
	err := p.db.QueryRowContext(
		ctx,
		query,
		product.Name,
		product.Description,
		product.Stock,
		product.IsActive,
	).Scan(
		&product.ID,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
