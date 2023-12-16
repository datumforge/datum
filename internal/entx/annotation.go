package entx

// CascadeAnnotationName is a name for our cascading delete annotation
var CascadeAnnotationName = "DATUM_CASCADE"

// CascadeAnnotation is an annotation used to indicate that an edge should be cascaded
type CascadeAnnotation struct {
	Field string
}

// Name returns the name of the CascadeAnnotation
func (a CascadeAnnotation) Name() string {
	return CascadeAnnotationName
}

// CascadeAnnotationField sets the field name of the edge containing the ID of a record from the current schema
func CascadeAnnotationField(fieldname string) *CascadeAnnotation {
	return &CascadeAnnotation{
		Field: fieldname,
	}
}
