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

// RequestError is a generic error when a request with the client fails
type RequestError struct {
	// StatusCode is the http response code that was returned
	StatusCode int
	// Body of the response
	Body string
}

// Error returns the RequestError in string format
func (e *RequestError) Error() string {
	return fmt.Sprintf("unable to process request (status %d): %s", e.StatusCode, strings.ToLower(e.Body))
}

// newRequestError returns an error when a datum client request fails
func newRequestError(statusCode int, body string) *RequestError {
	return &RequestError{
		StatusCode: statusCode,
		Body:       body,
	}
}
