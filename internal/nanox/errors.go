package idx

import (
	"errors"
)

// ErrUnsupportedType is returned when a value is provided of an unsupported type
var ErrUnsupportedType = errors.New("unsupported type")
