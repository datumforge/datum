package gcs

import (
	"context"
	"io"
	"mime"
	"path/filepath"

	"cloud.google.com/go/storage"
	"github.com/pkg/errors"
	"google.golang.org/api/option"

	gostorages "github.com/datumforge/datum/pkg/utils/storage"
)

// Storage is a Google Cloud Storage interface
type Storage struct {
	bucket *storage.BucketHandle
}

// Config is the configuration for Storage - need to blow this out but initial focus was on s3
type Config struct {
	// Enabled is a flag to enable or disable the storage
	Enabled bool `json:"enabled" koanf:"enabled"`
	// CredentialsFile is the path to the credentials file
	CredentialsFile string `json:"credentialsFile" koanf:"credentialsFile"`
	// Bucket is the name of the bucket
	Bucket string `json:"bucket" koanf:"bucket"`
}

// NewStorage returns a new GCP Storage with the provided configuration
func NewStorage(ctx context.Context, credentialsFile, bucket string) (*Storage, error) {
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Storage{bucket: client.Bucket(bucket)}, nil
}

// Save saves content to path
func (g *Storage) Save(ctx context.Context, content io.Reader, path string) (rerr error) {
	w := g.bucket.Object(path).NewWriter(ctx)
	w.ContentType = mime.TypeByExtension(filepath.Ext(path))

	// Close the writer on return
	defer func() {
		if err := w.Close(); err != nil {
			rerr = err
		}
	}()

	if _, err := io.Copy(w, content); err != nil {
		return errors.WithStack(err)
	}

	return errors.WithStack(rerr)
}

// Stat returns path metadata
func (g *Storage) Stat(ctx context.Context, path string) (*gostorages.Stat, error) {
	attrs, err := g.bucket.Object(path).Attrs(ctx)

	if errors.Is(err, storage.ErrObjectNotExist) {
		return nil, gostorages.ErrNotExist
	} else if err != nil {
		return nil, err
	}

	return &gostorages.Stat{
		ModifiedTime: attrs.Updated,
		Size:         attrs.Size,
	}, nil
}

// Open opens path for reading
func (g *Storage) Open(ctx context.Context, path string) (io.ReadCloser, error) {
	r, err := g.bucket.Object(path).NewReader(ctx)
	if errors.Is(err, storage.ErrObjectNotExist) {
		return nil, gostorages.ErrNotExist
	}

	return r, errors.WithStack(err)
}

// Delete deletes path
func (g *Storage) Delete(ctx context.Context, path string) error {
	return errors.WithStack(g.bucket.Object(path).Delete(ctx))
}

// OpenWithStat opens path for reading with file stats
func (g *Storage) OpenWithStat(ctx context.Context, path string) (io.ReadCloser, *gostorages.Stat, error) {
	r, err := g.bucket.Object(path).NewReader(ctx)

	if errors.Is(err, storage.ErrObjectNotExist) {
		return nil, nil, gostorages.ErrNotExist
	}

	return r, &gostorages.Stat{
		ModifiedTime: r.Attrs.LastModified,
		Size:         r.Attrs.Size,
	}, errors.WithStack(err)
}
