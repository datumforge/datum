package vars

// These variables are set at build time (using -ldflags -X ...)
// See Taskfile.yaml for more details.

var (
	ClerkRootURL = "unset" //nolint:gochecknoglobals
	LoginURL     = "unset" //nolint:gochecknoglobals
	ApiURL       = "unset" //nolint:gochecknoglobals
	Stage        = "unset" //nolint:gochecknoglobals
	CommitID     = "unset" //nolint:gochecknoglobals
	Version      = "unset" //nolint:gochecknoglobals
	BuildDate    = "unset" //nolint:gochecknoglobals
	Branch       = "unset" //nolint:gochecknoglobals
)
