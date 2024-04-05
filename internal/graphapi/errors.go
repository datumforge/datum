package graphapi

import (
	"errors"
	"fmt"
)

var (
	// ErrInternalServerError is returned when an internal error occurs.
	ErrInternalServerError = errors.New("internal server error")

	// ErrPermissionDenied is returned when the user is not authorized to perform the requested query or mutation
	ErrPermissionDenied = errors.New("you are not authorized to perform this action")

	// ErrCascadeDelete is returned when an error occurs while performing cascade deletes on associated objects
	ErrCascadeDelete = errors.New("error deleting associated objects")

	// ErrSubscriberNotFound is returned when a subscriber is not found
	ErrSubscriberNotFound = errors.New("subscriber not found")
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

func newCascadeDeleteError(err error) error {
	return fmt.Errorf("%w: %v", ErrCascadeDelete, err)
}
