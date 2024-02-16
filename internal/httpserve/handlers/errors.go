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

var (
	ErrTryLoginAgain             = response("Unable to login with those details - please try again!")
	ErrTryRegisterAgain          = response("Unable to register with those details - please try again!")
	ErrTryOrganizationAgain      = response("Unable to create or access that organization - please try again!")
	ErrTryProfileAgain           = response("Unable to create or access user profile - please try again!")
	ErrTryResendAgain            = response("Unable to resend email - please try again!")
	ErrMemberNotFound            = response("Team member with the specified ID was not found.")
	ErrMissingOrganizationName   = response("Organization name is required.")
	ErrMissingOrganizationDomain = response("Organization domain is required.")
	ErrOrganizationNotFound      = response("Organization with the specified ID was not found.")
	ErrLogBackIn                 = response("Logged out of your account - please log back in!")
	ErrVerifyEmail               = response("Please verify your email address and try again!")
	ErrInvalidEmail              = response("Please enter a valid email address.")
	ErrVerificationFailed        = response("Email verification failed. Please contact support@datum.net for assistance.")
	ErrSendPasswordResetFailed   = response("Unable to send password reset email. Please contact support@datum.net for assistance.")
	ErrPasswordResetFailed       = response("Unable to reset your password. Please contact support@datum.net for assistance.")
	ErrRequestNewInvite          = response("Invalid invitation link - please request a new one!")
	ErrSomethingWentWrong        = response("Oops - something went wrong!")
	ErrBadResendRequest          = response("Unable to resend email - please update request and try again.")
	ErrRequestNewReset           = response("Unable to reset your password - please request a new password reset.")

	AllResponses = map[string]struct{}{}
)

// response creates a standard error message to ensure uniqueness and testability for external packages
func response(msg string) string {
	if _, ok := AllResponses[msg]; ok {
		panic("duplicate error response defined: " + msg)
	}

	AllResponses[msg] = struct{}{}

	return msg
}

// StatusError decodes an error response from datum
// swagger:response StatusError
type StatusError struct {
	StatusCode int
	Reply      Reply
}

// Reply contains standard fields that are used for generic API responses and errors.
// swagger:response Reply
type Reply struct {
	Success    bool   `json:"success"`
	Error      string `json:"error,omitempty"`
	Unverified bool   `json:"unverified,omitempty"`
}
