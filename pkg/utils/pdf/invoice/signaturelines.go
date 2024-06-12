//nolint:mnd
package invoice

import (
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
)

// BuildSignatureLines prepares signatures of the buyer and seller
func (i *Invoice) BuildSignatureLines() {
	i.pdf.SetBackgroundColor(getMulledWine())
	i.pdf.Line(0.5)
	i.pdf.SetBackgroundColor(getWinterSky())

	i.pdf.Row(15, func() {
		i.pdf.Col(1, func() {
			i.pdf.Text("Notes:", props.Text{
				Top:   1,
				Style: consts.Bold,
				Size:  8,
				Align: consts.Left,
				Color: getMulledWine(),
			})
		})
		i.pdf.Col(3, func() {
			i.pdf.Text(i.Notes, props.Text{
				Top:   1,
				Style: consts.Italic,
				Size:  8,
				Align: consts.Left,
			})
		})
	})

	i.pdf.Row(15, func() {
		i.pdf.Col(5, func() {
			i.pdf.Signature("Buyer signature", props.Font{
				Size:   12.0,
				Style:  consts.Normal,
				Family: consts.Helvetica,
				Color: color.Color{
					Red:   10,
					Green: 20,
					Blue:  30,
				},
			})
		})
		i.pdf.ColSpace(1)
		i.pdf.Col(5, func() {
			i.pdf.Text(i.Signature, props.Text{
				Top:    5,
				Style:  consts.Normal,
				Family: consts.Helvetica,
				Size:   8,
				Align:  consts.Center,
			})
			i.pdf.Signature("Seller signature", props.Font{
				Size:   12.0,
				Style:  consts.Normal,
				Family: consts.Helvetica,
				Color: color.Color{
					Red:   10,
					Green: 20,
					Blue:  30,
				},
			})
		})
	})
}
