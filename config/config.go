package config

import (
	"fmt"
	"net/url"
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

type URLConfig struct {
	Base   string `split_words:"true" default:"https://datum.net"`
	Verify string `split_words:"true" default:"/verify"`
	Invite string `split_words:"true" default:"/invite"`
	Reset  string `split_words:"true" default:"/reset"`
}

func (c URLConfig) Validate() error {
	if c.Base == "" {
		return fmt.Errorf("invalid email url configuration: base URL is required")
	}

	if c.Invite == "" {
		return fmt.Errorf("invalid email url configuration: invite path is required")
	}

	if c.Verify == "" {
		return fmt.Errorf("invalid email url configuration: verify path is required")
	}

	if c.Reset == "" {
		return fmt.Errorf("invalid email url configuration: reset path is required")
	}

	return nil
}

// InviteURL constructs an invite URL from the token
func (c URLConfig) InviteURL(token string) (string, error) {
	if token == "" {
		return "", fmt.Errorf("token is required")
	}

	base, _ := url.Parse(c.Base)
	url := base.ResolveReference(&url.URL{Path: c.Invite, RawQuery: url.Values{"token": []string{token}}.Encode()})
	return url.String(), nil
}

// VerifyURL constructs a verify URL from the token
func (c URLConfig) VerifyURL(token string) (string, error) {
	if token == "" {
		return "", fmt.Errorf("token is required")
	}

	base, _ := url.Parse(c.Base)
	url := base.ResolveReference(&url.URL{Path: c.Verify, RawQuery: url.Values{"token": []string{token}}.Encode()})
	return url.String(), nil
}

// ResetURL constructs a reset URL from the token
func (c URLConfig) ResetURL(token string) (string, error) {
	if token == "" {
		return "", fmt.Errorf("token is required")
	}

	base, _ := url.Parse(c.Base)
	url := base.ResolveReference(&url.URL{Path: c.Reset, RawQuery: url.Values{"token": []string{token}}.Encode()})
	return url.String(), nil
}
