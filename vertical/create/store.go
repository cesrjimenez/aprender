package create

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) InsertOrder(ctx context.Context, customerID, itemID string, qty int) (string, error) {
	var id string

	q := `
		INSERT INTO orders (customer_id, item_id, quantity) 
        VALUES ($1, $2, $3) 
        RETURNING id
	`

	if err := s.db.QueryRowContext(ctx, q, customerID, itemID, qty).Scan(&id); err != nil {
		return "", fmt.Errorf("insert order failed: %w", err)
	}

	return id, nil
}
