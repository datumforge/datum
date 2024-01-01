package tokens

import "time"

// Config defines the configuration settings for authentication tokens used in the server
type Config struct {
	// Keys contains the kid as the key and a path to the pem file as the value
	KID             string            `required:"false"`                  // $DATUM_TOKEN_KID
	Keys            map[string]string `required:"false"`                  // $DATUM_TOKEN_KEYS
	Audience        string            `default:"https://datum.net"`       // $DATUM_TOKEN_AUDIENCE
	RefreshAudience string            `required:"false"`                  // $DATUM_TOKEN_REFRESH_AUDIENCE
	Issuer          string            `default:"https://auth.datum.net"`  // $DATUM_TOKEN_ISSUER
	AccessDuration  time.Duration     `split_words:"true" default:"1h"`   // $DATUM_TOKEN_ACCESS_DURATION
	RefreshDuration time.Duration     `split_words:"true" default:"2h"`   // $DATUM_TOKEN_REFRESH_DURATION
	RefreshOverlap  time.Duration     `split_words:"true" default:"-15m"` // $DATUM_TOKEN_REFRESH_OVERLAP
	CookieDomain    string            `default:"datum.net"`               // $DATUM_TOKEN_COOKIE_DOMAIN
}
