package datumclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/datumforge/datum/pkg/models"
	"github.com/datumforge/datum/pkg/utils/responses"
)

var (
	unsuccessful = models.Reply{Success: false}
	notFound     = models.Reply{Success: false, Error: "resource not found"}
	notAllowed   = models.Reply{Success: false, Error: "method not allowed"}
	unverified   = models.Reply{Success: false, Unverified: true, Error: responses.ErrVerifyEmail}
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

// Construct a new response for an error or simply return unsuccessful.
func ErrorResponse(err interface{}) models.Reply {
	if err == nil {
		return unsuccessful
	}

	rep := models.Reply{Success: false}
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

// NotFound returns a JSON 404 response for the API.
// NOTE: we know it's weird to put server-side handlers like NotFound and NotAllowed
// here in the client/api side package but it unifies where we keep our error handling
// mechanisms.
func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, notFound)
}

// NotAllowed returns a JSON 405 response for the API.
func NotAllowed(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, notAllowed)
}

// Unverified returns a JSON 403 response indicating that the user has not verified
// their email address.
func Unverified(c *gin.Context) {
	c.JSON(http.StatusForbidden, unverified)
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

// StatusError decodes an error response from Datum.
type StatusError struct {
	StatusCode int
	Reply      models.Reply
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
