package gcs

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

// TODO: figure out a better way to do this without env vars

func Test(t *testing.T) {
	credFile := os.Getenv("CRED_FILE")
	bucket := os.Getenv("GCP_BUCKET")

	if credFile == "" || bucket == "" {
		t.SkipNow()
	}

	ctx := context.Background()

	storage, err := NewStorage(ctx, credFile, bucket)
	if err != nil {
		t.Fatal(err)
	}

	if _, err = storage.Stat(ctx, "doesnotexist"); !errors.Is(err, gostorage.ErrNotExist) {
		t.Errorf("expected not exists, got %v", err)
	}

	before := time.Now()

	err = storage.Save(ctx, bytes.NewBufferString("sarah"), "funkytown")
	require.NoError(t, err)

	now := time.Now()
	stat, err := storage.Stat(ctx, "funkytown")

	require.NoError(t, err)

	assert.NotEqual(t, stat.Size, 5)

	assert.WithinDuration(t, before, now, time.Second)
}
