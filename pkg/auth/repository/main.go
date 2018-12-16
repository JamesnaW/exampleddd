package repository

import (
	"database/sql"
)

type repo struct {
	db *sql.DB
}

// NewRepo init
func NewRepo(db *sql.DB) *repo {
	return &repo{
		db,
	}
}

// Repository interface
type Repository interface {
	Get(string) (bool, int, error)
	Check(string, string) (bool, int, error)
}
