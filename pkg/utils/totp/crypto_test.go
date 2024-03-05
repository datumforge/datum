package totp

import (
	"encoding/base64"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestString(t *testing.T) {
	samples := "abcdefghijklmnopqrstuv"

	for i := 0; i <= 10; i++ {
		ln := 50
		random, err := String(ln, samples)

		if err != nil {
			t.Fatal("failed to generate random string", err)
		}

		if len(random) != 50 {
			t.Error("incorrect character count", cmp.Diff(
				len(random), 50,
			))
		}

		for _, v := range random {
			s := string(v)
			if !strings.Contains(samples, s) {
				t.Errorf("invalid character used in random string: %s", s)
			}
		}
	}
}

func TestStringB64(t *testing.T) {
	b64str, err := StringB64(50)
	if err != nil {
		t.Error("error generating random string", err)
	}

	_, err = base64.StdEncoding.DecodeString(b64str)
	if err != nil {
		t.Error("failed to decode base64 encoded string")
	}
}

func TestOTPHash(t *testing.T) {
	str := "the quick brown fox"
	hash, err := OTPHash(str)

	if err != nil {
		t.Error("error generating hash", err)
	}

	if str == hash {
		t.Error("string not hashed")
	}

	hash2, err := OTPHash(str)
	if err != nil {
		t.Error("error generating hash", err)
	}

	if hash != hash2 {
		t.Error("hashes do not match", cmp.Diff(hash, hash2))
	}
}
