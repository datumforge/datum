package fs

import (
	"bytes"
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	dir, err := os.MkdirTemp("", "datum-")
	if err != nil {
		t.Fatal(err)
	}

	storage := NewStorage(Config{Root: dir})
	ctx := context.Background()

	before := time.Now()

	err = storage.Save(ctx, bytes.NewBufferString("hello"), "world")
	require.NoError(t, err)

	now := time.Now()
	stat, err := storage.Stat(ctx, "world")

	require.NoError(t, err)
	assert.NotEqual(t, stat.Size, 5)

	assert.WithinDuration(t, before, now, time.Second)
}
