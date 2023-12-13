package datum

import (
	"errors"
	"fmt"
)

var (
	// ErrTokenRequired is returned when no authentication token is provided
	ErrTokenRequired = errors.New("DATUM_ACCESS_TOKEN not set")
)

// RequiredFieldMissingError is returned when a field is required but not provided
type RequiredFieldMissingError struct {
	Field      string
	ObjectType string
}

// Error returns the RequiredFieldMissingError in string format
func (e *RequiredFieldMissingError) Error() string {
	return fmt.Sprintf("%s is required", e.Field)
}

// NewRequiredFieldMissingError returns an error for a missing required field
func NewRequiredFieldMissingError(f string) *RequiredFieldMissingError {
	return &RequiredFieldMissingError{
		Field: f,
	}
}
