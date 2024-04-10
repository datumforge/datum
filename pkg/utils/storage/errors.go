package storage

import (
	"errors"
)

var (
	// ErrNotExist is returned when a path does not exist
	ErrNotExist = errors.New("does not exist")
)
