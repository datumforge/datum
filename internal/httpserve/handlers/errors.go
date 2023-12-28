package handlers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/datumforge/datum/internal/ent/generated"
)

var (
	// ErrBadRequest is returned when the request cannot be processed
	ErrBadRequest = errors.New("invalid request")

	// ErrProcessingRequest is returned when the request cannot be processed
	ErrProcessingRequest = errors.New("error processing request, please try again")

	// ErrMissingRequiredFields is returned when the login request has an empty username or password
	ErrMissingRequiredFields = errors.New("invalid request, missing username and/or password")

	// ErrDuplicate is returned when the request violates the unique constraints
	ErrDuplicate = errors.New("unique constraint violated on model")

	// ErrMissingRelation is returned when a foreign key restricted is violated
	ErrMissingRelation = errors.New("foreign key relation violated on model")

	// ErrNotNull is returned when a field is required but not provided
	ErrNotNull = errors.New("not null constraint violated on model")

	// ErrConstraint is returned when a database constraint is violted
	ErrConstraint = errors.New("database constraint violated")

	// ErrNotFound is returned when the requested object is not found
	ErrNotFound = errors.New("object not found in the database")
)

// InvalidEmailConfigError is returned when an invalid url configuration was provided
type InvalidEmailConfigError struct {
	// RequiredField that is missing
	RequiredField string
}

// Error returns the InvalidEmailConfigError in string format
func (e *InvalidEmailConfigError) Error() string {
	return fmt.Sprintf("invalid email url configuration: %s is required", e.RequiredField)
}

// newInvalidEmailConfigError returns an error for a missing required field in the email config
func newInvalidEmailConfigError(field string) *InvalidEmailConfigError {
	return &InvalidEmailConfigError{
		RequiredField: field,
	}
}

// MissingRequiredFieldError is returned when a required field was not provided in a request
type MissingRequiredFieldError struct {
	// RequiredField that is missing
	RequiredField string
}

// Error returns the InvalidEmailConfigError in string format
func (e *MissingRequiredFieldError) Error() string {
	return fmt.Sprintf("%s is required", e.RequiredField)
}

// newMissingRequiredField returns an error for a missing required field
func newMissingRequiredFieldError(field string) *MissingRequiredFieldError {
	return &MissingRequiredFieldError{
		RequiredField: field,
	}
}

// IsConstraintError returns true if the error resulted from a database constraint violation.
func IsConstraintError(err error) bool {
	var e *generated.ConstraintError
	return errors.As(err, &e) || IsUniqueConstraintError(err) || IsForeignKeyConstraintError(err)
}

// IsUniqueConstraintError reports if the error resulted from a DB uniqueness constraint violation.
// e.g. duplicate value in unique index.
func IsUniqueConstraintError(err error) bool {
	if err == nil {
		return false
	}

	for _, s := range []string{
		"Error 1062",                 // MySQL
		"violates unique constraint", // Postgres
		"UNIQUE constraint failed",   // SQLite
	} {
		if strings.Contains(err.Error(), s) {
			return true
		}
	}

	return false
}

// IsForeignKeyConstraintError reports if the error resulted from a database foreign-key constraint violation.
// e.g. parent row does not exist.
func IsForeignKeyConstraintError(err error) bool {
	if err == nil {
		return false
	}

	for _, s := range []string{
		"Error 1451",                      // MySQL (Cannot delete or update a parent row).
		"Error 1452",                      // MySQL (Cannot add or update a child row).
		"violates foreign key constraint", // Postgres
		"FOREIGN KEY constraint failed",   // SQLite
	} {
		if strings.Contains(err.Error(), s) {
			return true
		}
	}

	return false
}
