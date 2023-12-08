package tokens

import (
	"time"
)

type TokenConfig struct {
	Keys            map[string]string `required:"false"`                  // $DATUM_TOKEN_KEYS
	Audience        string            `default:"https://datum.net"`       // $DATUM_TOKEN_AUDIENCE
	RefreshAudience string            `required:"false"`                  // $DATUM_TOKEN_REFRESH_AUDIENCE
	Issuer          string            `default:"https://auth.datum.net"`  // $DATUM_TOKEN_ISSUER
	AccessDuration  time.Duration     `split_words:"true" default:"1h"`   // $DATUM_TOKEN_ACCESS_DURATION
	RefreshDuration time.Duration     `split_words:"true" default:"2h"`   // $DATUM_TOKEN_REFRESH_DURATION
	RefreshOverlap  time.Duration     `split_words:"true" default:"-15m"` // $DATUM_TOKEN_REFRESH_OVERLAP
}
