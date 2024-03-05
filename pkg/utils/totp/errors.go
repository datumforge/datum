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
	// ErrCannotDecodeOTPHash
	ErrCannotDecodeOTPHash = errors.New("cannot decode otp hash")

	// ErrInvalidOTPHashFormat
	ErrInvalidOTPHashFormat = errors.New("invalid otp hash format")

	// ErrFailedToHashCode
	ErrFailedToHashCode = errors.New("failed to hash code")

	// ErrCiphertextTooShort
	ErrCiphertextTooShort = errors.New("ciphertext too short")

	// ErrFailedToCreateCipherBlock
	ErrFailedToCreateCipherBlock = errors.New("failed to create cipher block")

	// ErrCannotDecodeSecret
	ErrCannotDecodeSecret = errors.New("cannot decode secret")

	// ErrCannotWriteSecret
	ErrCannotWriteSecret = errors.New("cannot write secret")

	// ErrFailedToDetermineSecretVersion
	ErrFailedToDetermineSecretVersion = errors.New("failed to determine secret version")

	// ErrFailedToCreateCipherText
	ErrFailedToCreateCipherText = errors.New("failed to create cipher text")

	// ErrNoSecretKeyForVersion
	ErrNoSecretKeyForVersion = errors.New("no secret key for version")

	// ErrNoSecretKey
	ErrNoSecretKey = errors.New("no secret key")

	// ErrFailedToValidateCode
	ErrFailedToValidateCode = errors.New("failed to validate code")

	// ErrCodeIsNoLongerValid
	ErrCodeIsNoLongerValid = errors.New("code is no longer valid")

	// ErrIncorrectCodeProvided
	ErrIncorrectCodeProvided = errors.New("incorrect code provided")

	// ErrCannotDecryptSecret
	ErrCannotDecryptSecret = errors.New("cannot decrypt secret")

	// ErrFailedToGetSecretForQR
	ErrFailedToGetSecretForQR = errors.New("failed to get secret for qr")

	// ErrFailedtoGenerateSecret
	ErrFailedtoGenerateSecret = errors.New("failed to generate secret")

	// ErrCannotHashOTPString
	ErrCannotHashOTPString = errors.New("cannot hash otp string")

	// ErrCannotGenerateRandomString
	ErrCannotGenerateRandomString = errors.New("cannot generate random string")
)
