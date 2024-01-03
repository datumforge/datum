package entdb

import (
	"errors"
	"fmt"
)

var (
	// ErrUnsupportedMWDriver is returned when an unsupported multiwrite driver is used
	ErrUnsupportedMWDriver = errors.New("unsupported multiwrite driver")

	// ErrUnsupportedDialect is returned when an unsupported dialect is used
	ErrUnsupportedDialect = errors.New("unsupported dialect")
)

func newDialectError(dialect string) error {
	return fmt.Errorf("%w: %s", ErrUnsupportedDialect, dialect)
}

func newMultiwriteDriverError(driver string) error {
	return fmt.Errorf("%w: %s", ErrUnsupportedMWDriver, driver)
}
