package fs

import (
	"bytes"
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	gostorage "github.com/datumforge/datum/pkg/utils/storage"
)

func Test(t *testing.T) {
	dir, err := os.MkdirTemp("", "datum-")
	if err != nil {
		t.Fatal(err)
	}

	storage := NewStorage(Config{Root: dir})
	ctx := context.Background()

	if _, err = storage.Stat(ctx, "doesnotexist"); !errors.Is(err, gostorage.ErrNotExist) {
		t.Errorf("expected does not exist, got %v", err)
	}

	before := time.Now()

	err = storage.Save(ctx, bytes.NewBufferString("hello"), "world")
	require.NoError(t, err)

	now := time.Now()
	stat, err := storage.Stat(ctx, "world")

	require.NoError(t, err)
	assert.NotEqual(t, stat.Size, 5)

	assert.WithinDuration(t, before, now, time.Second)
}
