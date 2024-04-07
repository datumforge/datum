package tables

import (
	"fmt"
	"io"
	"strings"

	"github.com/olekukonko/tablewriter"
)

const colWidth = 50

// TableOutputWriter is the interface to write out tables
type TableOutputWriter interface {
	SetHeaders(headers ...string)
	AddRow(items ...interface{})
	Render()
	RowCount() int
}

// convertToUpper will make sure all entries are UPPERCASED ALMOST LIKE THEY ARE SCREAMING AT YOU
func convertToUpper(headers []string) []string {
	head := []string{}
	for _, item := range headers {
		head = append(head, strings.ToUpper(item))
	}

	return head
}

// NewTableWriter gets a new instance of our table output writer
func NewTableWriter(output io.Writer, headers ...string) TableOutputWriter {
	table := tablewriter.NewWriter(output)
	table.SetBorder(false)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderLine(false)
	table.SetColWidth(colWidth)
	table.SetTablePadding("\t\t")
	table.SetHeader(convertToUpper(headers))

	t := &tableoutputwriter{}
	t.out = output
	t.table = table

	return t
}

// tableoutputwriter is our internal implementation of the TableOutputWriter
type tableoutputwriter struct {
	out   io.Writer
	table *tablewriter.Table
}

// SetHeaders sets the headers for our table and coverts to UPPERCASE for everyones viewing pleasure
func (t *tableoutputwriter) SetHeaders(headers ...string) {
	t.table.SetHeader(convertToUpper(headers))
}

// AddRow appends a new row to our table
func (t *tableoutputwriter) AddRow(items ...interface{}) {
	row := []string{}

	for _, item := range items {
		row = append(row, fmt.Sprintf("%v", item))
	}

	t.table.Append(row)
}

// RowCount gets the number of rows in the table
func (t *tableoutputwriter) RowCount() int {
	return t.table.NumLines()
}

// Render emits the generated table to the output once ready
func (t *tableoutputwriter) Render() {
	t.table.Render()

	// ensures a break line after we flush the tabwriter
	fmt.Fprintln(t.out)
}
