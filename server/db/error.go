package db

import "errors"

var (
	ErrNoID   = errors.New("ID not found with query")
	ErrExists = errors.New("ID was found")
)
