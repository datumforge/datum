package google

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	// ErrServerError
	ErrServerError = errors.New("server error")

	// ErrContextMissingGoogleUser
	ErrContextMissingGoogleUser = errors.New("context missing google user")

	// ErrFailedConstructingEndpointURL
	ErrFailedConstructingEndpointURL = errors.New("error constructing URL")

	// ErrUnableToGetGoogleUser
	ErrUnableToGetGoogleUser = errors.New("unable to get google user")

	// ErrCannotValidateGoogleUser
	ErrCannotValidateGoogleUser = errors.New("could not validate google user")
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
