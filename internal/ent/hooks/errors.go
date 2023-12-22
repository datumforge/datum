package hooks

import (
	"errors"
)

var (
	// ErrInternalServerError is returned when an internal error occurs.
	ErrInternalServerError = errors.New("internal server error")

	// ErrPersonalOrgsNoChildren is returned when personal org attempts to add a child org
	ErrPersonalOrgsNoChildren = errors.New("personal organizations are not allowed to have child organizations")
)
