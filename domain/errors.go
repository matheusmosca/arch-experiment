package domain

import "errors"

var (
	ErrInvalidParams   = errors.New("invalid params")
	ErrInvalidNotFound = errors.New("resource not found")
)
