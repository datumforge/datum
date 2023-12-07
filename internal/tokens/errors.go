package tokens

import (
	"errors"
	"fmt"
)

// Error constants
var (
	ErrInvalidKey      = errors.New("key is invalid")
	ErrInvalidKeyType  = errors.New("key is of invalid type")
	ErrHashUnavailable = errors.New("the requested hash function is unavailable")

	ErrCacheMiss    = errors.New("requested key is not in the cache")
	ErrCacheExpired = errors.New("requested key is expired")

	// ErrTokenManagerFailedInit returns when the token manager was not correctly provided signing keys
	ErrTokenManagerFailedInit = errors.New("token manager not initialized with signing keys")

	// ErrFailedRetreiveClaimsFromToken returns when claims can not be retreived from an access token
	ErrFailedRetreiveClaimsFromToken = errors.New("could not retrieve claims from access token")

	// ErrTokenMissingKid returns when the kid cannot be found in the header of the token
	ErrTokenMissingKid = errors.New("token does not have kid in header")

	// ErrFailedParsingKid returns when the kid could not be parsed
	ErrFailedParsingKid = errors.New("could not parse kid: %s")

	// ErrUnknownSigningKey returns when the signing key fetched does not match the loaded managed keys
	ErrUnknownSigningKey = errors.New("unknown signing key")
)

var (
	// The below block of error constants are only used in comparison(s) checks with ValidationError and are the inner standard errors which could occur when performing token validation; these won't be referenced in return error messages within the code directly and wrapped with our custom errors

	// ErrTokenMalformed returns when a token is malformed
	ErrTokenMalformed = errors.New("token is malformed")

	// ErrTokenUnverifiable
	ErrTokenUnverifiable = errors.New("token is unverifiable")

	// ErrTokenSignatureInvalid
	ErrTokenSignatureInvalid = errors.New("token signature is invalid")

	// ErrTokenInvalidAudience
	ErrTokenInvalidAudience = errors.New("token has invalid audience")

	// ErrTokenExpired
	ErrTokenExpired = errors.New("token is expired")

	// ErrTokenUsedBeforeIssued
	ErrTokenUsedBeforeIssued = errors.New("token used before issued")

	// ErrTokenInvalidIssuer
	ErrTokenInvalidIssuer = errors.New("token has invalid issuer")

	// ErrTokenNotValidYet
	ErrTokenNotValidYet = errors.New("token is not valid yet")

	// ErrTokenNotValid
	ErrTokenNotValid = errors.New("token is invalid")

	// ErrTokenInvalidID
	ErrTokenInvalidID = errors.New("token has invalid id")

	// ErrTokenInvalidClaims
	ErrTokenInvalidClaims = errors.New("token has invalid claims")
)

// The errors that might occur when parsing and validating a token
const (
	ValidationErrorMalformed        uint32 = 1 << iota // Token is malformed
	ValidationErrorUnverifiable                        // Token could not be verified because of signing problems
	ValidationErrorSignatureInvalid                    // Signature validation failed

	// Standard Claim validation errors
	ValidationErrorAudience      // AUD validation failed
	ValidationErrorExpired       // EXP validation failed
	ValidationErrorIssuedAt      // IAT validation failed
	ValidationErrorIssuer        // ISS validation failed
	ValidationErrorNotValidYet   // NBF validation failed
	ValidationErrorID            // JTI validation failed
	ValidationErrorClaimsInvalid // Generic claims validation error
)

// NewValidationError is a helper for constructing a ValidationError with a string error message
func NewValidationError(errorText string, errorFlags uint32) *ValidationError {
	return &ValidationError{
		text:   errorText,
		Errors: errorFlags,
	}
}

// ValidationError represents an error from Parse if token is not valid
type ValidationError struct {
	Inner  error
	Errors uint32
	text   string
}

// Error is the implementation of the err interface for ValidationError
func (e ValidationError) Error() string {
	i := e.Inner

	switch {
	case i != nil:
		return e.Inner.Error()
	case e.text != "":
		return e.text
	default:
		return "token is invalid"
	}
}

// Unwrap gives errors.Is and errors.As access to the inner errors defined above
func (e *ValidationError) Unwrap() error {
	return e.Inner
}

// Is checks if this ValidationError is of the supplied error. We are first checking for the exact
// error message by comparing the inner error message. If that fails, we compare using the error
// flags. This way we can use custom error messages and leverage errors.Is using the global error
// variables, plus I just learned how to use errors.Is today so this is pretty sweet
func (e *ValidationError) Is(err error) bool {
	// Check, if our inner error is a direct match
	if errors.Is(errors.Unwrap(e), err) {
		return true
	}

	// Otherwise, we need to match using our error flags
	switch err {
	case ErrTokenMalformed:
		return e.Errors&ValidationErrorMalformed != 0
	case ErrTokenUnverifiable:
		return e.Errors&ValidationErrorUnverifiable != 0
	case ErrTokenSignatureInvalid:
		return e.Errors&ValidationErrorSignatureInvalid != 0
	case ErrTokenInvalidAudience:
		return e.Errors&ValidationErrorAudience != 0
	case ErrTokenExpired:
		return e.Errors&ValidationErrorExpired != 0
	case ErrTokenUsedBeforeIssued:
		return e.Errors&ValidationErrorIssuedAt != 0
	case ErrTokenInvalidIssuer:
		return e.Errors&ValidationErrorIssuer != 0
	case ErrTokenNotValidYet:
		return e.Errors&ValidationErrorNotValidYet != 0
	case ErrTokenInvalidID:
		return e.Errors&ValidationErrorID != 0
	case ErrTokenInvalidClaims:
		return e.Errors&ValidationErrorClaimsInvalid != 0
	}

	return false
}

// ParseError is defining a custom error type called `ParseError`
type ParseError struct {
	Object string
	Value  string
	Err    error
}

// Error returns the ParseError in string format
func (e *ParseError) Error() string {
	return fmt.Sprintf("could not parse %s %s: %v", e.Object, e.Value, e.Err)
}

// The function newParseError creates a new ParseError object with the given parameters
func newParseError(o string, v string, err error) *ParseError {
	return &ParseError{
		Object: o,
		Value:  v,
		Err:    err,
	}
}
