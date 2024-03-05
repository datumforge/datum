package totp

const (
	defaultLength = 6
)

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
