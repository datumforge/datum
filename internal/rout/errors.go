package rout

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/utils/responses"
)

var (
	unsuccessful = Reply{Success: false}
	notFound     = Reply{Success: false, Error: "resource not found"}
	notAllowed   = Reply{Success: false, Error: "method not allowed"}
	unverified   = Reply{Success: false, Unverified: true, Error: responses.ErrVerifyEmail}
)

var (
	httpunsuccessful = echo.HTTPError{}
	unauthorized     = echo.HTTPError{Code: http.StatusUnauthorized, Message: responses.ErrTryLoginAgain}
)

var (
	ErrInvalidCredentials = errors.New("datum credentials are missing or invalid")
	ErrExpiredCredentials = errors.New("datum credentials have expired")
	ErrPasswordMismatch   = errors.New("passwords do not match")
	ErrPasswordTooWeak    = errors.New("password is too weak: use a combination of upper and lower case letters, numbers, and special characters")
	ErrMissingID          = errors.New("missing required id")
	ErrMissingField       = errors.New("missing required field")
	ErrInvalidField       = errors.New("invalid or unparsable field")
	ErrRestrictedField    = errors.New("field restricted for request")
	ErrConflictingFields  = errors.New("only one field can be set")
	ErrModelIDMismatch    = errors.New("resource id does not match id of endpoint")
	ErrUserExists         = errors.New("user or organization already exists")
	ErrInvalidUserClaims  = errors.New("user claims invalid or unavailable")
	ErrUnparsable         = errors.New("could not parse request")
	ErrUnknownUserRole    = errors.New("unknown user role")
)

// ErrorResponse constructs a new response for an error or simply returns unsuccessful
func ErrorResponse(err interface{}) Reply {
	if err == nil {
		return unsuccessful
	}

	rep := Reply{Success: false}
	switch err := err.(type) {
	case error:
		rep.Error = err.Error()
	case string:
		rep.Error = err
	case fmt.Stringer:
		rep.Error = err.String()
	case json.Marshaler:
		data, e := err.MarshalJSON()
		if e != nil {
			panic(err)
		}
		rep.Error = string(data)
	default:
		rep.Error = "unhandled error response"
	}

	return rep
}

// HTTPErrorResponse constructs a new response for an error or simply returns unsuccessful
func HTTPErrorResponse(err interface{}) *echo.HTTPError {
	if err == nil {
		return &httpunsuccessful
	}

	rep := echo.HTTPError{Code: http.StatusBadRequest}
	switch err := err.(type) {
	case error:
		rep.Message = err.Error()
	case string:
		rep.Message = err
	case fmt.Stringer:
		rep.Message = err.String()
	case json.Marshaler:
		data, e := err.MarshalJSON()
		if e != nil {
			panic(err)
		}

		rep.Message = string(data)
	default:
		rep.Message = "unhandled error response"
	}

	return &rep
}

func NotFound(c echo.Context) error {
	return c.JSON(http.StatusNotFound, notFound) //nolint:errcheck
}

// NotAllowed returns a JSON 405 response for the API.
func NotAllowed(c echo.Context) error {
	return c.JSON(http.StatusMethodNotAllowed, notAllowed) //nolint:errcheck
}

// Unverified returns a JSON 403 response indicating that the user has not verified
// their email address.
func Unverified(c echo.Context) error {
	return c.JSON(http.StatusForbidden, unverified) //nolint:errcheck
}

// Unauthorized returns a JSON 401 response indicating that the request failed authorization
func Unauthorized(c echo.Context) error {
	return c.JSON(http.StatusUnauthorized, unauthorized) //nolint:errcheck
}

// FieldError provides a general mechanism for specifying errors with specific API
// object fields such as missing required field or invalid field and giving some
// feedback about which fields are the problem
type FieldError struct {
	Field string `json:"field"`
	Err   error  `json:"error"`
}

func (e *FieldError) Error() string {
	return fmt.Sprintf("%s: %s", e.Err, e.Field)
}

func (e *FieldError) Is(target error) bool {
	return errors.Is(e.Err, target)
}

func (e *FieldError) Unwrap() error {
	return e.Err
}

func MissingField(field string) error {
	return &FieldError{Field: field, Err: ErrMissingField}
}

func InvalidField(field string) error {
	return &FieldError{Field: field, Err: ErrInvalidField}
}

func RestrictedField(field string) error {
	return &FieldError{Field: field, Err: ErrRestrictedField}
}

func ConflictingFields(fields ...string) error {
	return &FieldError{Field: strings.Join(fields, ", "), Err: ErrConflictingFields}
}

// StatusError decodes an error response from datum.
type StatusError struct {
	StatusCode int
	Reply      Reply
}

// Reply contains standard fields that are used for generic API responses and errors.
type Reply struct {
	Success    bool   `json:"success"`
	Error      string `json:"error,omitempty"`
	Unverified bool   `json:"unverified,omitempty"`
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("[%d] %s", e.StatusCode, e.Reply.Error)
}

// ErrorStatus returns the HTTP status code from an error or 500 if the error is not a
// StatusError.
func ErrorStatus(err error) int {
	if err == nil {
		return http.StatusOK
	}

	if e, ok := err.(*StatusError); !ok || e.StatusCode < 100 || e.StatusCode >= 600 {
		return http.StatusInternalServerError
	} else {
		return e.StatusCode
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

// NewMissingRequiredFieldError returns an error for a missing required field
func NewMissingRequiredFieldError(field string) *MissingRequiredFieldError {
	return &MissingRequiredFieldError{
		RequiredField: field,
	}
}
