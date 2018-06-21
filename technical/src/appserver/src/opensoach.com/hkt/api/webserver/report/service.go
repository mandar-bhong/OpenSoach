package report

import (
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	lmodels "opensoach.com/hkt/api/models"
	"opensoach.com/hkt/api/webserver/report/dbaccess"
	hktmodels "opensoach.com/hkt/models"
	gmodels "opensoach.com/models"
)

var SUB_MODULE_NAME = "HKT.API.Report"

type ReportService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service ReportService) GenerateReport(req hktmodels.DBGenerateReportRequestDataModel) (bool, interface{}) {

	dbErr, reportData := dbaccess.GetReportInfo(service.ExeCtx.SessionInfo.Product.NodeDbConn, req.ReportID)

	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	reportDataRecord := *reportData

	if len(reportDataRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	dberr, _, resultRows := dbaccess.GetReportQueryData(service.ExeCtx.SessionInfo.Product.NodeDbConn, reportDataRecord[0].ReportQuery)
	if dberr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	headerModel := hktmodels.ReportHeaderModel{}

	isJsonConvertSuccess := ghelper.ConvertFromJSONString(reportDataRecord[0].ReportHeader, &headerModel)

	if isJsonConvertSuccess == false {

	}

	exceldata := gmodels.ExcelData{}
	exceldata.Data = resultRows

	if req.Language == "en" {
		exceldata.Headers = headerModel.En
	} else {
		exceldata.Headers = headerModel.Hi
	}

	err, data := ghelper.CreateExcel(exceldata)
	if err != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while Creating Excel file.", dbErr)
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully Created Report Excel File")

	return true, data

}

func (service ReportService) ViewReport(req hktmodels.DBGenerateReportRequestDataModel) (bool, interface{}) {

	dbErr, reportData := dbaccess.GetReportInfo(service.ExeCtx.SessionInfo.Product.NodeDbConn, req.ReportID)

	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	reportDataRecord := *reportData

	if len(reportDataRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	dberr, _, resultRows := dbaccess.GetReportQueryData(service.ExeCtx.SessionInfo.Product.NodeDbConn, reportDataRecord[0].ReportQuery, req.QueryParams...)
	if dberr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	reportDataModel := hktmodels.DBGetReportDataModel{}
	reportDataModel.ReportId = reportDataRecord[0].ReportId
	reportDataModel.ReportCode = reportDataRecord[0].ReportCode
	reportDataModel.ReportDesc = reportDataRecord[0].ReportDesc
	reportDataModel.ReportHeader = reportDataRecord[0].ReportHeader
	reportDataModel.ReportData = resultRows

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched report data")

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

func (service ReportService) ReportLocationSummary(req lmodels.APIReportLocationSummaryRequest) (bool, interface{}) {

	dberr, reportData := dbaccess.GetReportInfo(service.ExeCtx.SessionInfo.Product.NodeDbConn, req.ReportID)

	if dberr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dberr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	reportDataRecord := *reportData

	if len(reportDataRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	filterModel := hktmodels.DBReportLocationSummaryFilterDataModel{}
	filterModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	filterModel.SpId = req.SpID

	dbErr, reportLocationDataModels := dbaccess.GetReportLocationSummary(service.ExeCtx.SessionInfo.Product.NodeDbConn, req, filterModel)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	apiResponse := lmodels.APIReportLocationSummaryResponse{}

	reportTaskSummaryModelList := []hktmodels.ReportTaskSummaryModel{}

	var reportTaskSummaryMap = map[string]*hktmodels.ReportTaskSummary{}

	for _, dbSummaryDataModel := range reportLocationDataModels {

		value, hasValue := reportTaskSummaryMap[dbSummaryDataModel.TaskName]

		if hasValue == false {
			value = &hktmodels.ReportTaskSummary{}
			reportTaskSummaryMap[dbSummaryDataModel.TaskName] = value
		}

		if dbSummaryDataModel.Status == 1 {
			value.Ontime = dbSummaryDataModel.Count
		} else {
			value.Delayed = dbSummaryDataModel.Count
		}
	}

	for key, value := range reportTaskSummaryMap {
		reportTaskSummaryModel := hktmodels.ReportTaskSummaryModel{}
		reportTaskSummaryModel.Taskname = key
		reportTaskSummaryModel.Ontime = value.Ontime
		reportTaskSummaryModel.Delayed = value.Delayed
		reportTaskSummaryModel.Total = value.Ontime + value.Delayed
		reportTaskSummaryModelList = append(reportTaskSummaryModelList, reportTaskSummaryModel)
	}

	apiResponse.ReportHeader = reportDataRecord[0].ReportHeader
	apiResponse.ReportTaskSummary = reportTaskSummaryModelList

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully report location summary per month")
	return true, apiResponse
}
