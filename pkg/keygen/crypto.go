package keygen

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"
)

// S256Challenge creates [RFC 7636]: https://datatracker.ietf.org/doc/html/rfc7636#section-4.2
func S256Challenge(code string) string {
	h := sha256.New()
	h.Write([]byte(code))

	return strings.TrimRight(base64.URLEncoding.EncodeToString(h.Sum(nil)), "=")
}

// SHA256 creates sha256 hash as defined in FIPS 180-4
func SHA256(text string) string {
	h := sha256.New()
	h.Write([]byte(text))

	return fmt.Sprintf("%x", h.Sum(nil))
}

// SHA512 creates sha512 hash as defined in FIPS 180-4
func SHA512(text string) string {
	h := sha512.New()
	h.Write([]byte(text))

	return fmt.Sprintf("%x", h.Sum(nil))
}

// HS256 creates a HMAC hash with sha256 digest algorithm.
func HS256(text string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(text))

	return fmt.Sprintf("%x", h.Sum(nil))
}

// HS512 creates a HMAC hash with sha512 digest algorithm
func HS512(text string, secret string) string {
	h := hmac.New(sha512.New, []byte(secret))
	h.Write([]byte(text))

	return fmt.Sprintf("%x", h.Sum(nil))
}

// Equal compares two hash strings for equality without leaking timing information
func Equal(hash1 string, hash2 string) bool {
	return subtle.ConstantTimeCompare([]byte(hash1), []byte(hash2)) == 1
}
