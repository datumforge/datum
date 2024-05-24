package rout

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	echo "github.com/datumforge/echox"

	"github.com/getkin/kin-openapi/openapi3"
)

// ErrorCode is returned along side error messages for better error handling
type ErrorCode string

var (
	ErrInvalidCredentials        = errors.New("datum credentials are missing or invalid")
	ErrExpiredCredentials        = errors.New("datum credentials have expired")
	ErrPasswordMismatch          = errors.New("passwords do not match")
	ErrPasswordTooWeak           = errors.New("password is too weak: use a combination of upper and lower case letters, numbers, and special characters")
	ErrMissingID                 = errors.New("missing required id")
	ErrMissingField              = errors.New("missing required field")
	ErrInvalidField              = errors.New("invalid or unparsable field")
	ErrRestrictedField           = errors.New("field restricted for request")
	ErrConflictingFields         = errors.New("only one field can be set")
	ErrModelIDMismatch           = errors.New("resource id does not match id of endpoint")
	ErrUserExists                = errors.New("user or organization already exists")
	ErrInvalidUserClaims         = errors.New("user claims invalid or unavailable")
	ErrUnparsable                = errors.New("could not parse request")
	ErrUnknownUserRole           = errors.New("unknown user role")
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

	unsuccessful     = Reply{Success: false}
	notFound         = Reply{Success: false, Error: "resource not found"}
	notAllowed       = Reply{Success: false, Error: "method not allowed"}
	unverified       = Reply{Success: false, Unverified: true, Error: ErrVerifyEmail}
	httpunsuccessful = echo.HTTPError{}
)

// response creates a standard error message to ensure uniqueness and testability for external packages
func response(msg string) string {
	if _, ok := AllResponses[msg]; ok {
		panic("duplicate error response defined: " + msg)
	}

	AllResponses[msg] = struct{}{}

	return msg
}

// FieldError provides a general mechanism for specifying errors with specific API
// object fields such as missing required field or invalid field and giving some
// feedback about which fields are the problem
type FieldError struct {
	Field string `json:"field"`
	Err   error  `json:"error"`
}

// StatusError decodes an error response from datum
type StatusError struct {
	StatusCode int   `json:"code" yaml:"code" description:"the response HTTP code also in the response payload for your parsing convenience"`
	Reply      Reply `json:"reply" yaml:"reply" description:"the Reply generated via the internal/rout package which contains a success bool and the corresponding message"`
}

type OpenAPIErrorResponse struct {
	schemaref   openapi3.SchemaRef
	statuserror StatusError
}

// Reply contains standard fields that are used for generic API responses and errors
type Reply struct {
	Success    bool      `json:"success" yaml:"success" description:"Whether or not the request was successful or not"`
	Error      string    `json:"error,omitempty" yaml:"error,omitempty" description:"The error message if the request was unsuccessful"`
	ErrorCode  ErrorCode `json:"error_code,omitempty" yaml:"error_code,omitempty" description:"The error code if the request was unsuccessful"`
	Unverified bool      `json:"unverified,omitempty" yaml:"unverified,omitempty"`
}

type Response struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Data    any     `json:"data,omitempty"`
	Errors  []error `json:"errors,omitempty"`
}

func (res Response) MarshalJSON() ([]byte, error) {
	var errors []string
	for _, err := range res.Errors {
		errors = append(errors, err.Error())
	}

	return json.Marshal(map[string]any{"data": res.Data, "errors": errors})
}

type GraphQLRequest struct {
	Query string `json:"query"`
}

type GraphQLResponse struct {
	Data   any     `json:"data"`
	Errors []error `json:"errors,omitempty"`
}

// InvalidRequest returns a JSON 400 response for the API
func InvalidRequest() StatusError {
	return StatusError{
		StatusCode: http.StatusBadRequest,
		Reply:      Reply{Success: false, Error: "invalid request", Unverified: false},
	}
}

func InternalError() StatusError {
	return StatusError{
		StatusCode: http.StatusInternalServerError,
		Reply:      Reply{Success: false, Error: "internal server error", Unverified: false},
	}
}

// BadRequest returns a JSON 400 response for the API
func BadRequest() StatusError {
	return StatusError{
		StatusCode: http.StatusBadRequest,
		Reply:      Reply{Success: false, Error: "bad request", Unverified: false},
	}
}

// InternalServerError returns a JSON 500 response for the API
func InternalServerError() StatusError {
	return StatusError{
		StatusCode: http.StatusInternalServerError,
		Reply:      Reply{Success: false, Error: "internal server error", Unverified: false},
	}
}

// Conflict returns a JSON 409 response for the API
func Conflict() StatusError {
	return StatusError{
		StatusCode: http.StatusConflict,
		Reply:      Reply{Success: false, Error: "conflict", Unverified: false},
	}
}

// Unauthorized returns a JSON 401 response for the API
func Unauthorized() StatusError {
	return StatusError{
		StatusCode: http.StatusUnauthorized,
		Reply:      Reply{Success: false, Error: "unauthorized", Unverified: false},
	}
}

func NotFound() StatusError {
	return StatusError{
		StatusCode: http.StatusNotFound,
		Reply:      Reply{Success: false, Error: "not found", Unverified: false},
	}
}

func Created() StatusError {
	return StatusError{
		StatusCode: http.StatusCreated,
		Reply:      Reply{Success: true, Error: ""},
	}
}

// MissingRequiredFieldError is returned when a required field was not provided in a request
type MissingRequiredFieldError struct {
	// RequiredField that is missing
	RequiredField string `json:"required_field"`
}

// ErrorResponseWithCode constructs a new response for an error the contains an error code
func ErrorResponseWithCode(err interface{}, code ErrorCode) Reply {
	rep := ErrorResponse(err)
	rep.ErrorCode = code

	return rep
}

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

// NotAllowed returns a JSON 405 response for the API.
func NotAllowed(c echo.Context) error {
	return c.JSON(http.StatusMethodNotAllowed, notAllowed) //nolint:errcheck
}

// Unverified returns a JSON 403 response indicating that the user has not verified
// their email address.
func Unverified(c echo.Context) error {
	return c.JSON(http.StatusForbidden, unverified) //nolint:errcheck
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

func (e *StatusError) Error() string {
	return fmt.Sprintf("[%d] %s", e.StatusCode, e.Reply.Error)
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
