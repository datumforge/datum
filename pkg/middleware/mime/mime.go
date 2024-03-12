package mime

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	echo "github.com/datumforge/echox"
	"github.com/datumforge/echox/middleware"
)

// Config defines the config for Mime middleware
type Config struct {
	// Skipper defines a function to skip middleware
	Enabled            bool               `json:"enabled" koanf:"enabled" default:"true"`
	Skipper            middleware.Skipper `json:"-" koanf:"-"`
	MimeTypesFile      string             `json:"mimeTypesFile" koanf:"mimeTypesFile" default:""`
	DefaultContentType string             `json:"defaultContentType" koanf:"defaultContentType" default:"application/data"`
}

// DefaultConfig is the default Gzip middleware config.
var DefaultConfig = Config{
	Skipper:            middleware.DefaultSkipper,
	MimeTypesFile:      "",
	DefaultContentType: "application/data",
}

// New creates a new middleware function with the default config
func New() echo.MiddlewareFunc {
	return NewWithConfig(DefaultConfig)
}

// NewWithConfig creates a new middleware function with the provided config
func NewWithConfig(config Config) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = DefaultConfig.Skipper
	}

	mimeTypes := loadMimeFile(config.MimeTypesFile)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			if !config.Enabled {
				return next(c)
			}

			ext := filepath.Ext(c.Request().URL.Path)
			mimeType := mimeTypes[ext]

			if mimeType == "" {
				mimeType = config.DefaultContentType
			}

			if mimeType != "" {
				c.Response().Header().Set(echo.HeaderContentType, mimeType)
			}

			return next(c)
		}
	}
}

func loadMimeFile(filename string) map[string]string {
	mimeTypes := make(map[string]string)

	f, err := os.Open(filename)
	if err != nil {
		return mimeTypes
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) <= 1 || fields[0][0] == '#' {
			continue
		}

		mimeType := fields[0]

		for _, ext := range fields[1:] {
			if ext[0] == '#' {
				break
			}

			mimeTypes[ext] = mimeType
		}
	}

	if err := scanner.Err(); err != nil {
		return mimeTypes
	}

	return mimeTypes
}
