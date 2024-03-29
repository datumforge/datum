package handlers

import (
	"errors"
	"strings"

	"github.com/datumforge/datum/internal/ent/generated"
)

var (
	// ErrBadRequest is returned when the request cannot be processed
	ErrBadRequest = errors.New("invalid request")

	// ErrProcessingRequest is returned when the request cannot be processed
	ErrProcessingRequest = errors.New("error processing request, please try again")

	// ErrMissingRequiredFields is returned when the login request has an empty username or password
	ErrMissingRequiredFields = errors.New("invalid request, missing username and/or password")

	// ErrNotFound is returned when the requested object is not found
	ErrNotFound = errors.New("object not found in the database")

	// ErrMissingField is returned when a field is missing duh
	ErrMissingField = errors.New("missing required field")

	// ErrInvalidCredentials is returned when the password is invalid or missing
	ErrInvalidCredentials = errors.New("datum credentials are missing or invalid")

	// ErrUnverifiedUser is returned when email_verified on the user is false
	ErrUnverifiedUser = errors.New("user is not verified")

	// ErrUnableToVerifyEmail is returned when user's email is not able to be verified
	ErrUnableToVerifyEmail = errors.New("could not verify email")

	// ErrMaxAttempts is returned when user has requested the max retry attempts to verify their email
	ErrMaxAttempts = errors.New("max attempts verifying email address")

	// ErrNoEmailFound is returned when using an oauth provider and the email address cannot be determined
	ErrNoEmailFound = errors.New("no email found from oauth provider")

	// ErrInvalidProvider is returned when registering a user with an unsupported oauth provider
	ErrInvalidProvider = errors.New("oauth2 provider not supported")

	// ErrNoAuthUser is returned when the user couldn't be identified by the request
	ErrNoAuthUser = errors.New("could not identify authenticated user in request")

	// ErrPassWordResetTokenInvalid is returned when the provided token and secret do not match the stored
	ErrPassWordResetTokenInvalid = errors.New("password reset token invalid")

	// ErrNonUniquePassword is returned when the password was already used
	ErrNonUniquePassword = errors.New("password was already used, please try again")

	// ErrPasswordTooWeak is returned when the password is too weak
	ErrPasswordTooWeak = errors.New("password is too weak: use a combination of upper and lower case letters, numbers, and special characters")

	// ErrMaxDeviceLimit is returned when the user has reached the max device limit
	ErrMaxDeviceLimit = errors.New("max device limit reached")

	// ErrDeviceAlreadyRegistered is returned when the device is already registered
	ErrDeviceAlreadyRegistered = errors.New("device already registered")

	// ErrSubscriberNotFound is returned when the subscriber is not found
	ErrSubscriberNotFound = errors.New("subscriber not found")

	// ErrExpiredToken is returned when the token has expired
	ErrExpiredToken = errors.New("token has expired")
)

// IsConstraintError returns true if the error resulted from a database constraint violation.
func IsConstraintError(err error) bool {
	var e *generated.ConstraintError
	return errors.As(err, &e) || IsUniqueConstraintError(err) || IsForeignKeyConstraintError(err)
}

// IsUniqueConstraintError reports if the error resulted from a DB uniqueness constraint violation.
// e.g. duplicate value in unique index.
func IsUniqueConstraintError(err error) bool {
	if err == nil {
		return false
	}

	for _, s := range []string{
		"Error 1062",                 // MySQL
		"violates unique constraint", // Postgres
		"UNIQUE constraint failed",   // SQLite
	} {
		if strings.Contains(err.Error(), s) {
			return true
		}
	}

	return false
}

// IsForeignKeyConstraintError reports if the error resulted from a database foreign-key constraint violation.
// e.g. parent row does not exist.
func IsForeignKeyConstraintError(err error) bool {
	if err == nil {
		return false
	}

	for _, s := range []string{
		"Error 1451",                      // MySQL (Cannot delete or update a parent row).
		"Error 1452",                      // MySQL (Cannot add or update a child row).
		"violates foreign key constraint", // Postgres
		"FOREIGN KEY constraint failed",   // SQLite
	} {
		if strings.Contains(err.Error(), s) {
			return true
		}
	}

	return false
}
