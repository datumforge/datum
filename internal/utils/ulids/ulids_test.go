package ulids_test

import (
	"sync"
	"testing"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/require"

	ulidlib "github.com/datumforge/datum/internal/utils/ulids"
)

func TestIsZero(t *testing.T) {
	testCases := []struct {
		input  ulid.ULID
		assert require.BoolAssertionFunc
	}{
		{ulid.ULID{}, require.True},
		{ulid.ULID{0x00}, require.True},
		{ulid.ULID{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, require.True},
		{ulid.Make(), require.False},
	}

	for _, tc := range testCases {
		tc.assert(t, ulidlib.IsZero(tc.input))
	}
}

func TestNew(t *testing.T) {
	// Should be able to concurrently create 100 ulids.
	var wg sync.WaitGroup

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()

			uid := ulidlib.New()

			require.False(t, ulidlib.IsZero(uid))
		}()
	}
	wg.Wait()
}

func TestFromTime(t *testing.T) {
	// Should be able to concurrently create 100 ulids from a timestamp.
	now := time.Now()
	vals := sync.Map{}

	var wg sync.WaitGroup

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()

			uid := ulidlib.FromTime(now)

			require.False(t, ulidlib.IsZero(uid))
			vals.Store(uid, struct{}{})
		}()
	}
	wg.Wait()

	nunique := 0

	vals.Range(func(key, value any) bool {
		nunique++
		return true
	})
	require.Equal(t, 100, nunique)
}

func TestParse(t *testing.T) {
	example := ulidlib.New()

	testCases := []struct {
		input    any
		expected ulid.ULID
		err      error
	}{
		{example.String(), example, nil},
		{example.Bytes(), example, nil},
		{example, example, nil},
		{[16]byte(example), example, nil},
		{"", ulidlib.Null, nil},
		{uint64(14), ulidlib.Null, ulidlib.ErrUnknownType},
		{"foo", ulidlib.Null, ulid.ErrDataSize},
		{[]byte{0x14, 0x21}, ulidlib.Null, ulid.ErrDataSize},
		{ulidlib.Null.String(), ulidlib.Null, nil},
	}

	for i, tc := range testCases {
		actual, err := ulidlib.Parse(tc.input)
		require.ErrorIs(t, err, tc.err, "could not compare error on test case %d", i)
		require.Equal(t, tc.expected, actual, "expected result not returned")

		if tc.err != nil {
			require.Panics(t, func() { ulidlib.MustParse(tc.input) })
		} else {
			require.Equal(t, tc.expected, ulidlib.MustParse(tc.input))
		}
	}
}

func TestBytes(t *testing.T) {
	example := ulidlib.New()
	expected := example.Bytes()

	testCases := []struct {
		input    any
		expected []byte
		err      error
	}{
		{example.String(), expected, nil},
		{example.Bytes(), expected, nil},
		{example, expected, nil},
		{[16]byte(example), expected, nil},
		{"", []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, nil},
		{uint64(14), nil, ulidlib.ErrUnknownType},
		{"foo", nil, ulid.ErrDataSize},
		{[]byte{0x14, 0x21}, nil, ulid.ErrDataSize},
		{ulidlib.Null.String(), []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, nil},
	}

	for i, tc := range testCases {
		actual, err := ulidlib.Bytes(tc.input)
		require.ErrorIs(t, err, tc.err, "could not compare error on test case %d", i)
		require.Equal(t, tc.expected, actual, "expected result not returned on test case %d", i)

		if tc.err != nil {
			require.Panics(t, func() { ulidlib.MustBytes(tc.input) })
		} else {
			require.Equal(t, tc.expected, ulidlib.MustBytes(tc.input))
		}
	}
}
