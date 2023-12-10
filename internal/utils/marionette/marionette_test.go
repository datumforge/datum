package marionette_test

import (
	"context"
	"os"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	backoff "github.com/cenkalti/backoff/v4"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/datumforge/datum/internal/utils/marionette"
)

// TestMain is a special function in Go that is used to control the execution of test functions.
// It is called by the testing framework before running any tests. In this case, it is used to run the
// tests and exit the program with the appropriate exit code.
func TestMain(m *testing.M) {
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestTasks(t *testing.T) {
	// NOTE: ensure the queue size is zero so that queueing blocks until all tasks are
	// queued to prevent a race condition with the call to stop.
	logger := zap.NewNop().Sugar()
	tm := marionette.New(marionette.Config{Logger: logger, Workers: 4, QueueSize: 0, ServerName: "test"})
	tm.Start()

	// Should be able to call start twice without panic
	tm.Start()

	// Queue basic tasks with no retries
	var completed int32

	for i := 0; i < 100; i++ {
		tm.Queue(marionette.TaskFunc(func(context.Context) error { // nolint: errcheck
			time.Sleep(1 * time.Millisecond)
			atomic.AddInt32(&completed, 1)
			return nil
		}))
	}

	require.True(t, tm.IsRunning())
	tm.Stop()

	// Should be able to call stop twice without panic
	tm.Stop()

	require.Equal(t, int32(100), completed)
	require.False(t, tm.IsRunning())

	// Should not be able to queue when the task manager is stopped
	err := tm.Queue(marionette.TaskFunc(func(context.Context) error { return nil }))
	require.ErrorIs(t, err, marionette.ErrTaskManagerStopped)
}

func TestTasksRetry(t *testing.T) {
	// This is a long running test, skip if in short mode
	if testing.Short() {
		t.Skip("skipping long running test in short mode")
	}

	logger := zap.NewNop().Sugar()

	// NOTE: ensure the queue size is zero so that queueing blocks until all tasks are
	// queued to prevent a race condition with the call to stop.
	tm := marionette.New(marionette.Config{Logger: logger, Workers: 4, QueueSize: 0, ServerName: "test"})
	tm.Start()

	// Create a state of tasks that hold the number of attempts and success
	var wg sync.WaitGroup

	state := make([]*TestTask, 0, 100)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		state = append(state, &TestTask{failUntil: 3, wg: &wg})
	}

	// Queue state tasks with a retry limit that will ensure they all succeed
	for _, retryTask := range state {
		tm.Queue(retryTask, marionette.WithRetries(5), marionette.WithBackoff(&backoff.ZeroBackOff{})) // nolint: errcheck
	}

	// Wait for all tasks to be completed and stop the task manager.
	wg.Wait()
	tm.Stop()

	// Analyze the results from the state
	var completed, attempts int
	for _, retryTask := range state {
		attempts += retryTask.attempts

		if retryTask.success {
			completed++
		}
	}

	require.Equal(t, 100, completed, "expected all tasks to have been completed")
	require.Equal(t, 300, attempts, "expected all tasks to have failed twice before success")
}

func TestTasksRetryFailure(t *testing.T) {
	// This is a long running test, skip if in short mode
	if testing.Short() {
		t.Skip("skipping long running test in short mode")
	}

	logger := zap.NewNop().Sugar()

	// NOTE: ensure the queue size is zero so that queueing blocks until all tasks are
	// queued to prevent a race condition with the call to stop.
	tm := marionette.New(marionette.Config{Logger: logger, Workers: 4, QueueSize: 0, ServerName: "test"})
	tm.Start()

	// Create a state of tasks that hold the number of attempts and success
	var wg sync.WaitGroup

	state := make([]*TestTask, 0, 100)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		state = append(state, &TestTask{failUntil: 5, wg: &wg})
	}

	// Queue state tasks with a retry limit that will ensure they all fail
	for _, retryTask := range state {
		tm.Queue(retryTask, marionette.WithRetries(1), marionette.WithBackoff(&backoff.ZeroBackOff{})) // nolint: errcheck
	}

	// Wait for all tasks to be completed and stop the task manager.
	time.Sleep(500 * time.Millisecond)
	tm.Stop()

	// Analyze the results from the state
	var completed, attempts int
	for _, retryTask := range state {
		attempts += retryTask.attempts

		if retryTask.success {
			completed++
		}
	}

	require.Equal(t, 0, completed, "expected all tasks to have failed")
	require.Equal(t, 200, attempts, "expected all tasks to have failed twice before no more retries")
}

