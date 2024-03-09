package totp

import (
	"errors"
	"fmt"
)

// Error represents an error within OTP/TOTP domain
type Error interface {
	Error() string
	Message() string
	Code() ErrCode
}

// ErrCode is a machine readable code representing an error within the authenticator domain
type ErrCode string

// ErrInvalidCode represents an error related to an invalid TOTP/OTP code
type ErrInvalidCode string

func (e ErrInvalidCode) Code() ErrCode   { return "invalid_code" }
func (e ErrInvalidCode) Error() string   { return fmt.Sprintf("[%s] %s", e.Code(), string(e)) }
func (e ErrInvalidCode) Message() string { return string(e) }

var (
	// ErrCannotDecodeOTPHash is an error representing a failure to decode an OTP hash
	ErrCannotDecodeOTPHash = errors.New("cannot decode otp hash")

	// ErrInvalidOTPHashFormat is an error representing an invalid OTP hash format
	ErrInvalidOTPHashFormat = errors.New("invalid otp hash format")

	// ErrFailedToHashCode is an error representing a failure to hash code
	ErrFailedToHashCode = errors.New("failed to hash code")

	// ErrCipherTextTooShort is an error representing a ciphertext that is too short
	ErrCipherTextTooShort = errors.New("ciphertext too short")

	// ErrFailedToCreateCipherBlock is an error representing a failure to create a cipher block
	ErrFailedToCreateCipherBlock = errors.New("failed to create cipher block")

	// ErrCannotDecodeSecret is an error representing a failure to decode a secret
	ErrCannotDecodeSecret = errors.New("cannot decode secret")

	// ErrCannotWriteSecret is an error representing a failure to write a secret
	ErrCannotWriteSecret = errors.New("cannot write secret")

	// ErrFailedToDetermineSecretVersion is an error representing a failure to determine secret version
	ErrFailedToDetermineSecretVersion = errors.New("failed to determine secret version")

	// ErrFailedToCreateCipherText is an error representing a failure to create cipher text
	ErrFailedToCreateCipherText = errors.New("failed to create cipher text")

	// ErrNoSecretKeyForVersion is an error representing no secret key for version
	ErrNoSecretKeyForVersion = errors.New("no secret key for version")

	// ErrNoSecretKey is an error representing no secret key
	ErrNoSecretKey = errors.New("no secret key")

	// ErrFailedToValidateCode is an error representing a failure to validate code
	ErrFailedToValidateCode = errors.New("failed to validate code")

	// ErrCodeIsNoLongerValid is an error representing a code that is no longer valid
	ErrCodeIsNoLongerValid = errors.New("code is no longer valid")

	// ErrIncorrectCodeProvided is an error representing an incorrect code provided
	ErrIncorrectCodeProvided = errors.New("incorrect code provided")

	// ErrCannotDecryptSecret is an error representing a failure to decrypt secret
	ErrCannotDecryptSecret = errors.New("cannot decrypt secret")

	// ErrFailedToGetSecretForQR is an error representing a failure to get secret for qr
	ErrFailedToGetSecretForQR = errors.New("failed to get secret for qr")

	// ErrFailedToGenerateSecret is an error representing a failure to generate secret
	ErrFailedToGenerateSecret = errors.New("failed to generate secret")

	// ErrCannotHashOTPString is an error representing a failure to hash otp string
	ErrCannotHashOTPString = errors.New("cannot hash otp string")

	// ErrCannotGenerateRandomString is an error representing a failure to generate random string
	ErrCannotGenerateRandomString = errors.New("cannot generate random string")
)
