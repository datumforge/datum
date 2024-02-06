package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"

	ent "github.com/datumforge/datum/internal/ent/generated"
)

func TestParseName(t *testing.T) {
	tests := []struct {
		name string
		user string
		want ent.CreateUserInput
	}{
		{
			name: "happy path",
			user: "Matty Anderson",
			want: ent.CreateUserInput{
				FirstName: "Matty",
				LastName:  "Anderson",
			},
		},
		{
			name: "very long name",
			user: "Matty Anderson Is The Best",
			want: ent.CreateUserInput{
				FirstName: "Matty",
				LastName:  "Anderson Is The Best",
			},
		},
		{
			name: "single name",
			user: "Matty",
			want: ent.CreateUserInput{
				FirstName: "Matty",
				LastName:  "",
			},
		},
		{
			name: "empty name",
			user: "",
			want: ent.CreateUserInput{
				FirstName: "",
				LastName:  "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseName(tt.user)

			assert.Equal(t, tt.want, got)
		})
	}
}
