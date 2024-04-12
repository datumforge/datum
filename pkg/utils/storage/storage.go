package storage

import (
	"context"
	"crypto/md5" // nolint: gosec
	"encoding/base64"
	"encoding/json"
	"io"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/spf13/afero"
)

// Storage is the package storage interface wrapper for absctracting the various storage backends
type Storage interface {
	// Delete deletes the content at the given path
	Delete(ctx context.Context, path string) error
	// FileSystem returns the underlying filesystem
	FileSystem() *afero.Afero
	// Open opens a reader for the content at the given path
	Open(ctx context.Context, path string) (io.ReadCloser, error)
	// OpenWithStat opens a reader for the content at the given path and returns the metadata
	OpenWithStat(ctx context.Context, path string) (io.ReadCloser, *Stat, error)
	// Save saves content to path
	Save(ctx context.Context, content io.Reader, path string) error
	// Stat returns path metadata
	Stat(ctx context.Context, path string) (*Stat, error)
	// Tags returns the tags for the given path
	Tags(string) (map[string]string, error)
	// TempFileSystem returns a temporary filesystem
	TempFileSystem() *afero.Afero
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

// ComputeChecksum calculates the MD5 checksum for the provided data. It expects that
// the passed io object will be seeked to its beginning and will seek back to the
// beginning after reading its content
func ComputeChecksum(data io.ReadSeeker) (string, error) {
	hash := md5.New() // nolint: gosec
	if _, err := io.Copy(hash, data); err != nil {
		return "", ErrCouldNotReadFile
	}

	if _, err := data.Seek(0, io.SeekStart); err != nil { // seek back to beginning of file
		return "", ErrCouldNotSeekFile
	}

	return base64.StdEncoding.EncodeToString(hash.Sum(nil)), nil
}

// DetectContentType leverages http.DetectContentType to identify the content type
// of the provided data. It expects that the passed io object will be seeked to its
// beginning and will seek back to the beginning after reading its content
func DetectContentType(data io.ReadSeeker) (string, error) {
	if _, err := data.Seek(0, io.SeekStart); err != nil { // seek back to beginning of file
		return "", ErrCouldNotSeekFile
	}

	// the default return value will default to application/octet-stream if unable to detect the MIME type
	contentType, readErr := mimetype.DetectReader(data)
	if readErr != nil {
		return "", ErrCouldNotDetectContentType
	}

	if _, err := data.Seek(0, io.SeekStart); err != nil { // seek back to beginning of file
		return "", ErrCouldNotSeekFile
	}

	return contentType.String(), nil
}
