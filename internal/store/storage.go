package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Product interface {
		Create(context.Context, *Product) error
	}
}

func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		Product: &ProductStore{db},
	}
}
