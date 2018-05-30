package report

import (
	core "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	"opensoach.com/hkt/api/webserver/report/dbaccess"
	hktmodels "opensoach.com/hkt/models"
	gmodels "opensoach.com/models"
)

var SUB_MODULE_NAME = "HKT.API.Report"

type ReportService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service ReportService) GenerateReport(reportID int64) (bool, interface{}) {

	dbErr, reportData := dbaccess.GetReportInfo(service.ExeCtx.SessionInfo.Product.NodeDbConn, reportID)

	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *reportData

	if len(dbRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	dberr, cols, resultRows := dbaccess.GetReportQueryData(service.ExeCtx.SessionInfo.Product.NodeDbConn, dbRecord[0].ReportQuery)
	if dberr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	exceldata := gmodels.ExcelData{}
	exceldata.Headers = cols
	exceldata.Data = resultRows

	err, data := core.CreateExcel(exceldata)
	if err != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while Creating Excel file.", dbErr)
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully Created Report Excel File")

	return true, data

}

func (service ReportService) ViewReport(reportID int64) (bool, interface{}) {

	dbErr, reportData := dbaccess.GetReportInfo(service.ExeCtx.SessionInfo.Product.NodeDbConn, reportID)

	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *reportData

	if len(dbRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	dberr, cols, resultRows := dbaccess.GetReportQueryData(service.ExeCtx.SessionInfo.Product.NodeDbConn, dbRecord[0].ReportQuery)
	if dberr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	reportDataModel := hktmodels.DBGetReportDataModel{}
	reportDataModel.ReportId = dbRecord[0].ReportId
	reportDataModel.ReportCode = dbRecord[0].ReportCode
	reportDataModel.ReportDesc = dbRecord[0].ReportDesc
	reportDataModel.ReportHeader = cols
	reportDataModel.ReportData = resultRows

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully Created Report Excel File")

	return true, reportDataModel

}

func (service ReportService) ReportShortList() (bool, interface{}) {

	dbErr, listData := dbaccess.GetReportShortDataList(service.ExeCtx.SessionInfo.Product.NodeDbConn)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched report short data list.")

	return true, listData

}
