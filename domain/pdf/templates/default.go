package templates

import (
	"github.com/asentientbanana/pausalkole-admin/domain/pdf/util"
	"github.com/asentientbanana/pausalkole-admin/models"
	"github.com/asentientbanana/pausalkole-admin/utils"
	"github.com/phpdave11/gofpdf"
	"log"
	"strconv"
)

func GenerateDefaultInvoicePdf(invoice models.Invoice) {
	document := gofpdf.New("P", "mm", "A4", "")

	// Define header
	document.SetHeaderFunc(func() {
		document.SetFont("Arial", "B", 14)
		document.Cell(0, 10, "Invoice ID: "+invoice.ID.String())
		document.Ln(20) // line break
	})

	document.AddPage()
	document.SetFont("Arial", "", 12)

	// Column widths
	colWidth := 90.0 // half of A4 (minus margins)
	lineHeight := 8.0

	leftX := 20.0
	rightX := 110.0 // roughly half page width
	y := 40.0

	const EntityRowMargin = 4

	// Agency
	document.CellFormat(colWidth, lineHeight, invoice.Agency.Name, "0", 0, "", false, 0, "")
	for i, agencyField := range invoice.Agency.Fields {
		_y := y + float64(i*EntityRowMargin)
		document.Text(leftX, _y, agencyField.Field+": "+agencyField.Value)
	}

	invoice.Recipient.Fields = append(invoice.Recipient.Fields, invoice.Agency.Fields...)
	invoice.Recipient.Fields = append(invoice.Recipient.Fields, invoice.Agency.Fields...)
	invoice.Recipient.Fields = append(invoice.Recipient.Fields, invoice.Agency.Fields...)
	invoice.Recipient.Fields = append(invoice.Recipient.Fields, invoice.Agency.Fields...)
	invoice.Recipient.Fields = append(invoice.Recipient.Fields, invoice.Agency.Fields...)

	invoice.Items = append(invoice.Items, invoice.Items...)
	invoice.Items = append(invoice.Items, invoice.Items...)
	invoice.Items = append(invoice.Items, invoice.Items...)
	invoice.Items = append(invoice.Items, invoice.Items...)
	invoice.Items = append(invoice.Items, invoice.Items...)

	infoHeightOffset, entityInfoRowNumber := util.CalculatePdfEntityInfoHeight(&invoice.Recipient, &invoice.Agency)

	//Calculate available space for items on the first page
	initialRowLimit := 15
	rowLimit := initialRowLimit
	//Magic number gotten with testing
	if entityInfoRowNumber > 40 {
		rowLimit = 5
	} else if entityInfoRowNumber > 30 {
		rowLimit = 7
	} else if entityInfoRowNumber > 15 {
		rowLimit = 10
	} else {
		rowLimit = entityInfoRowNumber
	}

	paginatedItems := util.GetPdfRowsPerPage(invoice, initialRowLimit, rowLimit)
	//Recipient

	pageCount := len(paginatedItems)
	for index, page := range paginatedItems {
		if index == 0 {
			document.CellFormat(colWidth, lineHeight, invoice.Recipient.Name, "0", 0, "", false, 0, "")
			for i, agencyField := range invoice.Recipient.Fields {
				_y := y + float64(i*EntityRowMargin)
				document.Text(rightX, _y, agencyField.Field+": "+agencyField.Value)

			}
			//Offset the element under the Agency and recipient info
			document.Ln(infoHeightOffset - 20)

			// Write a title
			document.Cell(0, 10, "Description")
			document.Ln(6)
			document.Cell(0, 10, invoice.Description)
			document.Ln(12)
		}

		GenerateDefaultInvoiceItemTable(document, page, invoice.Items)

		// Define footer
		document.SetFooterFunc(func() {
			document.SetY(-15) // position 15mm from bottom
			document.SetFont("Arial", "I", 8)
			document.CellFormat(0, 10,
				// Page number
				"Page "+strconv.Itoa(document.PageCount()),
				"", 0, "C", false, 0, "")
		})
		if pageCount != index+1 {
			document.AddPage()
		}
	}
	// Save to file
	err := document.OutputFileAndClose("list_with_header_footer.pdf")
	if err != nil {
		log.Fatal(err)
	}
}

func GenerateDefaultInvoiceItemTable(pdf *gofpdf.Fpdf, indexes []int, items []models.InvoiceItem) {

	pdf.CellFormat(10, 10, "#", "0", 0, "C", false, 0, "")
	pdf.CellFormat(90, 10, "Description", "0", 0, "C", false, 0, "")
	pdf.CellFormat(30, 10, "Metric", "0", 0, "C", false, 0, "")
	pdf.CellFormat(20, 10, "Quantity", "0", 0, "C", false, 0, "")
	pdf.CellFormat(20, 10, "Amount", "0", 0, "C", false, 0, "")
	pdf.Ln(-1)
	for i, row := range indexes {
		item := items[row]
		amount := utils.FormatFloatUtil(float64(item.Amount))

		pdf.CellFormat(10, 10, strconv.Itoa(i+1), "1", 0, "C", false, 0, "")
		pdf.CellFormat(90, 10, item.Description, "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 10, item.Metric, "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 10, strconv.Itoa(item.Quantity), "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 10, amount, "1", 0, "C", false, 0, "")
		pdf.Ln(-1)
	}
}
