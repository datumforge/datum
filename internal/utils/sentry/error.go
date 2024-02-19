package sentry

import (
	"errors"
	"fmt"
)

type ServiceError struct {
	msg  string
	args []interface{}
	err  error
}

func (e *ServiceError) Error() string {
	if e.msg == "" {
		return e.err.Error()
	}

	msg := fmt.Sprintf(e.msg, e.args...)
	return fmt.Sprintf("%s: %s", msg, e.err)
}

func (e *ServiceError) Is(target error) bool {
	return errors.Is(e.err, target)
}

func (e *ServiceError) Unwrap() error {
	return e.err
}
