package entx

// CascadeAnnotationName is a name for our cascading delete annotation
var CascadeAnnotationName = "DATUM_CASCADE"
var SchemaGenAnnotationName = "DATUM_SCHEMAGEN"

// CascadeAnnotation is an annotation used to indicate that an edge should be cascaded
type CascadeAnnotation struct {
	Field string
}

type SchemaGenAnnotation struct {
	Skip bool
}

// Name returns the name of the CascadeAnnotation
func (a CascadeAnnotation) Name() string {
	return CascadeAnnotationName
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

// SchemaGenSkip sets the whether schema generation should be skipped for this type
func SchemaGenSkip(skip bool) *SchemaGenAnnotation {
	return &SchemaGenAnnotation{
		Skip: skip,
	}
}
