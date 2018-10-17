package helper

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/jung-kurt/gofpdf"
	gmodels "opensoach.com/models"
)

func CreatePdf(pdfdatamodel gmodels.PdfDataModel) (error, []byte) {

	pdf := newReport(pdfdatamodel.ReportName)

	for i := 0; i < len(pdfdatamodel.PdfData); i++ {

		if pdfdatamodel.PdfData[i].IsSummary == true {

			pdf.SetFont("Times", "", 20)
			pdf.CellFormat(30, 10, "Summary", "", 0, "", false, 0, "")
			pdf.Ln(15)

			x, y := pdf.GetXY()

			for j := 0; j < len(pdfdatamodel.PdfData[i].Headers); j++ {

				if j < len(pdfdatamodel.PdfData[i].Headers)/2 {
					pdf.SetFont("Times", "B", 12)
					pdf.CellFormat(50, 10, pdfdatamodel.PdfData[i].Headers[j], "", 0, "", false, 0, "")
					pdf.SetFont("Times", "", 12)
					pdf.CellFormat(30, 10, pdfdatamodel.PdfData[i].Data[0][j], "", 0, "", false, 0, "")
					pdf.Ln(8)
				} else {
					pdf.SetXY(x+140, y)
					pdf.SetFont("Times", "B", 12)
					pdf.CellFormat(50, 10, pdfdatamodel.PdfData[i].Headers[j], "", 0, "", false, 0, "")
					pdf.SetFont("Times", "", 12)
					pdf.CellFormat(30, 10, pdfdatamodel.PdfData[i].Data[0][j], "", 0, "", false, 0, "")
					pdf.Ln(8)
					y = y + 8
				}
			}
			pdf.Ln(10)

		} else {

			pdf = header(pdf, pdfdatamodel.PdfData[i].Headers, pdfdatamodel.PdfData[i].ColsWidth, pdfdatamodel.StartDate, pdfdatamodel.EndDate)
			pdf = table(pdf, pdfdatamodel.PdfData[i].Data, pdfdatamodel.PdfData[i].ColsAlign, pdfdatamodel.PdfData[i].ColsWidth)

		}
	}

	if pdf.Err() {
		log.Fatalf("Failed creating PDF report: %s\n", pdf.Error())
	}

	// Saving Pdf File
	// err := savePDF(pdf)
	// if err != nil {
	// 	log.Fatalf("Cannot save PDF: %s|n", err)
	// }

	var b bytes.Buffer

	err := pdf.Output(&b)
	if err != nil {
		log.Fatalf("failed to write in buffer: %s|n", err)
	}

	return nil, b.Bytes()

}

func newReport(reportname string) *gofpdf.Fpdf {

	pdf := gofpdf.New("L", "mm", "A4", "")

	pageWd, _ := pdf.GetPageSize()

	pdf.SetHeaderFunc(func() {
		pdf.Image("logo.png", 10, 6, 25, 0, false, "", 0, "")
		pdf.SetXY(-25, 6)
		pdf.SetFont("Times", "", 10)
		pdf.Cell(0, 10, time.Now().Format("02/01/2006"))
		pdf.Ln(15)
		x, y := pdf.GetXY()
		pdf.SetDrawColor(192, 192, 192)
		pdf.Line(x-3, y, x+pageWd-16, y)
		pdf.Ln(5)
	})

	pdf.SetFooterFunc(func() {
		pdf.SetDrawColor(192, 192, 192)
		pdf.SetXY(-4, -12)
		pdf.SetFont("Times", "", 8)
		pdf.CellFormat(0, 10, fmt.Sprintf("%d/{nb}", pdf.PageNo()), "", 0, "C", false, 0, "")
		x, y := pdf.GetXY()
		pdf.Line(x+4, y, x-pageWd+17, y)
	})

	pdf.AliasNbPages("")
	pdf.AddPage()

	pdf.SetX((pageWd - 40) / 2)
	pdf.SetFont("Times", "", 25)
	pdf.CellFormat(40, 10, reportname, "", 0, "C", false, 0, "")
	pdf.Ln(15)

	return pdf
}

func header(pdf *gofpdf.Fpdf, hdr []string, colsWidth []float64, startdate, enddate string) *gofpdf.Fpdf {

	pageWd, _ := pdf.GetPageSize()

	if startdate != "" && enddate != "" {
		t1, _ := time.Parse("2006-01-02T15:04:05.000Z", startdate)
		t2, _ := time.Parse("2006-01-02T15:04:05.000Z", enddate)

		if startdate == enddate {
			pdf.SetX((pageWd - 40) / 2)
			pdf.SetFont("Times", "B", 14)
			pdf.CellFormat(40, 10, fmt.Sprintf("Date: %s", t1.Format("Mon Jan 2 2006")), "", 0, "C", false, 0, "")
			pdf.Ln(10)
		} else {
			pdf.SetX((pageWd - 40) / 2)
			pdf.SetFont("Times", "B", 14)
			pdf.CellFormat(40, 10, fmt.Sprintf("Date: %s - %s", t1.Format("Mon Jan 2 2006"), t2.Format("Mon Jan 2 2006")), "", 0, "C", false, 0, "")
			pdf.Ln(10)
		}

	}

	pdf.SetFont("Times", "B", 11)
	pdf.SetFillColor(57, 103, 102)
	pdf.SetTextColor(255, 255, 255)

	_, lineHt := pdf.GetFontSize()
	marginCell := 3.
	curx, y := pdf.GetXY()
	x := curx
	height := 0.

	for i, str := range hdr {

		lines := pdf.SplitLines([]byte(str), colsWidth[i])
		h := float64(len(lines))*lineHt + marginCell*float64(len(lines))
		if h > height {
			height = h
		}
	}

	for i, str := range hdr {

		width := colsWidth[i]
		pdf.SetLineWidth(0.1)
		pdf.SetDrawColor(192, 192, 192)
		pdf.Rect(x, y, width, height, "FD")
		pdf.MultiCell(width, lineHt+marginCell, str, "", "CM", false)
		x += width
		pdf.SetXY(x, y)

	}
	pdf.SetXY(curx, y+height)

	return pdf
}

func table(pdf *gofpdf.Fpdf, tbl [][]string, colsAlign []string, colsWidth []float64) *gofpdf.Fpdf {

	pdf.SetFillColor(240, 240, 240)
	pdf.SetTextColor(0, 0, 0)

	marginCell := 3.
	_, pageh := pdf.GetPageSize()
	_, _, _, mbottom := pdf.GetMargins()

	pdf.SetFont("Times", "", 11)

	for _, line := range tbl {

		curx, y := pdf.GetXY()
		x := curx
		height := 0.
		_, lineHt := pdf.GetFontSize()

		for i, str := range line {

			lines := pdf.SplitLines([]byte(str), colsWidth[i])
			h := float64(len(lines))*lineHt + marginCell*float64(len(lines))
			if h > height {
				height = h
			}
		}

		// add a new page if the height of the row doesn't fit on the page
		if pdf.GetY()+height > pageh-mbottom {
			pdf.AddPage()
			y = pdf.GetY()
		}

		for i, str := range line {
			width := colsWidth[i]
			pdf.SetLineWidth(0.1)
			pdf.SetDrawColor(192, 192, 192)
			pdf.Rect(x, y, width, height, "")
			pdf.MultiCell(width, lineHt+marginCell, str, "", colsAlign[i], false)
			x += width
			pdf.SetXY(x, y)
		}
		pdf.SetXY(curx, y+height)

	}
	return pdf
}

func savePDF(pdf *gofpdf.Fpdf) error {
	return pdf.OutputFileAndClose("report.pdf")
}
