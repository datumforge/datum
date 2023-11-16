package fga

import (
	"errors"
	"fmt"
)

var (
	// ErrFGAMissingHost is returned when a host is not provided
	ErrFGAMissingHost = errors.New("invalid OpenFGA config: missing host")

	// ErrMissingRelation is returned when a relation is empty in a tuple creation
	ErrMissingRelation = errors.New("unable to create tuple, missing relation")

	// ErrMissingObject is returned when a object is empty in a tuple creation
	ErrMissingObject = errors.New("unable to create tuple, missing object")
)

// InvalidEntityError is returned when an invalid openFGA entity is configured
type InvalidEntityError struct {
	EntityRepresentation string
}

// Error returns the InvalidEntityError in string format
func (e *InvalidEntityError) Error() string {
	return fmt.Sprintf("invalid entity representation: %s", e.EntityRepresentation)
}

func newInvalidEntityError(s string) *InvalidEntityError {
	return &InvalidEntityError{
		EntityRepresentation: s,
	}
}
