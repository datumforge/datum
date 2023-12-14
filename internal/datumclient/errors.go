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

// Error returns the RequiredFieldMissingError in string format
func (e *AuthenticationError) Error() string {
	return fmt.Sprintf("unable to authenticate (status %d): %s", e.StatusCode, strings.ToLower(e.Body))
}

// newAuthenticationError returns an error for a missing required field
func newAuthenticationError(statusCode int, body string) *AuthenticationError {
	return &AuthenticationError{
		StatusCode: statusCode,
		Body:       body,
	}
}
