package repository

import (
	"database/sql"
	"exampleddd/pkg/entity"
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
	Get(int) (entity.User, error)
}
