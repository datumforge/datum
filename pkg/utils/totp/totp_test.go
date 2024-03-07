package totp

import (
	"database/sql"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOTPManager(t *testing.T) {
	codeLength := 10
	svc := NewOTP(
		WithCodeLength(codeLength),
	)

	code, hash, err := svc.OTPCode("mitb@datum.net", Email)
	require.NoErrorf(t, err, "failed to create code: %v", err)

	assert.Len(t, code, codeLength, "incorrect code length")

	err = svc.ValidateOTP(code, hash)
	require.NoErrorf(t, err, "failed to validate code: %v", err)
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
	require.NoError(t, err)
	assert.NotNil(t, secret, "no secret generated")
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
	require.NoError(t, err)

	expectedString := "otpauth://totp/authenticator.local:+17853931234?algorithm=" +
		"SHA1&digits=6&issuer=authenticator.local&period=30&secret=" +
		"572JFGKOMDRA6KHE5O3ZV62I6BP352E7"
	assert.Equal(t, expectedString, qrString, "TOTP QR string does not match")
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
	require.NoError(t, err, "failed to encrypt secret")

	assert.NotEqual(t, secret, s, "value not encrypted")
	assert.True(t, strings.HasPrefix(s, "2:"), "value not encrypted with latest secret")

	s, err = svc.decrypt(s)
	require.NoError(t, err, "failed to decrypt secret")

	assert.Equal(t, secret, s, "value not decrypted")
}
func TestGenerateRecoveryCodes(t *testing.T) {
	o := &OTP{
		recoveryCodeCount:  5,
		recoveryCodeLength: 8,
	}

	codes := o.GenerateRecoveryCodes()

	assert.Len(t, codes, 5, "incorrect number of recovery codes generated")

	for _, code := range codes {
		assert.Len(t, code, 8, "incorrect recovery code length")
	}
}
