package dbaccess

import (
	"strings"

	"github.com/jmoiron/sqlx"
	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	hkthelper "opensoach.com/hkt/api/helper"
	lmodels "opensoach.com/hkt/api/models"
	"opensoach.com/hkt/constants"
	"opensoach.com/hkt/constants/dbquery"
	hktmodels "opensoach.com/hkt/models"
	pcconst "opensoach.com/prodcore/constants"
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

func GetReportLocationSummary(dbConn string, req lmodels.APIReportLocationSummaryRequest, filtermodel hktmodels.DBReportLocationSummaryFilterDataModel) (error, []hktmodels.DBReportLocationSummaryDataModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetReportLocationSummary")

	data := []hktmodels.DBReportLocationSummaryDataModel{}

	whereCondition := hkthelper.GetFilterConditionFormModel(filtermodel)

	if req.StartDate != nil && req.EndDate != nil {

		if whereCondition != "" {
			whereCondition = whereCondition + " and "
		}

		dbStartTime := req.StartDate.Format(pcconst.DB_TIME_FORMAT)
		dbEndTime := req.EndDate.Format(pcconst.DB_TIME_FORMAT)

		whereCondition = whereCondition + " txn_date between '" + dbStartTime + "' and '" + dbEndTime + "'"
	}

	if whereCondition != "" {
		whereCondition = " where " + whereCondition
	}

	query := strings.Replace(dbquery.QUERY_GET_REPORT_TASK_SUMMARY_PER_MONTH, "$WhereCondition$", whereCondition, 1)

	selectCtx := dbmgr.SelectContext{}
	selectCtx.DBConnection = dbConn
	selectCtx.Dest = &data
	selectCtx.Query = query
	selectCtx.QueryType = dbmgr.Query
	selectCtxErr := selectCtx.Select()
	if selectCtxErr != nil {
		return selectCtxErr, nil
	}

	return nil, data
}
