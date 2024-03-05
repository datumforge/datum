package totp

const (
	defaultLength             = 6
	defaultRecoveryCodeCount  = 16
	defaultRecoveryCodeLength = 10
	codePeriod                = 30
)

type Config struct {
	// Enabled is a flag to enable or disable the OTP service
	Enabled bool `json:"enabled" koanf:"enabled" default:"true"`
	// CodeLength is the length of the OTP code
	CodeLength int `json:"codeLength" koanf:"codeLength" default:"6"`
	// Issuer is the issuer for TOTP codes
	Issuer string `json:"issuer" koanf:"issuer" default:"datum"`
	// WithRedis configures the service with a redis client
	WithRedis bool `json:"redis" koanf:"redis" default:"true"`
	// Secret stores a versioned secret key for cryptography functions
	Secret string `json:"secret" koanf:"secret"`
	// RecoveryCodeCount is the number of recovery codes to generate
	RecoveryCodeCount int `json:"recoveryCodeCount" koanf:"recoveryCodeCount" default:"16"`
	// RecoveryCodeLength is the length of a recovery code
	RecoveryCodeLength int `json:"recoveryCodeLength" koanf:"recoveryCodeLength" default:"8"`
}

// NewOTP returns a new OTP validator
func NewOTP(options ...ConfigOption) TOTPManager {
	s := OTP{
		codeLength:         defaultLength,
		ttl:                codePeriod,
		recoveryCodeCount:  defaultRecoveryCodeCount,
		recoveryCodeLength: defaultRecoveryCodeLength,
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
		s.issuer = issuer
	}
}

// WithRecoveryCodeCount configures the service with a number of recovery codes to generate
func WithRecoveryCodeCount(count int) ConfigOption {
	return func(s *OTP) {
		s.recoveryCodeCount = count
	}
}

// WithRecoveryCodeLength configures the service with the length of recovery codes to generate
func WithRecoveryCodeLength(length int) ConfigOption {
	return func(s *OTP) {
		s.recoveryCodeLength = length
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
