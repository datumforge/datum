package config

import (
	"github.com/datumforge/datum/pkg/utils/storage/fs"
	"github.com/datumforge/datum/pkg/utils/storage/gcs"
	"github.com/datumforge/datum/pkg/utils/storage/s3"
)

// Config is the configuration for the storage backend
// This exists primarily for ease of use with koanf
type Config struct {
	S3config  *s3.Config  `json:"s3" koanf:"s3"`
	GCSconfig *gcs.Config `json:"gcs" koanf:"gcs"`
	FSconfig  *fs.Config  `json:"fs" koanf:"fs"`
}
