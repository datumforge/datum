package marionette

import (
	"errors"
	"fmt"
	"time"

	"go.uber.org/zap"
)

var (
	// ErrTaskManagerStopped
	ErrTaskManagerStopped = errors.New("the task manager is not running")

	// ErrUnschedulable
	ErrUnschedulable = errors.New("cannot schedule a task with a zero valued timestamp")

	// ErrNoWorkers
	ErrNoWorkers = errors.New("invalid configuration: at least one worker must be specified")

	// ErrNoServerName
	ErrNoServerName = errors.New("invalid configuration: no server name specified")
)

// Error keeps track of task failures and reports the failure context
type Error struct {
	err      error         // a wrapped error that describes the overall error and can be specified by the user.
	attempts int           // number of times the task was attempted
	taskerrs []error       // the error that was returned by each task for each failed retry
	duration time.Duration // the amount of time the task was tried before failure
}

// Errorw takes an error as input and returns a pointer to a new `Error` struct with the specified error set as the `err` field
func Errorw(err error) *Error {
	return &Error{err: err}
}

// Errorf returns a pointer to a new `Error` struct with the specified error message formatted as a string and arguments
func Errorf(format string, a ...any) *Error {
	return &Error{err: fmt.Errorf(format, a...)} // nolint: goerr113
}

// Error implements the error interface and gives a high level message about failures
func (e *Error) Error() string {
	if e.err != nil {
		return fmt.Sprintf("after %d attempts: %s", e.attempts, e.err.Error())
	}

	return fmt.Sprintf("task failed after %d attempts", e.attempts)
}

// Is checks if the error is the user specified target. If the wrapped user error is nil
// then it checks if the error is one of the task errors, otherwise returns false
func (e *Error) Is(target error) bool {
	if e.err != nil {
		return errors.Is(e.err, target)
	}

	for _, err := range e.taskerrs {
		if errors.Is(err, target) {
			return true
		}
	}

	return false
}

// Unwrap returns the underlying user specified error, even if it is nil
func (e *Error) Unwrap() error {
	return e.err
}

// Append adds a task failure (or nil) to the array of task errors and increment attempts
func (e *Error) Append(err error) {
	e.attempts++
	e.taskerrs = append(e.taskerrs, err)
}

// Since sets the duration of processing the task to the time since the input timestamp
func (e *Error) Since(started time.Time) {
	e.duration = time.Since(started)
}

// Dict returns a log event that can be used with the *zap.Dict method for logging
// details about the error including the number of attempts, each individual attempt
// error and the total duration of processing the task before failure.
func (e *Error) Dict() {
	logger := zap.NewExample()
	logger.Info("task error",
		zap.Dict("task event",
			zap.Error(e.err),
			zap.Int("attempts", e.attempts),
			zap.Errors("task_errors", e.taskerrs),
			zap.Duration("duration", e.duration)))
}
