package fs

import (
	"bytes"
	"context"
	"errors"
	"os"
	"testing"
	"time"

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

	if err := storage.Save(ctx, bytes.NewBufferString("hello"), "world"); err != nil {
		t.Fatal(err)
	}

	now := time.Now()
	stat, err := storage.Stat(ctx, "world")

	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if stat.Size != 5 {
		t.Errorf("expected size to be %d, got %d", 5, stat.Size)
	}

	if stat.ModifiedTime.Before(before) {
		t.Errorf("expected modtime to be after %v, got %v", before, stat.ModifiedTime)
	}

	if stat.ModifiedTime.After(now) {
		t.Errorf("expected modtime to be before %v, got %v", now, stat.ModifiedTime)
	}
}
