package entx_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/datumforge/datum/internal/entx"

	"github.com/stretchr/testify/assert"
)

func TestCascadeAnnotation(t *testing.T) {
	f := gofakeit.Name()
	ca := entx.CascadeAnnotationField(f)

	assert.Equal(t, ca.Name(), entx.CascadeAnnotationName)
	assert.Equal(t, ca.Field, f)
}

func TestSchemaGenAnnotation(t *testing.T) {
	s := gofakeit.Bool()
	sa := entx.SchemaGenSkip(s)

	assert.Equal(t, sa.Name(), entx.SchemaGenAnnotationName)
	assert.Equal(t, sa.Skip, s)
}
