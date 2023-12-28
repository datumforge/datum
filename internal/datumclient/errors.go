package datumclient

import (
	"fmt"
	"strings"
)

// AuthenticationError is returned when a user cannot be authenticated
type AuthenticationError struct {
	// StatusCode is the http response code that was returned
	StatusCode int
	// Body of the response
	Body string
}

// Error returns the AuthenticationError in string format
func (e *AuthenticationError) Error() string {
	return fmt.Sprintf("unable to authenticate (status %d): %s", e.StatusCode, strings.ToLower(e.Body))
}

// newAuthenticationError returns an error when authentication to datum fails
func newAuthenticationError(statusCode int, body string) *AuthenticationError {
	return &AuthenticationError{
		StatusCode: statusCode,
		Body:       body,
	}
}

// RegistrationError is returned when a user cannot be registered
type RegistrationError struct {
	// StatusCode is the http response code that was returned
	StatusCode int
	// Body of the response
	Body string
}

// Error returns the RegistrationError in string format
func (e *RegistrationError) Error() string {
	return fmt.Sprintf("unable to register new user (status %d): %s", e.StatusCode, strings.ToLower(e.Body))
}

// newRegistrationError returns an error when a new user cannot be registered
func newRegistrationError(statusCode int, body string) *RegistrationError {
	return &RegistrationError{
		StatusCode: statusCode,
		Body:       body,
	}
}
