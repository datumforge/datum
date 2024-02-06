package datum

import (
	"errors"
	"fmt"
)

var (
	// ErrTokenRequired is returned when no authentication token is provided
	ErrTokenRequired = errors.New("DATUM_ACCESS_TOKEN not set")

	// ErrInvalidRole is returned when an invalid role is provided for a member
	ErrInvalidRole = errors.New("invalid role, only member and admin are allowed")

	// ErrInvalidInviteStatus is returned when an invalid status is provided for an invite
	ErrInvalidInviteStatus = errors.New("invalid status, only sent, required, accepted, expired are allowed")

	// ErrUnsupportedProvider is returned when an invalid provider is specified during login
	ErrUnsupportedProvider = errors.New("invalid provider, only Github and Google are supported")
)

// RequiredFieldMissingError is returned when a field is required but not provided
type RequiredFieldMissingError struct {
	// Field contains the required field that was missing from the input
	Field string
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
