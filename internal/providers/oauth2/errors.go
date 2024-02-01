package oauth2

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	// ErrContextMissingToken
	ErrContextMissingToken = errors.New("oauth2: Context missing Token")

	// ErrContextMissingStateValue
	ErrContextMissingStateValue = errors.New("oauth2: Context missing state value")

	// ErrInvalidState
	ErrInvalidState = errors.New("oauth2: Invalid OAuth2 state parameter")

	// ErrFailedToGeneerateToken
	ErrFailedToGeneerateToken = errors.New("failed to generate token")

	// ErrMissingCodeOrState
	ErrMissingCodeOrState = errors.New("oauth2: Request missing code or state")
)

// DefaultFailureHandler responds with a 400 status code and message parsed from the context
var DefaultFailureHandler = http.HandlerFunc(failureHandler)

func failureHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	err := ErrorFromContext(ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// ErrorFromContext always returns some non-nil error
	http.Error(w, "", http.StatusBadRequest)
}

var (
	// ErrContextMissingErrorValue is returned when the context does not have an error value
	ErrContextMissingErrorValue = fmt.Errorf("context missing error value")

	// ErrTheMagicalNonError is an error which exists, for reasons of existing
	ErrTheMagicalNonError = fmt.Errorf("some error")
)
