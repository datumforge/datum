package entfga

// Annotations of the fga extension
type Annotations struct {
	ObjectType string `json:"object_type,omitempty"` // Object type for the fga relationship
}

// Name of the annotation
func (Annotations) Name() string {
	return "Authz"
}
