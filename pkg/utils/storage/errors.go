package storage

import (
	"errors"
)

var (
	// ErrNotExist is returned when a path does not exist
	ErrNotExist = errors.New("does not exist")

	// ErrCouldNotSeekFile is returned when a file could not be seeked
	ErrCouldNotSeekFile = errors.New("could not seek file")

	// ErrCouldNotReadFile is returned when a file could not be read
	ErrCouldNotReadFile = errors.New("could not read file")

	// ErrCouldNotWriteFile is returned when a file could not be written
	ErrCouldNotWriteFile = errors.New("could not write file")

	// ErrCouldNotDetectContentType is returned when a content type could not be detected
	ErrCouldNotDetectContentType = errors.New("could not detect content type")
)
