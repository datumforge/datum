package tokens

// TokenResponse is the request response when exchanging an oauth token from one provider
// to a Datum issued token
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
}
