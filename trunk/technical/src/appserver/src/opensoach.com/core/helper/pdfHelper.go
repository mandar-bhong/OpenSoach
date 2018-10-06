package helper

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/jung-kurt/gofpdf"
	gmodels "opensoach.com/models"
)

func CreatePdf(excelDataList []gmodels.ExcelData) (error, []byte) {

	pdf := newReport(excelDataList[0].StartDate, excelDataList[0].EndDate)

	pdf = header(pdf, excelDataList[0].Headers)
	pdf = table(pdf, excelDataList[0].Data)

	if pdf.Err() {
		log.Fatalf("Failed creating PDF report: %s\n", pdf.Error())
	}

	//Saving Pdf File
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

func newReport(startdate, enddate string) *gofpdf.Fpdf {

	pdf := gofpdf.New("L", "mm", "A4", "")

	pdf.SetHeaderFunc(func() {
		pdf.SetXY(5, 0)
		pdf.SetFont("Times", "", 10)
		pdf.Cell(0, 10, time.Now().Format("02/01/2006"))
		pdf.Ln(15)
	})

	pdf.SetFooterFunc(func() {
		pdf.SetXY(-15, -15)
		pdf.SetFont("Times", "", 8)
		pdf.CellFormat(0, 10, fmt.Sprintf("%d/{nb}", pdf.PageNo()), "", 0, "C", false, 0, "")
	})

	pdf.AliasNbPages("")
	pdf.AddPage()

	pdf.SetFont("Times", "", 28)
	pdf.Cell(40, 10, "Vehicle Service Report")
	pdf.Ln(15)

	fmt.Println(startdate, enddate)
	t1, _ := time.Parse("2006-01-02T15:04:05.000Z", startdate)
	t2, _ := time.Parse("2006-01-02T15:04:05.000Z", enddate)

	if startdate == enddate {
		pdf.SetFont("Times", "B", 14)
		pdf.Cell(40, 10, fmt.Sprintf("Date: %s", t1.Format("02/01/2006")))
		pdf.Ln(10)
	} else {
		pdf.SetFont("Times", "B", 14)
		pdf.Cell(40, 10, fmt.Sprintf("Date: %s - %s", t1.Format("02/01/2006"), t2.Format("02/01/2006")))
		pdf.Ln(10)
	}

	return pdf
}

func header(pdf *gofpdf.Fpdf, hdr []string) *gofpdf.Fpdf {

	pdf.SetFont("Times", "B", 11)
	pdf.SetFillColor(250, 144, 43)
	pdf.SetTextColor(255, 255, 255)

	_, lineHt := pdf.GetFontSize()
	marginCell := 2.
	curx, y := pdf.GetXY()
	x := curx
	colWidth := 25.
	height := 0.

	for _, str := range hdr {

		lines := pdf.SplitLines([]byte(str), colWidth)
		h := float64(len(lines))*lineHt + marginCell*float64(len(lines))
		if h > height {
			height = h
		}
	}

	for _, str := range hdr {

		pdf.Rect(x, y, colWidth, height, "FD")
		pdf.MultiCell(colWidth, lineHt+marginCell, str, "", "CM", false)
		x += colWidth
		pdf.SetXY(x, y)

	}
	pdf.SetXY(curx, y+height)

	return pdf
}

func table(pdf *gofpdf.Fpdf, tbl [][]string) *gofpdf.Fpdf {

	pdf.SetFillColor(240, 240, 240)
	pdf.SetTextColor(0, 0, 0)

	marginCell := 2.
	_, pageh := pdf.GetPageSize()
	_, _, _, mbottom := pdf.GetMargins()
	colWidth := 25.

	pdf.SetFont("Times", "", 10)

	alignList := []string{"C", "C", "C", "C", "C", "C", "C", "C", "C", "R", "R"}

	//set individual column width
	// colWidthList := []float64{}

	for _, line := range tbl {

		curx, y := pdf.GetXY()
		x := curx
		height := 0.
		_, lineHt := pdf.GetFontSize()

		for _, str := range line {

			lines := pdf.SplitLines([]byte(str), colWidth)
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
			width := colWidth
			pdf.Rect(x, y, width, height, "")
			pdf.MultiCell(width, lineHt+marginCell, str, "", alignList[i], false)
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
