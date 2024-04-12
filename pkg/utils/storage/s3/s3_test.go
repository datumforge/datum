package s3

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
	"github.com/datumforge/datum/pkg/utils/ulids"
)

// TODO: figure out a better way to do this without env vars
var (
	accessKeyID     = os.Getenv("ACCESS_KEY_ID")
	secretAccessKey = os.Getenv("SECRET_ACCESS_KEY")
	region          = os.Getenv("AWS_REGION")
	bucket          = os.Getenv("S3_BUCKET")
)

func Test(t *testing.T) {
	keyNamespace := ulids.New().String()

	if accessKeyID == "" ||
		secretAccessKey == "" ||
		region == "" ||
		bucket == "" {
		t.SkipNow()
	}

	storageConfig := Config{
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
		Region:          region,
		Bucket:          bucket,
	}

	storage, err := NewStorage(storageConfig, keyNamespace)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	if _, err = storage.Stat(ctx, "doesnotexist"); !errors.Is(err, gostorage.ErrNotExist) {
		t.Errorf("expected not exists, got %v", err)
	}

	before := time.Now()

	var testString *string

	err = storage.Save(ctx, bytes.NewBufferString("mitb"), "ugh", "meow", "meowmeow", testString)
	require.NoError(t, err)

	now := time.Now().Add(time.Second)

	stat, err := storage.Stat(ctx, "ugh")

	require.NoError(t, err)
	assert.NotEqual(t, stat.Size, 5)
	assert.WithinDuration(t, before, now, time.Second)
}
