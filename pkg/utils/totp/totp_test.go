package totp

import (
	"database/sql"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestOTPManager(t *testing.T) {
	codeLength := 10
	svc := NewOTP(WithCodeLength(codeLength))
	code, hash, err := svc.OTPCode("mitb@datum.net", Email)

	if err != nil {
		t.Fatal("failed to create code:", err)
	}

	if len(code) != codeLength {
		t.Errorf("incorrect code length, want %v got %v", len(code), codeLength)
	}

	err = svc.ValidateOTP(code, hash)
	if err != nil {
		t.Error("failed to validate code:", err)
	}
}

func TestTOTPSecret(t *testing.T) {
	svc := NewOTP(
		WithIssuer("authenticator.local"),
		WithSecret(Secret{Version: 0, Key: "secret-key"}),
	)
	user := &User{
		IsTOTPAllowed:     true,
		IsEmailOTPAllowed: false,
		Phone: sql.NullString{
			String: "+17853931234",
			Valid:  true,
		},
	}

	secret, err := svc.TOTPSecret(user)
	if err != nil {
		t.Error("expected nil error, received:", err)
	}

	if secret == "" {
		t.Error("no secret generated")
	}
}

func TestTOTPQRString(t *testing.T) {
	svc := NewOTP(
		WithIssuer("authenticator.local"),
		WithSecret(Secret{
			Version: 1,
			Key:     "9f0c6da662f018b58b04a093e2dbb2e1d8d54250",
		}),
	)
	user := &User{
		IsTOTPAllowed:     true,
		IsEmailOTPAllowed: false,
		TFASecret:         "1:usrJIgtKY9j58GgLpKIaoJqNbwylphfzyJcoyRRg1Ow52/7j6KoRpky8tFLZlgrY",
		Phone: sql.NullString{
			String: "+17853931234",
			Valid:  true,
		},
	}
	qrString, err := svc.TOTPQRString(user)

	if err != nil {
		t.Error("expected nil error, received:", err)
	}

	expectedString := "otpauth://totp/authenticator.local:+17853931234?algorithm=" +
		"SHA1&digits=6&issuer=authenticator.local&period=30&secret=" +
		"572JFGKOMDRA6KHE5O3ZV62I6BP352E7"
	if !cmp.Equal(qrString, expectedString) {
		t.Error("TOTP QR string does not match",
			cmp.Diff(qrString, expectedString))
	}
}

func TestEncryptsWithLatestSecret(t *testing.T) {
	svc := &OTP{
		secrets: []Secret{
			{Version: 0, Key: "key-0"},
			{Version: 1, Key: "key-1"},
			{Version: 2, Key: "key-2"},
		},
	}
	secret := "some-secret-value"
	s, err := svc.encrypt(secret)

	if err != nil {
		t.Error("failed to encrypt", err)
	}

	if s == secret {
		t.Error("value not encrypted")
	}

	if !strings.HasPrefix(s, "2:") {
		t.Error("value not encrypted with latest secret")
	}

	s, err = svc.decrypt(s)

	if err != nil {
		t.Error("failed to decrypt secret", err)
	}

	if s != secret {
		t.Error("value not decrypted")
	}
}
