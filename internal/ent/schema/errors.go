package schema

import (
	"errors"
)

var (
	// ErrInvalidTokenSize is returned when session token size is invalid
	ErrInvalidTokenSize = errors.New("invalid token size")

	// ErrContainsSpaces is returned when field contains spaces
	ErrContainsSpaces = errors.New("field should not contain spaces")

	// ErrInternalServerError is returned when an internal error occurs.
	ErrInternalServerError = errors.New("internal server error")

	// ErrNotFound is returned when a resource is not found or the user does not have permissions to the resource
	ErrNotFound = errors.New("not found")
)
