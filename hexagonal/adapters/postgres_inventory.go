package adapters

import (
	"database/sql"
	"fmt"
)

type PostgresInventory struct{ db *sql.DB }

func NewPostgresInventory(db *sql.DB) *PostgresInventory {
	return &PostgresInventory{db: db}
}

func (p *PostgresInventory) Reserve(itemID string) error {
	_, err := p.db.Exec(`UPDATE inventory SET reserved = TRUE WHERE item_id = $1`, itemID)
	if err != nil {
		return fmt.Errorf("reserve failed: %w", err)
	}
	return nil
}
