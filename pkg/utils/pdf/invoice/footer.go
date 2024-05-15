// nolint: gomnd
package invoice

import (
	"fmt"
	"strconv"

	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

// BuildPDFFooter prepares footer on the invoice
func (i *Invoice) BuildPDFFooter() {
	i.pdf.RegisterFooter(func() {
		i.pdf.SetAliasNbPages("{nbs}")
		currentPage := strconv.Itoa(i.pdf.GetCurrentPage())
		i.pdf.Row(6, func() {
			i.pdf.Col(12, func() {
				i.pdf.Text("Need help? See your invoice by logging into console.datum.net", props.Text{
					Top:   1,
					Style: fontstyle.Normal,
					Size:  8,
					Align: align.Left,
					Color: getMulledWine(),
				})
			})
		})
		i.pdf.Row(6, func() {
			i.pdf.Col(12, func() {
				i.pdf.Text(fmt.Sprintf("Page %s of {nbs}", currentPage), props.Text{
					Top:   1,
					Style: fontstyle.Normal,
					Size:  8,
					Align: align.Center,
					Color: getMulledWine(),
				})
			})
		})
	})
}
