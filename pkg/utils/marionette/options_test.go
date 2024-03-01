package marionette

import (
	"testing"
	"time"

	backoff "github.com/cenkalti/backoff/v4"
	"github.com/stretchr/testify/require"
)

func TestOptions(t *testing.T) {
	makeOptions := func(opts ...Option) *TaskHandler {
		tm := &TaskManager{}
		return tm.WrapTask(nil, opts...)
	}

	t.Run("Defaults", func(t *testing.T) {
		opts := makeOptions()
		require.Equal(t, 0, opts.retries)
		require.Nil(t, opts.backoff)
		require.Equal(t, time.Duration(0), opts.timeout)
		require.NotNil(t, opts.err)
		require.False(t, opts.queuedAt.IsZero())
	})

	t.Run("User", func(t *testing.T) {
		opts := makeOptions(
			WithRetries(42),
			WithBackoff(backoff.NewConstantBackOff(1*time.Second)),
			WithTimeout(1*time.Second),
			WithError(ErrTaskManagerStopped),
		)

		require.Equal(t, 42, opts.retries)
		require.IsType(t, &backoff.ConstantBackOff{}, opts.backoff)
		require.Equal(t, 1*time.Second, opts.timeout)
		require.ErrorIs(t, opts.err, ErrTaskManagerStopped)
		require.False(t, opts.queuedAt.IsZero())
	})

	t.Run("Practical", func(t *testing.T) {
		opts := makeOptions(
			WithRetries(2),
			WithErrorf("%s wicked this way comes", "something"),
		)

		require.Equal(t, 2, opts.retries)
		require.IsType(t, &backoff.ExponentialBackOff{}, opts.backoff)
		require.Equal(t, time.Duration(0), opts.timeout)
		require.EqualError(t, opts.err, "after 0 attempts: something wicked this way comes")
		require.False(t, opts.queuedAt.IsZero())
	})
}
