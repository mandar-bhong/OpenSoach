package helper

import (
	"bufio"
	"bytes"

	"github.com/tealeg/xlsx"
	gmodels "opensoach.com/models"
)

func CreateExcel(excelData gmodels.ExcelData) (error, []byte) {

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")

	if err != nil {
		return err, nil
	}

	row := sheet.AddRow()
	for _, headerItem := range excelData.Headers {
		row.AddCell().Value = headerItem
	}

	for _, rowlist := range excelData.Data {
		row = sheet.AddRow()
		for _, colData := range rowlist {
			row.AddCell().Value = colData
		}
	}

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	file.Write(writer)

	return nil, b.Bytes()
}
