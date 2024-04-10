package storage

import (
	"context"
	"encoding/json"
	"io"
	"time"
)

// Storage is the package storage interface wrapper for absctracting the various storage backends
type Storage interface {
	// Save saves content to path
	Save(ctx context.Context, content io.Reader, path string) error
	// Stat returns path metadata
	Stat(ctx context.Context, path string) (*Stat, error)
	// Open opens a reader for the content at the given path
	Open(ctx context.Context, path string) (io.ReadCloser, error)
	// OpenWithStat opens a reader for the content at the given path and returns the metadata
	OpenWithStat(ctx context.Context, path string) (io.ReadCloser, *Stat, error)
	// Delete deletes the content at the given path
	Delete(ctx context.Context, path string) error
}

// Stat contains metadata about content stored in storage
// size and last modified time aren't fetched every time for lower processing overhead
type Stat struct {
	ModifiedTime time.Time
	Size         int64
}

// Unmarshals the content of the given body and stores it in the provided pointer
func UnmarshalTo(body io.ReadCloser, to interface{}) error {
	bs, err := io.ReadAll(body)

	body.Close()

	if err != nil {
		return err
	}

	return json.Unmarshal(bs, to)
}

// unmarshalToRawMessages unmarshals the content of the given body and stores it in a juicy slice of json.RawMessage
func UnmarshalToRawMessages(body io.ReadCloser) ([]json.RawMessage, error) {
	var data []json.RawMessage
	err := UnmarshalTo(body, &data)

	return data, err
}
