package fs

import (
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/datumforge/datum/pkg/utils/storage"
)

// Storage is a local filesystem storage interface
type Storage struct {
	root string
}

// Config is the configuration for Storage
type Config struct {
	// Enabled is a flag to enable or disable the storage
	Enabled bool `json:"enabled" koanf:"enabled"`
	// Root is the root directory for the filesystem storage
	Root string `json:"root" koanf:"root"`
}

// NewStorage returns a new filesystem storage with the provided configuration
func NewStorage(cfg Config) *Storage {
	return &Storage{root: cfg.Root}
}

// abs returns the absolute path of a given path
func (fs *Storage) abs(path string) string {
	return filepath.Join(fs.root, path)
}

// Save saves content to the provided path
func (fs *Storage) Save(ctx context.Context, content io.Reader, path string) error {
	abs := fs.abs(path)
	if err := os.MkdirAll(filepath.Dir(abs), 0755); err != nil { // nolint:gomnd
		return errors.WithStack(err)
	}

	w, err := os.Create(abs)
	if err != nil {
		return errors.WithStack(err)
	}

	defer w.Close()

	if _, err := io.Copy(w, content); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// Stat returns path metadata
func (fs *Storage) Stat(ctx context.Context, path string) (*storage.Stat, error) {
	fi, err := os.Stat(fs.abs(path))
	if os.IsNotExist(err) {
		return nil, storage.ErrNotExist
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	return &storage.Stat{
		ModifiedTime: fi.ModTime(),
		Size:         fi.Size(),
	}, nil
}

// Open opens path for reading
func (fs *Storage) Open(ctx context.Context, path string) (io.ReadCloser, error) {
	f, err := os.Open(fs.abs(path))
	if os.IsNotExist(err) {
		return nil, storage.ErrNotExist
	}

	return f, errors.WithStack(err)
}

// Delete deletes path
func (fs *Storage) Delete(ctx context.Context, path string) error {
	return os.Remove(fs.abs(path))
}

// OpenWithStat opens path for reading with file stats
func (fs *Storage) OpenWithStat(ctx context.Context, path string) (io.ReadCloser, *storage.Stat, error) {
	f, err := os.Open(fs.abs(path))
	if os.IsNotExist(err) {
		return nil, nil, storage.ErrNotExist
	}

	stat, err := f.Stat()
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	return f, &storage.Stat{
		ModifiedTime: stat.ModTime(),
		Size:         stat.Size(),
	}, nil
}
