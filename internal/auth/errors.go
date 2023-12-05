package auth

import (
	"errors"
)

var (
	// ErrJWTMissingInvalid is returned when the JWT is missing or invalid
	ErrJWTMissingInvalid = errors.New("JWT token missing or invalid")

	// ErrJWTClaimsInvalid is returned when the token could not be cast to jwt.MapClaims
	ErrJWTClaimsInvalid = errors.New("JWT claims missing or invalid")

	// ErrSubjectNotFound is returned when the sub is not found in the JWT claims
	ErrSubjectNotFound = errors.New("JWT claims missing subject")
)
