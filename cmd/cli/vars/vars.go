package vars

// These variables are set at build time (using -ldflags -X ...)
// See Taskfile.yaml for more details.

var (
	ClerkRootURL = "unset" //nolint:gochecknoglobals
	LoginURL     = "unset" //nolint:gochecknoglobals
	ApiURL       = "unset" //nolint:gochecknoglobals
	Stage        = "unset" //nolint:gochecknoglobals

	// Oauth Way
	AuthURL       = "unset" //nolint:gochecknoglobals
	TokenURL      = "unset" //nolint:gochecknoglobals
	OauthClientID = "unset" //nolint:gochecknoglobals
	RedirectURL   = "unset" //nolint:gochecknoglobals
	Audience      = "unset"
)
