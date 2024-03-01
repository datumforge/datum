package marionette_test

import (
	"context"
	"fmt"
	"sync"
)

// TestTask is defining the fields for a test task including failed, attempts, success, and the respective wait group
type TestTask struct {
	failUntil int
	attempts  int
	success   bool
	wg        *sync.WaitGroup
}

// Do method is responsible for executing the task
func (t *TestTask) Do(ctx context.Context) error {
	t.attempts++
	if t.attempts < t.failUntil {
		t.success = false
		return fmt.Errorf("task errored on attempt %d", t.attempts) // nolint: goerr113
	}

	t.success = true
	t.wg.Done()

	return nil
}

// String returns a string representation of the `TestTask` object, which is "test task"
func (t *TestTask) String() string {
	return "test task"
}
