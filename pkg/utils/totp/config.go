package totp

const (
	defaultLength = 6
)

type Config struct {
	// Enabled is a flag to enable or disable the OTP service
	Enabled bool `json:"enabled" koanf:"enabled" default:"true"`
	// CodeLength is the length of the OTP code
	CodeLength int `json:"codeLength" koanf:"codeLength" default:"6"`
	// Issuer is the issuer for TOTP codes
	Issuer string `json:"issuer" koanf:"issuer" default:"Datum"`
	// WithRedis configures the service with a redis client
	WithRedis bool `json:"redis" koanf:"redis" default:"true"`
}

// NewOTP returns a new OTP validator
func NewOTP(options ...ConfigOption) TOTPManager {
	s := OTP{
		codeLength: defaultLength,
	}

	for _, opt := range options {
		opt(&s)
	}

	return &s
}

// ConfigOption configures the validator
type ConfigOption func(*OTP)

// WithCodeLength configures the service with a length for random code generation
func WithCodeLength(length int) ConfigOption {
	return func(s *OTP) {
		s.codeLength = length
	}
}

// WithIssuer configures the service with a TOTP issuing domain
func WithIssuer(issuer string) ConfigOption {
	return func(s *OTP) {
		s.totpIssuer = issuer
	}
}

// WithSecret sets a new versioned Secret on the client
func WithSecret(x Secret) ConfigOption {
	return func(s *OTP) {
		s.secrets = append(s.secrets, x)
	}
}

// WithRedis configures the service with a redis client
func WithRedis(db otpRedis) ConfigOption {
	return func(s *OTP) {
		s.db = db
	}
}
