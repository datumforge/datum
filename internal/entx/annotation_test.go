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

func TestCascadeThroughAnnotation(t *testing.T) {
	f := gofakeit.Name()
	s := gofakeit.Name()
	schemas := []entx.ThroughCleanup{
		{
			Through: s,
			Field:   f,
		},
	}
	ca := entx.CascadeThroughAnnotationField(schemas)

	assert.Equal(t, ca.Name(), entx.CascadeThroughAnnotationName)
	assert.Equal(t, ca.Schemas[0].Field, f)
	assert.Equal(t, ca.Schemas[0].Through, s)
}

func TestSchemaGenAnnotation(t *testing.T) {
	s := gofakeit.Bool()
	sa := entx.SchemaGenSkip(s)

	assert.Equal(t, sa.Name(), entx.SchemaGenAnnotationName)
	assert.Equal(t, sa.Skip, s)
}
