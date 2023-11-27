package store

import (
	"encoding/base64"
	"encoding/json"
)

// Token represents a token object
type Token struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	SessionState     string `json:"session_state"`
	Scope            string `json:"scope"`
	IDToken          string `json:"id_token"`
}

// Parse function takes a string and optional options, and returns a Token and an error
func Parse(v string, opts ...OptionsParse) (*Token, error) {
	var o optionsParse
	for _, opt := range opts {
		opt(&o)
	}

	cookieRecord := Token{}
	vByte := []byte(v)

	if o.b64 {
		var err error
		vByte, err = base64.StdEncoding.DecodeString(v)

		if err != nil {
			return nil, err
		}
	}

	// parse cookie
	if err := json.Unmarshal(vByte, &cookieRecord); err != nil {
		return nil, err
	}

	return &cookieRecord, nil
}

type optionsParse struct {
	b64 bool
}

// OptionsParse allows the `Parse` function to accept optional parameters by providing a way to customize its behavior
type OptionsParse func(*optionsParse)

func WithBase64(v bool) OptionsParse {
	return func(o *optionsParse) {
		o.b64 = v
	}
}
