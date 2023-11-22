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

	// ErrMissingObjectOnDeletion is returned when a object is empty in a tuple deletion
	ErrMissingObjectOnDeletion = errors.New("unable to delete tuple, missing object")

	// ErrFailedToTransformModel is returned when the FGA model cannot be transformed to JSON
	ErrFailedToTransformModel = errors.New("failed to transform fga model")
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

// WritingTuplesError is returned when an error is returned writing a relationship tuple
type WritingTuplesError struct {
	User          string
	Relation      string
	Object        string
	Operation     string
	ErrorResponse error
}

// Error returns the InvalidEntityError in string format
func (e *WritingTuplesError) Error() string {
	return fmt.Sprintf("failed to %s tuple to OpenFGA store: (user: %s; relation: %s; object: %s), error: %v", e.Operation, e.User, e.Relation, e.Object, e.ErrorResponse.Error())
}

func newWritingTuplesError(user, relation, object, operation string, err error) *WritingTuplesError {
	return &WritingTuplesError{
		User:          user,
		Relation:      relation,
		Object:        object,
		Operation:     operation,
		ErrorResponse: err,
	}
}
