package passwd_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	. "github.com/datumforge/datum/internal/passwd"
)

func TestDerivedKey(t *testing.T) {
	// Cannot create derived key for empty password
	_, err := CreateDerivedKey("")
	require.EqualError(t, err, "cannot create derived key for empty password")

	// Create a derived key from a password
	passwd, err := CreateDerivedKey("theeaglefliesatmidnight")
	require.NoError(t, err)

	verified, err := VerifyDerivedKey(passwd, "theeaglefliesatmidnight")
	require.NoError(t, err)
	require.True(t, verified)

	verified, err = VerifyDerivedKey(passwd, "thesearentthedroidsyourelookingfor")
	require.NoError(t, err)
	require.False(t, verified)

	// Create a derived key from a password
	passwd2, err := CreateDerivedKey("lightning")
	require.NoError(t, err)
	require.NotEqual(t, passwd, passwd2)
}

func TestDerivedKeyDetail(t *testing.T) {
	// Cannot verify empty derived key or password
	errmsg := "cannot verify empty derived key or password"
	_, err := VerifyDerivedKey("", "foo")
	require.EqualError(t, err, errmsg)
	_, err = VerifyDerivedKey("foo", "")
	require.EqualError(t, err, errmsg)

	// Parse failures
	errmsg = "cannot parse encoded derived key, does not match regular expression"
	_, err = VerifyDerivedKey("notarealkey", "supersecretpassword")
	require.EqualError(t, err, errmsg)

	dk := "$pbkdf2$v=19$m=65536,t=1,p=2$FrAEw4rWRDpyIZXR/QSzpg==$chQikgApfQfSaPZ7idk6caqBk79xRalpPUs4Ro/hywM="
	errmsg = "current code only works with the the dk protcol \"argon2id\""
	_, err = VerifyDerivedKey(dk, "supersecretpassword")
	require.EqualError(t, err, errmsg)

	dk = "$argon2id$v=13212$m=65536,t=1,p=2$FrAEw4rWRDpyIZXR/QSzpg==$chQikgApfQfSaPZ7idk6caqBk79xRalpPUs4Ro/hywM="
	errmsg = "expected argon2id version 19"
	_, err = VerifyDerivedKey(dk, "supersecretpassword")
	require.EqualError(t, err, errmsg)

	dk = "$argon2id$v=19$m=65536,t=999999999999999999,p=2$FrAEw4rWRDpyIZXR/QSzpg==$chQikgApfQfSaPZ7idk6caqBk79xRalpPUs4Ro/hywM="
	errmsg = "could not parse time \"999999999999999999\": strconv.ParseUint: parsing \"999999999999999999\": value out of range"
	_, err = VerifyDerivedKey(dk, "supersecretpassword")
	require.EqualError(t, err, errmsg)

	dk = "$argon2id$v=19$m=999999999999999999,t=1,p=2$FrAEw4rWRDpyIZXR/QSzpg==$chQikgApfQfSaPZ7idk6caqBk79xRalpPUs4Ro/hywM="
	errmsg = "could not parse memory \"999999999999999999\": strconv.ParseUint: parsing \"999999999999999999\": value out of range"
	_, err = VerifyDerivedKey(dk, "supersecretpassword")
	require.EqualError(t, err, errmsg)

	dk = "$argon2id$v=19$m=65536,t=1,p=999999999999999999$FrAEw4rWRDpyIZXR/QSzpg==$chQikgApfQfSaPZ7idk6caqBk79xRalpPUs4Ro/hywM="
	errmsg = "could not parse threads \"999999999999999999\": strconv.ParseUint: parsing \"999999999999999999\": value out of range"
	_, err = VerifyDerivedKey(dk, "supersecretpassword")
	require.EqualError(t, err, errmsg)

	dk = "$argon2id$v=19$m=65536,t=1,p=2$==FrAEw4rWRDpyIZXR/QSzpg==$chQikgApfQfSaPZ7idk6caqBk79xRalpPUs4Ro/hywM="
	errmsg = "could not parse salt: illegal base64 data at input byte 0"
	_, err = VerifyDerivedKey(dk, "supersecretpassword")
	require.EqualError(t, err, errmsg)

	dk = "$argon2id$v=19$m=65536,t=1,p=2$FrAEw4rWRDpyIZXR/QSzpg==$==chQikgApfQfSaPZ7idk6caqBk79xRalpPUs4Ro/hywM="
	errmsg = "could not parse derived key: illegal base64 data at input byte 0"
	_, err = VerifyDerivedKey(dk, "supersecretpassword")
	require.EqualError(t, err, errmsg)
}

func TestIsDerivedKey(t *testing.T) {
	testCases := []struct {
		input  string
		assert require.BoolAssertionFunc
	}{
		{"", require.False},
		{"notarealkey", require.False},
		{"$pbkdf2$v=19$m=65536,t=1,p=2$FrAEw4rWRDpyIZXR/QSzpg==$chQikgApfQfSaPZ7idk6caqBk79xRalpPUs4Ro/hywM=", require.True},
		{"$argon2id$v=19$m=65536,t=1,p=2$z9oBHDE02w+jIaTCyC0cgQ==$i0eOYj2V59sG4iqRN0douuwZ6+7QG32KbSP6XmV8lx8=", require.True},
	}

	for i, tc := range testCases {
		tc.assert(t, IsDerivedKey(tc.input), "test case %d failed", i)
	}
}
