package rows

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

// Writer is defining an interface called `Writer`
type Writer interface {
	Write([]string) error
}

// NewTabRowWriter is a function that creates a new instance of the `TabRowWriter` struct,
// which implements the `Writer` interface. It takes a pointer to a `tabwriter.Writer` as a parameter
// and returns a `Writer` interface.
func NewTabRowWriter(w *tabwriter.Writer) Writer {
	if w == nil {
		w = tabwriter.NewWriter(os.Stdout, 1, 0, 4, ' ', 0) //nolint:gomnd
	}

	return &TabRowWriter{*w}
}

// TabRowWriter struct is defining a new struct type called `TabRowWriter`. This struct is
// used to implement the `Writer` interface. It has a single field, `tabwriter.Writer`, which is
// embedded within the struct. This allows the `TabRowWriter` struct to have access to all the methods
// and fields of the embedded `tabwriter.Writer`.
type TabRowWriter struct {
	tabwriter.Writer
}

// Write method is implementing the `Write` method of the `Writer`
// interface for the `TabRowWriter` struct. It takes a slice of strings called `record` as a parameter
// and returns an error.
func (w *TabRowWriter) Write(record []string) error {
	fmt.Fprintln(&w.Writer, strings.Join(record, "\t"))
	return nil
}
