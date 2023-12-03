package auth

import "errors"

var (
	ErrEmptyClientOauthState = errors.New("auth client's expected oauth state is empty")
	ErrMismatchedOauthState  = errors.New("mismatched oauth2 state")
	ErrBase64NumBytes        = errors.New("number of bytes is not divisible by three")
	ErrFailedToGetOauthToken = errors.New("failed to get oauth2 token")
)
