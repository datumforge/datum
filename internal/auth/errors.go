package auth

import (
	"fmt"
)

// ReadError is defining a custom error type called `ReadFailed`. It is a struct
// that holds intermediary values for comparison in errors
type ReadError struct {
	Object string
	Value  string
	Err    error
}

// Error returns the ReadError in string format
func (e *ReadError) Error() string {
	return fmt.Sprintf("could not read %s %s: %v", e.Object, e.Value, e.Err)
}

func newReadError(o string, v string, err error) *ReadError {
	return &ReadError{
		Object: o,
		Value:  v,
		Err:    err,
	}
}
