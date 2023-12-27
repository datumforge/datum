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

	// ErrMissingRequiredFields is returned when the login request has an empty username or password
	ErrMissingRequiredFields = errors.New("invalid request, missing username and/or password")

	ErrDuplicate = errors.New("unique constraint violated on model")

	ErrMissingRelation = errors.New("foreign key relation violated on model")

	ErrNotNull = errors.New("not null constraint violated on model")

	ErrConstraint = errors.New("database constraint violated")

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

func constraint(dberr sqlite3.Error) *ConstraintError { //nolint:unused
	// String parsing seems to be the only way to deal with error handling for sqlite3
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
