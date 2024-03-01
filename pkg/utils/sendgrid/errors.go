package sendgrid

import (
	"errors"
)

var (
	ErrNoRecipientFound = errors.New("no recipient found")
)
