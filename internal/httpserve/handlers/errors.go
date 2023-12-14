package handlers

import "errors"

var (
	// ErrBadRequest is returned when the request cannot be processed
	ErrBadRequest = errors.New("invalid request")

	// ErrMissingRequiredFields is returned when the login request has an empty username or password
	ErrMissingRequiredFields = errors.New("invalid request, missing username and/or password")
)
