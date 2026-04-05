package service

import (
	"database/sql"
	"strings"
	"testear/store"
)

type Service interface {
	GetUpperCaseUserName(id int) (string, error)
}

type service struct {
	store store.Store
}

func New(db *sql.DB) Service {
	return &service{
		store: store.New(db),
	}
}

func (s *service) GetUpperCaseUserName(id int) (string, error) {
	name, err := s.store.GetUserByID(id)
	if err != nil {
		return "", err
	}

	return strings.ToUpper(name), nil
}
