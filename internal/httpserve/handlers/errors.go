package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/datumforge/datum/internal/ent/generated"

	echo "github.com/datumforge/echox"
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

	// ErrConstraint is returned when a database constraint is violated
	ErrConstraint = errors.New("database constraint violated")

	// ErrNotFound is returned when the requested object is not found
	ErrNotFound = errors.New("object not found in the database")

	// ErrMissingField is returned when a field is missing duh
	ErrMissingField = errors.New("missing required field")

	// ErrInvalidCredentials is returned when the password is invalid or missing
	ErrInvalidCredentials = errors.New("datum credentials are missing or invalid")

	// ErrUnverifiedUser is returned when email_verified on the user is false
	ErrUnverifiedUser = errors.New("user is not verified")

	// ErrUnableToVerifyEmail is returned when user's email is not able to be verified
	ErrUnableToVerifyEmail = errors.New("could not verify email")

	// ErrNoAuthUser is returned when the user couldn't be identified by the request
	ErrNoAuthUser = errors.New("could not identify authenticated user in request")

	unsuccessful = echo.HTTPError{}
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

// ErrorResponse constructs a new response for an error or simply returns unsuccessful
func ErrorResponse(err interface{}) *echo.HTTPError {
	if err == nil {
		return &unsuccessful
	}

	rep := echo.HTTPError{Code: http.StatusBadRequest}
	switch err := err.(type) {
	case error:
		rep.Message = err.Error()
	case string:
		rep.Message = err
	case fmt.Stringer:
		rep.Message = err.String()
	case json.Marshaler:
		data, e := err.MarshalJSON()
		if e != nil {
			panic(err)
		}

		rep.Message = string(data)
	default:
		rep.Message = "unhandled error response"
	}

	return &rep
}
