package mixin

import (
	"errors"
	"fmt"

	"entgo.io/ent"
)

var (
	// ErrInvalidTokenSize is returned when session token size is invalid
	ErrInvalidTokenSize = errors.New("invalid token size")
)

// UnexpectedAuditError is returned when an unexpected audit log call is received
type UnexpectedAuditError struct {
	MutationType ent.Mutation
}

// Error returns the UnexpectedAuditError in string format
func (e *UnexpectedAuditError) Error() string {
	return fmt.Sprintf("unexpected audit log call from mutation type: %T", e.MutationType)
}

func newUnexpectedAuditError(arg ent.Mutation) *UnexpectedAuditError {
	return &UnexpectedAuditError{
		MutationType: arg,
	}
}

// UnexpectedMutationTypeError is returned when an unexpected mutation type is parsed
type UnexpectedMutationTypeError struct {
	MutationType ent.Mutation
}

// Error returns the UnexpectedAuditError in string format
func (e *UnexpectedMutationTypeError) Error() string {
	return fmt.Sprintf("unexpected mutation type: %T", e.MutationType)
}

func newUnexpectedMutationTypeError(arg ent.Mutation) *UnexpectedMutationTypeError {
	return &UnexpectedMutationTypeError{
		MutationType: arg,
	}
}
