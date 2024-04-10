package gcs

import (
	"bytes"
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

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

	if err := storage.Save(ctx, bytes.NewBufferString("sarah"), "funkytown"); err != nil {
		t.Fatal(err)
	}

	now := time.Now()
	stat, err := storage.Stat(ctx, "funkytown")

	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if stat.Size != 5 {
		t.Errorf("expected size to be %d, got %d", 5, stat.Size)
	}

	assert.WithinDuration(t, before, now, time.Second)
}
