package api

import (
	"errors"
	"fmt"
)

var (
	// ErrInternalServerError is returned when an internal error occurs.
	ErrInternalServerError = errors.New("internal server error")

	// ErrPermissionDenied is returned when the user is not authorized to perform the requested query or mutation
	ErrPermissionDenied = errors.New("you are not authorized to perform this action")
)

// PermissionDeniedError is returned when user is not authorized to perform the requested query or mutation
type PermissionDeniedError struct {
	Action     string
	ObjectType string
}

// Error returns the PermissionDeniedError in string format
func (e *PermissionDeniedError) Error() string {
	return fmt.Sprintf("you are not authorized to perform this action: %s on %s", e.Action, e.ObjectType)
}

// newPermissionDeniedError returns a PermissionDeniedError
func newPermissionDeniedError(a string, o string) *PermissionDeniedError {
	return &PermissionDeniedError{
		Action:     a,
		ObjectType: o,
	}
}
