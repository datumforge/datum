package entx

// CascadeAnnotationName is a name for our cascading delete annotation
var CascadeAnnotationName = "DATUM_CASCADE"

// CascadeThroughAnnotationName is a name for our cascading through edge delete annotation
var CascadeThroughAnnotationName = "DATUM_CASCADE_THROUGH"

// SchemaGenAnnotationName is a name for our graphql schema generation annotation
var SchemaGenAnnotationName = "DATUM_SCHEMAGEN"

// CascadeAnnotation is an annotation used to indicate that an edge should be cascaded
type CascadeAnnotation struct {
	Field string
}

// CascadeThroughAnnotation is an annotation used to indicate that an edge should be cascaded through
type CascadeThroughAnnotation struct {
	Schemas []ThroughCleanup
}

type ThroughCleanup struct {
	Field   string
	Through string
}

type SchemaGenAnnotation struct {
	Skip bool
}

// Name returns the name of the CascadeAnnotation
func (a CascadeAnnotation) Name() string {
	return CascadeAnnotationName
}

// Name returns the name of the CascadeThroughAnnotation
func (a CascadeThroughAnnotation) Name() string {
	return CascadeThroughAnnotationName
}

func (a SchemaGenAnnotation) Name() string {
	return SchemaGenAnnotationName
}

// CascadeAnnotationField sets the field name of the edge containing the ID of a record from the current schema
func CascadeAnnotationField(fieldname string) *CascadeAnnotation {
	return &CascadeAnnotation{
		Field: fieldname,
	}
}

// CascadeThroughAnnotationField sets the field name of the edge containing the ID of a record from the current schema
func CascadeThroughAnnotationField(schemas []ThroughCleanup) *CascadeThroughAnnotation {
	return &CascadeThroughAnnotation{
		Schemas: schemas,
	}
}

// SchemaGenSkip sets the whether schema generation should be skipped for this type
func SchemaGenSkip(skip bool) *SchemaGenAnnotation {
	return &SchemaGenAnnotation{
		Skip: skip,
	}
}
