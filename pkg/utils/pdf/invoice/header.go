// nolint: gomnd
package invoice

import (
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
)

// BuildInvoiceHeader prepares header on the invoice
func (i *Invoice) BuildInvoiceHeader() {
	i.pdf.RegisterHeader(func() {
		i.pdf.Row(20, func() {
			//			i.pdf.ColSpace(8)
			i.pdf.Col(4, func() {
				err := i.pdf.FileImage("pkg/utils/pdf/assets/images/logo.png", props.Rect{
					Center: true,
					//					Percent: 75,
				})
				if err != nil {
					panic("Could not load image")
				}
			})
		})
		i.pdf.Row(30, func() {
			i.pdf.Col(5, func() {
				i.pdf.Text("Invoice", props.Text{
					Size:  24,
					Style: consts.Normal,
					Align: consts.Left,
				})
				i.pdf.Text(i.Number, props.Text{
					Top:   12,
					Size:  18,
					Style: consts.Normal,
				})
				i.pdf.ColSpace(8)
			})
			i.pdf.Row(20, func() {
				i.pdf.ColSpace(8)
				i.pdf.Col(4, func() {
					i.pdf.Text("Date of issue:", props.Text{
						Size:  8,
						Style: consts.Normal,
						Align: consts.Center,
						Color: getMulledWine(),
					})
					i.pdf.Text(i.IssueDate, props.Text{
						Size:  8,
						Style: consts.Normal,
						Align: consts.Right,
					})
					i.pdf.Text("Date of sale:", props.Text{
						Top:   6,
						Size:  8,
						Style: consts.Normal,
						Align: consts.Center,
						Color: getMulledWine(),
					})
					i.pdf.Text(i.SaleDate, props.Text{
						Top:   6,
						Size:  8,
						Style: consts.Normal,
						Align: consts.Right,
					})
					i.pdf.Text("Due date:", props.Text{
						Top:   12,
						Size:  8,
						Style: consts.Normal,
						Align: consts.Center,
						Color: getMulledWine(),
					})
					i.pdf.Text(i.DueDate, props.Text{
						Top:   12,
						Size:  8,
						Style: consts.Normal,
						Align: consts.Right,
					})
				})
			})
		})
	})
}
