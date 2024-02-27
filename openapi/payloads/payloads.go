package payloads

// ResponseError is used as a payload for all errors. Use NewResponseError function
// to create new type to set some fields correctly.
type ResponseError struct {
	// HTTP status code
	HTTPStatusCode int `json:"-" yaml:"-"`

	// user facing error message
	Message string `json:"msg,omitempty" yaml:"msg,omitempty"`

	// full root cause
	Error string `json:"error" yaml:"error"`

	// build commit
	Version string `json:"version" yaml:"version"`

	// environment (prod or stage or ephemeral)
	Environment string `json:"environment,omitempty" yaml:"environment"`
}

type PubkeyRequest struct {
	Name string `json:"name" yaml:"name" description:"Enter the name of the newly created pubkey."`
	Body string `json:"body" yaml:"body" description:"Add a public part of a SSH key pair."`
}

type PubkeyResponse struct {
	ID                int64  `json:"id" yaml:"id"`
	AccountID         int64  `json:"-" yaml:"-"`
	Name              string `json:"name" yaml:"name"`
	Body              string `json:"body" yaml:"body"`
	Type              string `json:"type,omitempty" yaml:"type,omitempty"`
	Fingerprint       string `json:"fingerprint,omitempty" yaml:"fingerprint,omitempty"`
	FingerprintLegacy string `json:"fingerprint_legacy,omitempty" yaml:"fingerprint_legacy,omitempty"`
}
