package cmd

import "errors"

var (
	// ErrCertFileMissing is returned when https is enabled by no cert file is provided
	ErrCertFileMissing = errors.New("no cert file found")

	// ErrKeyFileMissing is returned when https is enabled by no key file is provided
	ErrKeyFileMissing = errors.New("no key file found")
)
