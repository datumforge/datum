package totp

import (
	"context"
	"database/sql"
	"time"
)

// TokenState represents a state of a JWT token.
// A token may represent an intermediary state prior
// to authorization (ex. TOTP code is required)
type TokenState string

// DeliveryMethod represents a mechanism to send messages to users
type DeliveryMethod string

// TFAOptions represents options a user may use to complete 2FA
type TFAOptions string

// MessageType describes a classification of a Message
type MessageType string

// Manager manages the protocol for SMS/Email 2FA codes and TOTP codes
type Manager struct {
	TOTPManager TOTPManager
}

const (
	// OTPEmail allows a user to complete TFA with an OTP code delivered via email
	OTPEmail TFAOptions = "otp_email"
	// OTPPhone allows a user to complete TFA with an OTP code delivered via phone
	OTPPhone TFAOptions = "otp_phone"
	// TOTP allows a user to complete TFA with a TOTP device or application
	TOTP TFAOptions = "totp"
	// Phone is a delivery method for text messages
	Phone DeliveryMethod = "phone"
	// Email is a delivery method for email
	Email = "email"
	// OTPAddress is a message containing an OTP code for contact verification
	OTPAddress MessageType = "otp_address"
	// OTPResend is a message containing an OTP code
	OTPResend MessageType = "otp_resend"
	// OTPLogin is a message containing an OTP code for login
	OTPLogin MessageType = "otp_login"
	// OTPSignup is a message containing an OTP code for signup
	OTPSignup MessageType = "otp_signup"
)

// User represents a user who is registered with the service
type User struct {
	// ID is a unique ID for the user
	ID string
	// Phone number associated with the account
	Phone sql.NullString
	// Email address associated with the account
	Email sql.NullString
	// TFASecret is a a secret string used to generate 2FA TOTP codes
	TFASecret string
	// IsPhoneAllowed specifies a user may complete authentication by verifying an OTP code delivered through SMS
	IsPhoneOTPAllowed bool
	// IsEmailOTPAllowed specifies a user may complete authentication by verifying an OTP code delivered through email
	IsEmailOTPAllowed bool
	// IsTOTPAllowed specifies a user may complete authentication by verifying a TOTP code
	IsTOTPAllowed bool
}

// DefaultOTPDelivery returns the default OTP delivery method
func (u *User) DefaultOTPDelivery() DeliveryMethod {
	if u.Email.String != "" {
		return Email
	}

	return Phone
}

// DefaultName returns the default name for a user (email or phone)
func (u *User) DefaultName() string {
	if u.Email.String != "" {
		return u.Email.String
	}

	return u.Phone.String
}

// Token is a token that provides proof of User authentication
type Token struct {
	// Phone is a User's phone number
	Phone string `json:"phone_number"`
	// CodeHash is the hash of a randomly generated code used
	// to validate an OTP code and escalate the token to an
	// authorized token
	CodeHash string `json:"code,omitempty"`
	// Code is the unhashed value of CodeHash. This value is
	// not persisted and returned to the client outside of the JWT
	// response through an alternative mechanism (e.g. Email). It is
	// validated by ensuring the SHA512 hash of the value matches the
	// CodeHash embedded in the token
	Code string `json:"-"`
	// TFAOptions represents available options a user may use to complete
	// 2FA.
	TFAOptions []TFAOptions `json:"tfa_options"`
}

// Message is a message to be delivered to a user
type Message struct {
	// Type describes the classification of a Message
	Type MessageType
	// Subject is a human readable subject describe the Message
	Subject string
	// Delivery type of the message (e.g. phone or email)
	Delivery DeliveryMethod
	// Vars contains key/value variables to populate
	// templated content
	Vars map[string]string
	// Content of the message
	Content string
	// Delivery address of the user (e.g. phone or email)
	Address string
	// ExpiresAt is the latest time we can attempt delivery
	ExpiresAt time.Time
	// DeliveryAttempts is the total amount of delivery attempts made
	DeliveryAttempts int
}

// TOTPManager manages the protocol for SMS/Email 2FA codes and TOTP codes
type TOTPManager interface {
	// TOTPQRString returns a URL string used for TOTP code generation
	TOTPQRString(u *User) (string, error)
	// TOTPSecret creates a TOTP secret for code generation
	TOTPSecret(u *User) (string, error)
	// OTPCode creates a random OTP code and hash
	OTPCode(address string, method DeliveryMethod) (code, hash string, err error)
	// ValidateOTP checks if a User email/sms delivered OTP code is valid
	ValidateOTP(code, hash string) error
	// ValidateTOTP checks if a User TOTP code is valid
	ValidateTOTP(ctx context.Context, user *User, code string) error
	// GenerateRecoveryCodes creates a set of recovery codes for a user
	GenerateRecoveryCodes() []string
}
