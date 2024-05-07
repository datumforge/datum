package emitter

// Listener is a function type that can handle events of any type
type Listener func(Event) error

// listenerItem stores a listener along with its unique identifier and priority
type listenerItem struct {
	listener Listener
	priority Priority
}

// ListenerOption is a function type that configures listener behavior
type ListenerOption func(*listenerItem)

// WithPriority sets the priority of a listener
func WithPriority(priority Priority) ListenerOption {
	return func(item *listenerItem) {
		item.priority = priority
	}
}
