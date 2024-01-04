package auth

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
	ErrUnauthenticated  = errors.New("request is unauthenticated")
	ErrNoClaims         = errors.New("no claims found on the request context")
	ErrNoUserInfo       = errors.New("no user info found on the request context")
	ErrInvalidAuthToken = errors.New("invalid authorization token")
	ErrAuthRequired     = errors.New("this endpoint requires authentication")
	ErrNoPermission     = errors.New("user does not have permission to perform this operation")
	ErrNoAuthUser       = errors.New("could not identify authenticated user in request")
	ErrNoAuthUserData   = errors.New("could not retrieve user data")
	ErrIncompleteUser   = errors.New("user is missing required fields")
	ErrUnverifiedUser   = errors.New("user is not verified")
	ErrCSRFVerification = errors.New("csrf verification failed for request")
	ErrParseBearer      = errors.New("could not parse bearer token from authorization header")
	ErrNoAuthorization  = errors.New("no authorization header in request")
	ErrNoRequest        = errors.New("no request found on the context")
	ErrRateLimit        = errors.New("rate limit reached: too many requests")
	ErrNoRefreshToken   = errors.New("no refresh token available on request")
	ErrRefreshDisabled  = errors.New("re-authentication with refresh tokens disabled")
	ErrShitWentBad      = errors.New("shit went bad")
)

var (
	unsuccessful = echo.HTTPError{}
	notFound     = echo.HTTPError{Code: http.StatusNotFound, Message: "resource not found"}
	notAllowed   = echo.HTTPError{Code: http.StatusMethodNotAllowed, Message: "method not allowed"}
	unverified   = echo.HTTPError{Code: http.StatusForbidden, Message: responses.ErrVerifyEmail}
	unauthorized = echo.HTTPError{Code: http.StatusUnauthorized, Message: responses.ErrTryLoginAgain}
)

var (
	ErrInvalidCredentials = errors.New("datum credentials are missing or invalid")
	ErrExpiredCredentials = errors.New("datum credentials have expired")
	ErrPasswordMismatch   = errors.New("passwords do not match")
	ErrPasswordTooWeak    = errors.New("password is too weak: use a combination of upper and lower case letters, numbers, and special characters")
	ErrNonUniquePassword  = errors.New("password was already used, please try again")
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
func ErrorResponse(err interface{}) *echo.HTTPError {
	if err == nil {
		return &unsuccessful
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

// NotFound returns a JSON 404 response for the API.
// NOTE: we know it's weird to put server-side handlers like NotFound and NotAllowed
// here in the client/api side package but it unifies where we keep our error handling
// mechanisms.
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
// feedback about which fields are the problem.
// TODO: allow multiple field errors to be specified in one response.
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
