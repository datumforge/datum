package marionette

import (
	"context"
	"time"

	backoff "github.com/cenkalti/backoff/v4"
)

// Option configures the task beyond the input context allowing for retries or backoff
// delays in task processing when there are failures or other task-specific handling
type Option func(*TaskHandler)

// WithRetries specifies the number of times to retry a task when it returns an error (default 0)
func WithRetries(retries int) Option {
	return func(o *TaskHandler) {
		o.retries = retries
	}
}

// WithBackoff strategy to use when retrying (default exponential backoff)
func WithBackoff(backoff backoff.BackOff) Option {
	return func(o *TaskHandler) {
		o.backoff = backoff
	}
}

// WithContext specifies a base context to be used as the parent context when the task is executed
// and on all subsequent retries
//
// NOTE: it is recommended that this context does not contain a deadline, otherwise the
// deadline may expire before the specified number of retries - use WithTimeout instead
func WithContext(ctx context.Context) Option {
	return func(o *TaskHandler) {
		o.ctx = ctx
	}
}

// WithTimeout specifies a timeout to add to the context before passing it into the task function
func WithTimeout(timeout time.Duration) Option {
	return func(o *TaskHandler) {
		o.timeout = timeout
	}
}

// WithError logs a specific error if all retries failed under the provided context. This error
// will be bundled with the errors that caused the retry failure and reported in a
// single error log message.
func WithError(err error) Option {
	return func(o *TaskHandler) {
		o.err = Errorw(err)
	}
}

// WithErrorf logs a specific error as WithError but using fmt.Errorf semantics to create the err.
func WithErrorf(format string, a ...any) Option {
	return func(o *TaskHandler) {
		o.err = Errorf(format, a...)
	}
}
