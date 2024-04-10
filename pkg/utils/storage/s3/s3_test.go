package s3

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
var (
	accessKeyID     = os.Getenv("ACCESS_KEY_ID")
	secretAccessKey = os.Getenv("SECRET_ACCESS_KEY")
	region          = os.Getenv("AWS_REGION")
	bucket          = os.Getenv("S3_BUCKET")
)

func Test(t *testing.T) {
	if accessKeyID == "" ||
		secretAccessKey == "" ||
		region == "" ||
		bucket == "" {
		t.SkipNow()
	}

	storage, err := NewStorage(Config{
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
		Region:          region,
		Bucket:          bucket,
	})
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	if _, err = storage.Stat(ctx, "doesnotexist"); !errors.Is(err, gostorage.ErrNotExist) {
		t.Errorf("expected not exists, got %v", err)
	}

	before := time.Now()

	if err := storage.Save(ctx, bytes.NewBufferString("mitb"), "ugh"); err != nil {
		t.Fatal(err)
	}

	now := time.Now().Add(time.Second)

	stat, err := storage.Stat(ctx, "ugh")

	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if stat.Size != 5 {
		t.Errorf("expected size to be %d, got %d", 5, stat.Size)
	}

	assert.WithinDuration(t, before, now, time.Second)
}
