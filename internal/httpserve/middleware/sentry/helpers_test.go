package sentry

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLimitTagValue(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "Short string",
			str:  "i3WEG3605Kj7i3WEG3605Kj7i3WEG3605Kj7i3WEG3605Kj7i3WEG3605Kj7i3WEG3605Kj7i3WEG3605Kj7i3WEG3605Kj7i3WEG3605Kj7",
			want: "i3WEG3605Kj7i3WEG3605Kj7i3WEG3605Kj7i3WEG3605Kj7i3WEG3605Kj7i3WEG3605Kj7i3WEG3605Kj7i3WEG3605Kj7i3WEG3605Kj7",
		},
		{
			name: "Long string containing \\n",
			str:  "05Kj7z2AXCl603gMJu6B23z2sD05\nKj7z2AXCl603gMJu6B23z2sD05Kj7z\n2AXCl603gMJu6B23z2sD05Kj7z2AXCl603gMJu6B23z2sD05Kj7z2AXCl603gMJu6B23z2sD05Kj7z2AXCl603gMJu6B23z2sD05Kj7z2AXCl603gMJu6B23z2sD05Kj7z2AXCl603gMJu6B23z2sD05Kj7z2AXCl603gMJu6B23z2sD05Kj7z2AXCl603gMJu6B23z2sD05Kj7z2AXCl603gMJu6B23z2sD",
			want: "05Kj7z2AXCl603gMJu6B23z2sD05 Kj7z2AXCl603gMJu6B23z2sD05Kj7z 2AXCl603gMJu6B23z2sD05Kj7z2AXCl603gMJu6B23z2sD05Kj7z2AXCl603gMJu6B23z2sD05Kj7z2AXCl603gMJu6B23z2sD05Kj7z2AXCl603gMJu6B23z2sD05Kj7z2AXCl60...",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.True(t, len(prepareTagValue(tt.str)) <= 200)
			require.Equal(t, tt.want, prepareTagValue(tt.str))
		})
	}
}

func TestLimitTagName(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "Short string",
			str:  "i3WEG3605Kj7",
			want: "i3WEG3605Kj7",
		},
		{
			name: "Long string",
			str:  "05Kj7z2AXCl603gMJu6B23z2sD05Kj7z2AX",
			want: "05Kj7z2AXCl603gMJu6B23z2sD05Kj7z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.True(t, len(prepareTagName(tt.str)) <= 32)
			require.Equal(t, tt.want, prepareTagName(tt.str))
		})
	}
}
