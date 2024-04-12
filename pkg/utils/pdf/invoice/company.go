// nolint: gomnd
package invoice

import (
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
)

// BuildCompanyDetails prepares rows with Buyer and Seller contact details on the invoice
func (i *Invoice) BuildCompanyDetails() {
	i.pdf.Row(7, func() {
		i.pdf.SetBackgroundColor(getMulledWine())
		i.pdf.Col(3, func() {
			i.pdf.Text("Seller", props.Text{
				Top:   1.5,
				Size:  9,
				Style: consts.Normal,
				Align: consts.Center,
				Color: getWinterSky(),
			})
		})
		i.pdf.ColSpace(4)
		i.pdf.Col(5, func() {
			i.pdf.Text("Buyer", props.Text{
				Top:   1.5,
				Size:  9,
				Style: consts.Normal,
				Align: consts.Center,
				Color: getWinterSky(),
			})
		})
	})

	i.pdf.SetBackgroundColor(getWinterSky())
	i.pdf.Row(10, func() {
		i.pdf.Col(2, func() {
			i.pdf.Text("Name:  ", props.Text{
				Top:   2,
				Style: consts.Normal,
				Align: consts.Left,
				Color: getMulledWine(),
			})
		})
		i.pdf.Col(3, func() {
			i.pdf.Text(i.Company.Seller.Name, props.Text{
				Top:   2,
				Style: consts.Normal,
				Align: consts.Left,
			})
		})
		i.pdf.ColSpace(2)
		i.pdf.Col(2, func() {
			i.pdf.Text("Name:  ", props.Text{
				Top:   2,
				Style: consts.Normal,
				Align: consts.Left,
				Color: getMulledWine(),
			})
		})
		i.pdf.Col(3, func() {
			i.pdf.Text(i.Company.Buyer.Name, props.Text{
				Top:   2,
				Style: consts.Normal,
				Align: consts.Left,
			})
		})
	})
	i.pdf.Row(10, func() {
		i.pdf.Col(2, func() {
			i.pdf.Text("Address:  ", props.Text{
				Top:   3,
				Style: consts.Normal,
				Align: consts.Left,
				Color: getMulledWine(),
			})
		})
		i.pdf.Col(3, func() {
			i.pdf.Text(i.Company.Seller.Address, props.Text{
				Top:   3,
				Style: consts.Normal,
				Align: consts.Left,
			})
		})
		i.pdf.ColSpace(2)
		i.pdf.Col(2, func() {
			i.pdf.Text("Address:  ", props.Text{
				Top:   3,
				Style: consts.Normal,
				Align: consts.Left,
				Color: getMulledWine(),
			})
		})
		i.pdf.Col(3, func() {
			i.pdf.Text(i.Company.Buyer.Address, props.Text{
				Top:   2,
				Style: consts.Normal,
				Align: consts.Left,
			})
		})
	})
	i.pdf.Row(7, func() {
		i.pdf.Col(2, func() {
			i.pdf.Text("VAT Number:  ", props.Text{
				Top:   3,
				Style: consts.Normal,
				Align: consts.Left,
				Color: getMulledWine(),
			})
		})
		i.pdf.Col(3, func() {
			i.pdf.Text(i.Company.Seller.VAT, props.Text{
				Top:   3,
				Style: consts.Normal,
				Align: consts.Left,
			})
		})
		i.pdf.ColSpace(2)
		i.pdf.Col(2, func() {
			i.pdf.Text("VAT Number:  ", props.Text{
				Top:   3,
				Style: consts.Normal,
				Align: consts.Left,
				Color: getMulledWine(),
			})
		})
		i.pdf.Col(3, func() {
			i.pdf.Text(i.Company.Buyer.VAT, props.Text{
				Top:   3,
				Style: consts.Normal,
				Align: consts.Left,
			})
		})
	})
	i.pdf.Row(2, func() { // intentionally blank
	})
}
