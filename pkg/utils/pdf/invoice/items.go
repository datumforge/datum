// nolint: gomnd
package invoice

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
)

// BuildInvoiceLineItems prepares Tablelist with items on the invoice with calculated tax amounts and total gross amounts
func (i *Invoice) BuildInvoiceLineItems() {
	backgroundColor := getWinterSky()
	header := getHeader()
	items := i.getItems()
	taxes, totals := i.countTax()
	contents := appendItems(items, taxes, totals)

	i.pdf.SetBackgroundColor(getMulledWine())
	i.pdf.Row(2, func() {
		i.pdf.Col(12, func() { // intentionally empty
		})
	})
	i.pdf.SetBackgroundColor(getWinterSky())
	i.pdf.TableList(header, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Style:     consts.Normal,
			Size:      8,
			GridSizes: []uint{1, 3, 1, 2, 1, 1, 3},
			Color:     getMulledWine(),
		},
		ContentProp: props.TableListContent{
			Style:     consts.Normal,
			Size:      10,
			GridSizes: []uint{1, 3, 1, 2, 1, 1, 3},
		},
		Align:                consts.Center,
		AlternatedBackground: &backgroundColor,
		HeaderContentSpace:   1,
		Line:                 false,
	})

	i.pdf.Row(2, func() { // intentionally empty - created some padding between the total box and the line items above
	})

	i.pdf.Row(8, func() {
		i.pdf.ColSpace(8)
		i.pdf.SetBackgroundColor(getMulledWine())
		i.pdf.Col(2, func() {
			i.pdf.Text("Total:", props.Text{
				Top:   3,
				Style: consts.Normal,
				Size:  8,
				Align: consts.Center,
				Color: getWinterSky(),
			})
		})
		i.pdf.Col(2, func() {
			i.pdf.Text(fmt.Sprintf("%s %s", calculateInvoiceSum(contents), i.Currency), props.Text{
				Top:   3,
				Style: consts.Normal,
				Size:  8,
				Align: consts.Center,
				Color: getWinterSky(),
			})
		})
	})
}

func getHeader() []string {
	return []string{"No", "Description", "Quantity", "Unit net price", "VAT rate", "VAT amount", "Total gross price"}
}

// calculateInvoiceSum iterating over a slice of strings representing the items on the invoice. It extracts the total gross
// price of each item from the last element of each row in the slice, converts it to a float value, and
// accumulates these values to calculate the overall sum of the invoice. The function then returns the
// total sum as a string with two decimal places
func calculateInvoiceSum(values [][]string) string {
	var sum float64

	for _, value := range values {
		num, _ := strconv.ParseFloat(value[len(value)-1], 64)
		sum += num
	}

	return strconv.FormatFloat(sum, 'f', 2, 64)
}

// countTax calculates the tax amount and total gross price for each item in the invoice
func (i *Invoice) countTax() ([]float64, []float64) {
	var (
		taxes  []float64
		totals []float64
	)

	for _, item := range i.Items {
		vat := item.VATRate
		price := item.UnitPrice
		quantity := item.Quantity

		tax := quantity * (vat * price / 100)
		total := quantity*price + tax

		taxes = append(taxes, tax)
		totals = append(totals, total)
	}

	return taxes, totals
}

// appendItems takes a juicy slice of strings along with two juicy slices of float64 values representing taxes
// and total amounts - it appends the tax and total amounts to each row in the input slice, along with a
// sequential number at the beginning of each row it returns the modified slice with the added
// tax and total amounts for each item
func appendItems(values [][]string, taxes []float64, totals []float64) [][]string {
	for i := range values {
		values[i] = append([]string{strconv.Itoa(i)}, values[i]...)
		values[i] = append(values[i], strconv.FormatFloat(taxes[i], 'f', 2, 64))
		values[i] = append(values[i], strconv.FormatFloat(totals[i], 'f', 2, 64))
	}

	return values
}

// getItems retrieves the items from the `Items` field of the `Invoice` struct and converts them into a
// juicy slice of strings
func (i *Invoice) getItems() [][]string {
	var items [][]string

	v := reflect.Indirect(reflect.ValueOf(i.Items))
	if v.Kind() != reflect.Slice {
		return nil
	}

	for j := 0; j < v.Len(); j++ {
		e := reflect.Indirect(v.Index(j))

		if e.Kind() != reflect.Struct {
			return nil
		}

		var element []string

		for j := 0; j < e.NumField(); j++ {
			element = append(element, fmt.Sprint(e.Field(j).Interface()))
		}

		items = append(items, element)
	}

	return items
}
