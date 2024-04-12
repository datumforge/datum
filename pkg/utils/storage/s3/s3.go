package s3

import (
	"bytes"
	"context"
	"io"
	"mime"
	"net/http"
	"path"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/pkg/errors"
	"github.com/spf13/afero"

	"github.com/datumforge/datum/pkg/utils/storage"
)

// keyNamespace == organizationID
// key == documentID

// Storage wraps AWS S3 storage interface and s3 client pointer
type Storage struct {
	bucket       string
	s3           *s3.Client
	keyNamespace string
	uploader     *manager.Uploader
	fs           *afero.Afero
	tempFs       *afero.Afero
}

// Config is the configuration for Storage
type Config struct {
	// Enabled is a flag to enable or disable the storage
	Enabled bool `json:"enabled" koanf:"enabled" default:"false"`
	// AccessKeyID is the access key id
	AccessKeyID string `json:"accessKeyID" koanf:"accessKeyID"`
	// Bucket is the name of the bucket
	Bucket string `json:"bucket" koanf:"bucket" default:"yourbucketname"`
	// Endpoint is the endpoint to use for the s3 client
	Endpoint string `json:"endpoint" koanf:"endpoint"`
	// Region is the region to use for the s3 client
	Region string `json:"region" koanf:"region" default:"us-region-a"`
	// SecretAccessKey is the secret access key
	SecretAccessKey string `json:"secretAccessKey" koanf:"secretAccessKey"`
	// UploadConcurrency is the number of goroutines to spin up when uploading parts
	UploadConcurrency *int64 `json:"uploadConcurrency" koanf:"uploadConcurrency" default:"5"`
	// CustomHTTPClient is a custom http client wrapper for s3 interfaces
	CustomHTTPClient CustomAPIHTTPClient `json:"-" koanf:"-"`
}

// CustomAPIHTTPClient is a custom http client wrapper for s3 interfaces
type CustomAPIHTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// withUploaderConcurrency sets the concurrency of the uploader which is number of goroutines to spin up when uploading parts
func withUploaderConcurrency(concurrency int64) func(uploader *manager.Uploader) {
	return func(uploader *manager.Uploader) {
		uploader.Concurrency = int(concurrency)
	}
}

// NewStorage returns a new Storage with the provided configuration
func NewStorage(cfg Config, keyNamespace string) (*Storage, error) {
	var fs = afero.NewMemMapFs()

	var tempFs = afero.NewMemMapFs()

	awscfg := aws.Config{
		Credentials: credentials.NewStaticCredentialsProvider(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		Region:      *aws.String(cfg.Region),
	}

	var uploaderopts []func(uploader *manager.Uploader)

	client := s3.NewFromConfig(awscfg, func(o *s3.Options) {
		if cfg.Endpoint != "" {
			o.BaseEndpoint = aws.String(cfg.Endpoint)
		}

		if cfg.CustomHTTPClient != nil {
			o.HTTPClient = cfg.CustomHTTPClient
		}
	})

	if cfg.UploadConcurrency != nil {
		uploaderopts = append(uploaderopts, withUploaderConcurrency(*cfg.UploadConcurrency))
	}

	return &Storage{
		bucket:       cfg.Bucket,
		s3:           client,
		keyNamespace: keyNamespace,
		uploader:     manager.NewUploader(client, uploaderopts...),
		fs:           &afero.Afero{Fs: fs},
		tempFs:       &afero.Afero{Fs: tempFs},
	}, nil
}

// Save saves content to path inside of a bucket (bucket is set in the config)
func (s *Storage) Save(ctx context.Context, content io.Reader, path, checksum, key string, tags *string) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}

	namespacedKey := filepath.Join(s.keyNamespace, path)

	input := &s3.PutObjectInput{
		ACL:        types.ObjectCannedACLPublicRead,
		Body:       content,
		Bucket:     aws.String(s.bucket),
		Key:        aws.String(namespacedKey),
		ContentMD5: &checksum,
	}

	if tags != nil {
		input.Tagging = tags
	}

	contenttype := mime.TypeByExtension(filepath.Ext(path)) // first, detect content type from extension
	if contenttype == "" {
		// second, detect content type from first 512 bytes of content
		data := make([]byte, 512) // nolint:gomnd

		n, err := content.Read(data)
		if err != nil {
			return err
		}

		contenttype = http.DetectContentType(data)

		input.Body = io.MultiReader(bytes.NewReader(data[:n]), content)
	}

	if contenttype != "" {
		input.ContentType = aws.String(contenttype)
	}

	_, err := s.uploader.Upload(ctx, input)

	return errors.WithStack(err)
}

