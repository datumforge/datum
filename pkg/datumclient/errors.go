package datumclient

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrFailedToGetOauthToken = errors.New("failed to get oauth2 token")
	ErrNoCookieJarSet        = errors.New("client does not have a cookie jar, cannot set cookies")
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
