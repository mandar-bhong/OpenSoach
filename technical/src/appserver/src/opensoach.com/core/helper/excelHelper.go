package helper

import (
	"bufio"
	"bytes"

	"github.com/tealeg/xlsx"
	gmodels "opensoach.com/models"
)

func CreateExcel(excelDataList []gmodels.ExcelData) (error, []byte) {

	file := xlsx.NewFile()

	for i := 0; i < len(excelDataList); i++ {

		if excelDataList[i].IsVertical == false {

			sheet, err := file.AddSheet(excelDataList[i].SheetName)

			if err != nil {
				return err, nil
			}

			for j := 0; j < len(excelDataList[i].Headers); j++ {
				row := sheet.AddRow()
				row.AddCell().Value = excelDataList[i].Headers[j]
				for _, data := range excelDataList[i].Data {
					row.AddCell().Value = data[j]
				}
			}

		} else {
			sheet, err := file.AddSheet(excelDataList[i].SheetName)

			if err != nil {
				return err, nil
			}

			row := sheet.AddRow()
			for _, headerItem := range excelDataList[i].Headers {
				row.AddCell().Value = headerItem
			}

			for _, rowlist := range excelDataList[i].Data {
				row = sheet.AddRow()
				for _, colData := range rowlist {
					row.AddCell().Value = colData
				}
			}
		}
	}

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	file.Write(writer)

	return nil, b.Bytes()
}
