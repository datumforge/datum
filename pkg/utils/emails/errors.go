package emails

import (
	"errors"
	"fmt"
)

var (
	// ErrMissingSubject is returned when the email is missing a subject
	ErrMissingSubject = errors.New("missing email subject")

	// ErrMissingSender is returned when the email sender field is missing
	ErrMissingSender = errors.New("missing email sender")

	// ErrMissingRecipient is returned when the email recipient is missing
	ErrMissingRecipient = errors.New("missing email recipient")

	// ErrEmailUnparsable is returned when an email address could not be parsed
	ErrEmailUnparsable = errors.New("could not parse email address")

	// ErrSendgridNotEnabled is returned when no sendgrid API key is present
	ErrSendgridNotEnabled = errors.New("sendgrid is not enabled, cannot add contact")

	// ErrFailedToCreateEmailClient is returned when the client cannot instantiate due to a missing API key
	ErrFailedToCreateEmailClient = errors.New("cannot create email client without API key")

	// ErrEmailArchiveOnlyInTestMode is returned when Archive is enabled but Testing mode is not enabled
	ErrEmailArchiveOnlyInTestMode = errors.New("invalid configuration: email archiving is only supported in testing mode")

	// ErrEmailNotParseable
	ErrEmailNotParseable = errors.New("invalid configuration: from email is unparsable")

	// ErrAdminEmailNotParseable
	ErrAdminEmailNotParseable = errors.New("invalid configuration: admin email is unparsable")

	// ErrBothAdminAndFromRequired
	ErrBothAdminAndFromRequired = errors.New("invalid configuration: admin and from emails are required if sendgrid is enabled")
)

// InvalidEmailConfigError is returned when an invalid url configuration was provided
type InvalidEmailConfigError struct {
	// RequiredField that is missing
	RequiredField string
}

// Error returns the InvalidEmailConfigError in string format
func (e *InvalidEmailConfigError) Error() string {
	return fmt.Sprintf("invalid email url configuration: %s is required", e.RequiredField)
}

// newInvalidEmailConfigError returns an error for a missing required field in the email config
func newInvalidEmailConfigError(field string) *InvalidEmailConfigError {
	return &InvalidEmailConfigError{
		RequiredField: field,
	}
}

// MissingRequiredFieldError is returned when a required field was not provided in a request
type MissingRequiredFieldError struct {
	// RequiredField that is missing
	RequiredField string
}

// Error returns the InvalidEmailConfigError in string format
func (e *MissingRequiredFieldError) Error() string {
	return fmt.Sprintf("%s is required", e.RequiredField)
}

// newMissingRequiredField returns an error for a missing required field
func newMissingRequiredFieldError(field string) *MissingRequiredFieldError {
	return &MissingRequiredFieldError{
		RequiredField: field,
	}
}
