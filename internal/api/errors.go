package api

import "errors"

var (
	// ErrInternalServerError is returned when an internal error occurs.
	ErrInternalServerError = errors.New("internal server error")

	ErrNotFoundError = errors.New("not found")
)
