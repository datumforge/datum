// nolint: gomnd
package invoice

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"

	"github.com/johnfercher/maroto/v2/pkg/props"
)

// BuildBankDetails prepares rows with Bank details on the invoice
func (i *Invoice) BuildBankDetails() {
	i.pdf.SetBackgroundColor(getMulledWine())
	i.pdf.Line(0.5)
	i.pdf.SetBackgroundColor(getWinterSky())

	i.pdf.Row(20, func() {
		i.pdf.Col(3, func() {
			i.pdf.Text("Account no:", props.Text{
				Style: fontstyle.Normal,
				Size:  8,
				Align: align.Left,
				Color: getMulledWine(),
			})
			i.pdf.Text(i.Bank.AccountNumber, props.Text{
				Top:   3,
				Style: fontstyle.Normal,
				Size:  8,
				Align: align.Left,
			})
		})
		i.pdf.Col(2, func() {
			i.pdf.Text("Bank/SWIFT: ", props.Text{
				Style: fontstyle.Normal,
				Size:  8,
				Align: align.Left,
				Color: getMulledWine(),
			})
			i.pdf.Text(i.Bank.Swift, props.Text{
				Top:   3,
				Style: fontstyle.Normal,
				Size:  8,
				Align: align.Left,
			})
		})
	})
}
