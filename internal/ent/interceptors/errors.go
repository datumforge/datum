package interceptors

import (
	"errors"
)

var (
	// ErrInternalServerError is returned when an internal error occurs.
	ErrInternalServerError = errors.New("internal server error")

	// ErrUnableToRetrieveUserID is returned when the user cannot be retrieved from the context
	ErrUnableToRetrieveUserID = errors.New("unable to retrieve user from context")
)