func TestTasksRetryBackoff(t *testing.T) {
	logger := zap.NewNop().Sugar()

	// NOTE: ensure the queue size is zero so that queueing blocks until all tasks are
	// queued to prevent a race condition with the call to stop.
	tm := marionette.New(marionette.Config{Logger: logger, Workers: 4, QueueSize: 0, ServerName: "test"})
	tm.Start()

	// Create a state of tasks that hold the number of attempts and success
	var wg sync.WaitGroup

	state := make([]*TestTask, 0, 100)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		state = append(state, &TestTask{failUntil: 3, wg: &wg})
	}

	// Queue state tasks with a retry limit that will ensure they all succeed
	for _, retryTask := range state {
		tm.Queue(retryTask, marionette.WithRetries(5), marionette.WithBackoff(backoff.NewConstantBackOff(10*time.Millisecond))) // nolint: errcheck
	}

	// Wait for all tasks to be completed and stop the task manager.
	wg.Wait()
	tm.Stop()

	// Analyze the results from the state
	var completed, attempts int
	for _, retryTask := range state {
		attempts += retryTask.attempts

		if retryTask.success {
			completed++
		}
	}

	// TODO: how to check if backoff was respected?
	require.Equal(t, 100, completed, "expected all tasks to have been completed")
	require.Equal(t, 300, attempts, "expected all tasks to have failed twice before success")
}

func TestTasksRetryContextCanceled(t *testing.T) {
	// This is a long running test, skip if in short mode
	if testing.Short() {
		t.Skip("skipping long running test in short mode")
	}

	logger := zap.NewNop().Sugar()

	// NOTE: ensure the queue size is zero so that queueing blocks until all tasks are
	// queued to prevent a race condition with the call to stop.

	tm := marionette.New(marionette.Config{Logger: logger, Workers: 4, QueueSize: 0, ServerName: "test"})
	tm.Start()

	var completed, attempts int32

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	// Queue tasks that are getting canceled
	for i := 0; i < 100; i++ {
		tm.Queue(marionette.TaskFunc(func(ctx context.Context) error { // nolint: errcheck
			atomic.AddInt32(&attempts, 1)
			if err := ctx.Err(); err != nil {
				return err
			}

			atomic.AddInt32(&completed, 1)
			return nil
		}), marionette.WithRetries(1),
			marionette.WithBackoff(&backoff.ZeroBackOff{}),
			marionette.WithContext(ctx),
		)
	}

	// Wait for all tasks to be completed and stop the task manager.
	time.Sleep(500 * time.Millisecond)
	tm.Stop()

	require.Equal(t, int32(0), completed, "expected all tasks to have been canceled")
	require.Equal(t, int32(200), attempts, "expected all tasks to have failed twice before no more retries")
}

func TestTasksRetrySuccessAndFailure(t *testing.T) {
	// This is a long running test, skip if in short mode
	if testing.Short() {
		t.Skip("skipping long running test in short mode")
	}

	logger := zap.NewNop().Sugar()

	// Test non-retry tasks alongside retry tasks
	// NOTE: ensure the queue size is zero so that queueing blocks until all tasks are
	// queued to prevent a race condition with the call to stop.
	tm := marionette.New(marionette.Config{Logger: logger, Workers: 4, QueueSize: 0, ServerName: "test"})
	tm.Start()

	// Create a state of tasks that hold the number of attempts and success
	var wg sync.WaitGroup

	state := make([]*TestTask, 0, 100)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		state = append(state, &TestTask{failUntil: 2, wg: &wg})
	}

	// Queue state tasks with a retry limit that will ensure they all fail
	// First 50 have a retry, second 50 do not.
	for i, retryTask := range state {
		if i < 50 {
			tm.Queue(retryTask, marionette.WithRetries(2), marionette.WithBackoff(&backoff.ZeroBackOff{})) // nolint: errcheck
		} else {
			tm.Queue(retryTask, marionette.WithBackoff(&backoff.ZeroBackOff{})) // nolint: errcheck
		}
	}

	// Wait for all tasks to be completed and stop the task manager.
	time.Sleep(500 * time.Millisecond)
	tm.Stop()

	// Analyze the results from the state
	var completed, attempts int
	for _, retryTask := range state {
		attempts += retryTask.attempts

		if retryTask.success {
			completed++
		}
	}

	require.Equal(t, 50, completed, "expected all tasks to have failed")
	require.Equal(t, 150, attempts, "expected all tasks to have failed twice before no more retries")
}

func TestQueue(t *testing.T) {
	// A simple test to ensure that tm.Stop() will wait until all items in the queue are finished.
	var wg sync.WaitGroup

	queue := make(chan int32, 64)

	var final int32

	wg.Add(1)

	go func() {
		for num := range queue {
			time.Sleep(1 * time.Millisecond)
			atomic.SwapInt32(&final, num)
		}

		wg.Done()
	}()

	for i := int32(1); i < 101; i++ {
		queue <- i
	}

	close(queue)
	wg.Wait()
	require.Equal(t, int32(100), final)
}
