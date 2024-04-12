package invoice

import (
	"errors"
)

var (
	// ErrCouldNotDoFontStuff is returned when we could not do font stuff
	ErrCouldNotDoFontStuff = errors.New("could not do font stuff")

	// ErrCouldNotSaveInvoice is returned when we could not save the invoice
	ErrCouldNotSaveInvoice = errors.New("could not save the invoice")

	// ErrCouldNotSetInvoiceLayout is returned when we could not set the invoice layout
	ErrCouldNotSetInvoiceLayout = errors.New("could not set the invoice layout")

	// ErrFailedParsingYAML is returned when the invoice input yaml could not be parsed
	ErrFailedParsingYAML = errors.New("could not unmarshal yaml values")
)
