package repository

import (
	"database/sql"
	"errors"
)

var (
	ErrNotFound        = errors.New("not found")
	ErrFailedToAdd     = errors.New("failed to add post")
	ErrFailedToPrepare = errors.New("failed to prepare request")
)

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) *repository {
	return &repository{db: db}
}
