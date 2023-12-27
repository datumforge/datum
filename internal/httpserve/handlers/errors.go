package handlers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mattn/go-sqlite3"
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

type ValidationError struct {
	err error
}

func invalid(err error) *ValidationError { //nolint:unused
	return &ValidationError{err}
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s", e.err)
}

func (e *ValidationError) Is(target error) bool {
	return errors.Is(e.err, target)
}

func (e *ValidationError) Unwrap() error {
	return e.err
}

// ConstraintError attempts to parse a sqlite3.ErrConstraint error into a model error.
type ConstraintError struct {
	err   error
	dberr sqlite3.Error
}

func newConstraintError(dberr sqlite3.Error) *ConstraintError {
	errs := dberr.Error()

	switch {
	case strings.HasPrefix(errs, "UNIQUE"):
		return &ConstraintError{err: ErrDuplicate, dberr: dberr}
	case strings.HasPrefix(errs, "FOREIGN KEY"):
		return &ConstraintError{err: ErrMissingRelation, dberr: dberr}
	case strings.HasPrefix(errs, "NOT NULL"):
		return &ConstraintError{err: ErrNotNull, dberr: dberr}
	default:
		return &ConstraintError{err: ErrConstraint, dberr: dberr}
	}
}

func (e *ConstraintError) Error() string {
	if e.dberr.Code == sqlite3.ErrConstraint {
		return e.err.Error()
	}

	return e.dberr.Error()
}

func (e *ConstraintError) Is(target error) bool {
	return errors.Is(e.err, target)
}

func (e *ConstraintError) Unwrap() error {
	return e.err
}
