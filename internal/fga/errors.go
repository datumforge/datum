package fga

import (
	"errors"
	"fmt"
)

var (
	// ErrFGAMissingHost is returned when a host is not provided
	ErrFGAMissingHost = errors.New("invalid OpenFGA config: missing host")
)

// InvalidEntityError is returned when an invalid openFGA entity is configured
type InvalidEntityError struct {
	EntityRepresentation string
}

// Error returns the InvalidEntityError in string format
func (e *InvalidEntityError) Error() string {
	return fmt.Sprintf("invalid entity representation: %T", e.EntityRepresentation)
}

func newInvalidEntityError(s string) *InvalidEntityError {
	return &InvalidEntityError{
		EntityRepresentation: s,
	}
}
