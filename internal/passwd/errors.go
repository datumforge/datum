package passwd

import (
	"errors"
	"fmt"
)

// Error constants
var (
	// ErrCannotCreateDK is returned when the provided password is empty or the derived key creation failed
	ErrCannotCreateDK = errors.New("cannot create derived key for empty password")

	// ErrCouldNotGenerate is returned when a derived key of specified length failed to be generated
	ErrCouldNotGenerate = fmt.Errorf("could not generate %d length", dkSLen)

	// ErrUnableToVerify is returned when attempting to verify an empty derived key or empty password
	ErrUnableToVerify = errors.New("cannot verify empty derived key or password")

	// ErrCannotParseDK is returned when the encoded derived key fails to be parsed due to part(s) mismatch
	ErrCannotParseDK = errors.New("cannot parse encoded derived key, does not match regular expression")

	// ErrCannotParseEncodedEK is returned when the derived key parts do not match the desired part length
	ErrCannotParseEncodedEK = errors.New("cannot parse encoded derived key, matched expression does not contain enough subgroups")

	// ErrWeakPassword is returned when the password provided does not meet minimum complexity requirements
	ErrWeakPassword = errors.New("provided password is too weak, try adding numbers or special characters")
)

// ParseError is defining a custom error type called `ParseError`. It is a struct
// that holds intermediary values for comparison in errors
type ParseError struct {
	Object string
	Value  string
	Err    error
}

// Error returns the ParseError in string format
func (e *ParseError) Error() string {
	return fmt.Sprintf("could not parse %s %s: %v", e.Object, e.Value, e.Err)
}

func newParseError(o string, v string, err error) *ParseError {
	return &ParseError{
		Object: o,
		Value:  v,
		Err:    err,
	}
}
