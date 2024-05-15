// nolint: gomnd
package invoice

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"gopkg.in/yaml.v3"
)

// New returns Invoice struct loaded with values from YAML and prepares PDF struct
func NewInvoice(file []byte) (*Invoice, error) {
	invoice := &Invoice{}
	if err := yaml.Unmarshal(file, &invoice); err != nil {
		return nil, ErrFailedParsingYAML
	}

	invoice.pdf = core.NewMaroto(orientation.Vertical, pagesize.A4)

	err := invoice.setPDFLayout()
	if err != nil {
		return nil, ErrCouldNotSetInvoiceLayout
	}

	return invoice, nil
}

func (i *Invoice) setPDFLayout() error {
	i.pdf.SetFirstPageNb(1)
	i.pdf.SetPageMargins(10, 15, 10)
	err := i.setFonts()

	if err != nil {
		return ErrCouldNotDoFontStuff
	}
	// tie everything together
	i.BuildInvoiceHeader()
	i.BuildPDFFooter()
	i.BuildCompanyDetails()
	i.BuildBankDetails()
	i.BuildInvoiceLineItems()
	i.BuildSignatureLines()

	_, height := i.pdf.GetPageSize()
	current := i.pdf.GetCurrentOffset()
	filler := height - current - 60
	i.pdf.Row(filler, func() { // intentionally blank
	})

	return nil
}

// SaveToPdf saves Invoice to a PDF file and closes the file
func (i *Invoice) SaveToPdf(outputPath string) error {
	if err := i.pdf.OutputFileAndClose(outputPath); err != nil {
		return ErrCouldNotSaveInvoice
	}

	return nil
}

// Save saves Invoice to bytes and closes it
func (i *Invoice) SaveAsBytes() ([]byte, error) {
	bytes, err := i.pdf.Output()

	if err != nil {
		return nil, ErrCouldNotSaveInvoice
	}

	return bytes.Bytes(), err
}

// Invoice parameters.
type Invoice struct {
	pdf       core.Maroto
	Number    string  `yaml:"number"`
	IssueDate string  `yaml:"issueDate"`
	SaleDate  string  `yaml:"saleDate"`
	DueDate   string  `yaml:"dueDate"`
	Notes     string  `yaml:"notes"`
	Company   Company `yaml:"company"`
	Bank      Bank    `yaml:"bank"`
	Items     []*Item `yaml:"items"`
	Currency  string  `yaml:"currency"`
	Signature string  `yaml:"signature"`
	Options   Options `yaml:"options"`
}

// Company details of buyer and seller
type Company struct {
	Buyer  Buyer  `yaml:"buyer"`
	Seller Seller `yaml:"seller"`
}

// Buyer company details
type Buyer struct {
	Name    string `yaml:"name"`
	Address string `yaml:"address"`
	VAT     string `yaml:"vat"`
}

// Seller company details
type Seller struct {
	Name    string `yaml:"name"`
	Address string `yaml:"address"`
	VAT     string `yaml:"vat"`
}

// Bank details on the invoice
type Bank struct {
	AccountNumber string `yaml:"accountNumber"`
	Swift         string `yaml:"swift"`
}

// Item parameters
type Item struct {
	Description string  `yaml:"description"`
	Quantity    float64 `yaml:"quantity"`
	UnitPrice   float64 `yaml:"unitPrice"`
	VATRate     float64 `yaml:"vatRate"`
}

// Options of the PDF document
type Options struct {
	FontFamily string `yaml:"font" default:"Arial"`
}
