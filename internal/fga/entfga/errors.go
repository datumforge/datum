package entfga

import (
	"errors"
)

var (
	// ErrUnsupportedType is returned when the object type is not supported
	ErrUnsupportedType = errors.New("unsupported type")

	// ErrMissingRole is returned when an update request is made that contains no role
	ErrMissingRole = errors.New("missing role in update")
)
