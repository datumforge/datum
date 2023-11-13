package api

import "errors"

var (
	// ErrInternalServerError is returned when an internal error occurs.
	ErrInternalServerError = errors.New("internal server error")

	// ErrNotFound is returned when a resource is not found or the user does not have permissions to the resource
	ErrNotFound = errors.New("not found")
)
