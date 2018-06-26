package dbaccess

import (
	"github.com/jmoiron/sqlx"
	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	"opensoach.com/hkt/constants"
	"opensoach.com/hkt/constants/dbquery"
	hktmodels "opensoach.com/hkt/models"
)

var SUB_MODULE_NAME = "HKT.API.Report.DB"

func GetReportInfo(dbConn string, reportid int64) (error, *[]hktmodels.DBSplNodeReportTemplateTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetReportInfo")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBSplNodeReportTemplateTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.QueryType = dbmgr.AutoQuery
	selDBCtx.TableName = constants.DB_SPL_NODE_REPORT_TEMPLATE_TBL
	selDBCtx.Dest = data
	selErr := selDBCtx.SelectById(reportid)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func GetReportQueryData(dbConn string, query string, args ...interface{}) (error, []string, [][]string) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetReportQueryData")

	engine, err := sqlx.Connect("mysql", dbConn)

	if err != nil {
		return err, nil, nil
	}

	rows, err := engine.Queryx(query, args...)

	if err != nil {
		return err, nil, nil
	}

	cols, err := rows.Columns()
	if err != nil {
		//fmt.Println("Failed to get columns", err)
		return err, nil, nil
	}

	// Result is your slice string.
	rawResult := make([][]byte, len(cols))

	var resultRows [][]string

	dest := make([]interface{}, len(cols)) // A temporary interface{} slice
	for i, _ := range rawResult {
		dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
	}

	for rows.Next() {

		result := make([]string, len(cols))

		err = rows.Scan(dest...)
		if err != nil {
			//fmt.Println("Failed to scan row", err)
			return err, nil, nil
		}

		for i, raw := range rawResult {
			if raw == nil {
				result[i] = ""
			} else {
				result[i] = string(raw)
			}
		}

		resultRows = append(resultRows, result)
	}

	return nil, cols, resultRows
}

func GetReportShortDataList(dbConn string) (error, *[]hktmodels.DBReportTemplateShortDataModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetReportShortDataList")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBReportTemplateShortDataModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_SPL_NODE_REPORT_TEMPLATE_TABLE_SELECT_SHORT_DATA_LIST
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select()
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func GetReportInfoByCode(dbConn string, reportcode string) (error, *[]hktmodels.DBSplNodeReportTemplateTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetReportInfoByCode")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBSplNodeReportTemplateTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Query = dbquery.QUERY_SELECT_REPORT_TEMPLATE_BY_REPORT_CODE
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(reportcode)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}