// Save saves content to path inside of a bucket (bucket is set in the config)
func (s *Storage) SaveQuick(ctx context.Context, content io.Reader, path string) error {
	input := &s3.PutObjectInput{
		ACL:    types.ObjectCannedACLPublicRead,
		Body:   content,
		Bucket: aws.String(s.bucket),
		Key:    aws.String(path),
	}

	contenttype := mime.TypeByExtension(filepath.Ext(path)) // first, detect content type from extension
	if contenttype == "" {
		// second, detect content type from first 512 bytes of content
		data := make([]byte, 512) // nolint:gomnd

		n, err := content.Read(data)
		if err != nil {
			return err
		}

		contenttype = http.DetectContentType(data)

		input.Body = io.MultiReader(bytes.NewReader(data[:n]), content)
	}

	if contenttype != "" {
		input.ContentType = aws.String(contenttype)
	}

	_, err := s.uploader.Upload(ctx, input)

	return errors.WithStack(err)
}

// Stat returns metadata about the object found in the provided path
func (s *Storage) Stat(ctx context.Context, path string) (*storage.Stat, error) {
	input := &s3.HeadObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(path),
	}

	out, err := s.s3.HeadObject(ctx, input)

	var notfounderr *types.NotFound

	if err != nil {
		if errors.As(err, &notfounderr) {
			return nil, storage.ErrNotExist
		}

		return nil, errors.WithStack(err)
	}

	return &storage.Stat{
		ModifiedTime: *out.LastModified,
		Size:         *out.ContentLength,
	}, nil
}

// Open opens path for reading and returns a reader which can be used to read the content
func (s *Storage) Open(ctx context.Context, path string) (io.ReadCloser, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(path),
	}

	out, err := s.s3.GetObject(ctx, input)

	var notsuckkeyerr *types.NoSuchKey

	if err != nil {
		if errors.As(err, &notsuckkeyerr) {
			return nil, storage.ErrNotExist
		}

		return nil, errors.WithStack(err)
	}

	return out.Body, nil
}

// Delete deletes path which uses the configured bucket and takes path as input
func (s *Storage) Delete(ctx context.Context, path string) error {
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(path),
	}

	_, err := s.s3.DeleteObject(ctx, input)
	if err != nil {
		return errors.WithStack(err)
	}

	return errors.WithStack(err)
}

// OpenWithStat opens the provided path for reading with file stats included
func (s *Storage) OpenWithStat(ctx context.Context, path string) (io.ReadCloser, *storage.Stat, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(path),
	}

	out, err := s.s3.GetObject(ctx, input)

	var notsuckkeyerr *types.NoSuchKey

	if err != nil {
		if errors.As(err, &notsuckkeyerr) {
			return nil, nil, storage.ErrNotExist
		}

		return nil, nil, errors.WithStack(err)
	}

	return out.Body, &storage.Stat{
		ModifiedTime: *out.LastModified,
		Size:         *out.ContentLength,
	}, nil
}

// PresignedURL returns a URL that provides access to a file for 15 minutes
func (s *Storage) PresignedURL(key string, contentType string) (string, error) {
	namespacedKey := path.Join(s.keyNamespace, key)
	presignClient := s3.NewPresignClient(s.s3)

	req, err := presignClient.PresignGetObject(context.Background(),
		&s3.GetObjectInput{
			Bucket:                     &s.bucket,
			Key:                        &namespacedKey,
			ResponseContentType:        &contentType,
			ResponseContentDisposition: StringPointer("attachment"),
		},
		func(opts *s3.PresignOptions) {
			opts.Expires = 15 * time.Minute // nolint:gomnd
		},
	)

	if err != nil {
		return "", errors.Wrap(err, "could not generate presigned URL")
	}

	return req.URL, nil
}

// Tags returns the tags for a specified key
func (s *Storage) Tags(key string) (map[string]string, error) {
	tags := make(map[string]string)

	namespacedKey := path.Join(s.keyNamespace, key)

	input := &s3.GetObjectTaggingInput{
		Bucket: &s.bucket,
		Key:    &namespacedKey,
	}

	result, err := s.s3.GetObjectTagging(context.Background(), input)
	if err != nil {
		return tags, errors.Wrap(err, "get object tagging on s3 failed")
	}

	for _, tag := range result.TagSet {
		tags[*tag.Key] = *tag.Value
	}

	return tags, nil
}

func StringPointer(s string) *string {
	return &s
}

// FileSystem returns the underlying afero filesystem
func (s *Storage) FileSystem() *afero.Afero {
	return s.fs
}

// TempFileSystem returns the temporary afero filesystem
func (s *Storage) TempFileSystem() *afero.Afero {
	return s.tempFs
}
