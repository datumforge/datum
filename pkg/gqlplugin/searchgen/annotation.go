package searchgen

import "encoding/json"

// SearchFieldAnnotationName is a name for the search field annotation
var SearchFieldAnnotationName = "DATUM_SEARCH"

// SearchFieldAnnotation is an annotation used to indicate that the field should be searchable
type SearchFieldAnnotation struct {
	// Searchable indicates that the field should be searchable
	Searchable bool
}

// Name returns the name of the SearchFieldAnnotation
func (a SearchFieldAnnotation) Name() string {
	return SearchFieldAnnotationName
}

// FieldSearchable returns a new SearchFieldAnnotation with the searchable flag set
func FieldSearchable() *SearchFieldAnnotation {
	return &SearchFieldAnnotation{
		Searchable: true,
	}
}

// Decode unmarshalls the SearchFieldAnnotation
func (a *SearchFieldAnnotation) Decode(annotation interface{}) error {
	buf, err := json.Marshal(annotation)
	if err != nil {
		return err
	}

	return json.Unmarshal(buf, a)
}
