package store

import "database/sql"

type Store interface {
	GetUserByID(id int) (string, error)
}

type store struct {
	db *sql.DB
}

func New(db *sql.DB) Store {
	return &store{db: db}
}

func (s *store) GetUserByID(id int) (string, error) {
	q := `SELECT name FROM users WHERE id = $1`

	var name string
	err := s.db.QueryRow(q, id).Scan(&name)
	if err != nil {
		return "", err
	}

	return name, nil
}
