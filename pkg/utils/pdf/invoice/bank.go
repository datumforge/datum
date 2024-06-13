//nolint:mnd
package invoice

import (
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
)

// BuildBankDetails prepares rows with Bank details on the invoice
func (i *Invoice) BuildBankDetails() {
	i.pdf.SetBackgroundColor(getMulledWine())
	i.pdf.Line(0.5)
	i.pdf.SetBackgroundColor(getWinterSky())

	i.pdf.Row(20, func() {
		i.pdf.Col(3, func() {
			i.pdf.Text("Account no:", props.Text{
				Style: consts.Normal,
				Size:  8,
				Align: consts.Left,
				Color: getMulledWine(),
			})
			i.pdf.Text(i.Bank.AccountNumber, props.Text{
				Top:   3,
				Style: consts.Normal,
				Size:  8,
				Align: consts.Left,
			})
		})
		i.pdf.Col(2, func() {
			i.pdf.Text("Bank/SWIFT: ", props.Text{
				Style: consts.Normal,
				Size:  8,
				Align: consts.Left,
				Color: getMulledWine(),
			})
			i.pdf.Text(i.Bank.Swift, props.Text{
				Top:   3,
				Style: consts.Normal,
				Size:  8,
				Align: consts.Left,
			})
		})
	})
}
