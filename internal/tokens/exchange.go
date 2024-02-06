package tokens

// ExchangeTokenResponse is the request response when exchanging an oauth token from one provider
// to a Datum issued token
type ExchangeTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}
