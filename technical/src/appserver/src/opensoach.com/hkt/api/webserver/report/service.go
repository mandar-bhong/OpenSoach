package report

import (
	"strconv"
	"strings"

	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	lmodels "opensoach.com/hkt/api/models"
	"opensoach.com/hkt/api/webserver/report/dbaccess"
	hktmodels "opensoach.com/hkt/models"
	gmodels "opensoach.com/models"
	pcmodels "opensoach.com/prodcore/models"
)

var SUB_MODULE_NAME = "HKT.API.Report"

type ReportService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service ReportService) GenerateReport(req lmodels.APIGenerateReportRequestModel) (bool, interface{}) {

	issuccess, respData := service.ViewReport(req.APIViewReportRequestModel)
	if issuccess == false {
		logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Failed to generate report failed to get report data")
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	exceldatalist := []gmodels.ExcelData{}

	if len(respData.([]hktmodels.DBGetReportDataModel)) < 2 {
		logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Failed to generate report insufficient report data")
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	taskSummaryData := respData.([]hktmodels.DBGetReportDataModel)[0]

	taskDetailsData := respData.([]hktmodels.DBGetReportDataModel)[1]

	exceldata := gmodels.ExcelData{}
	exceldata.SheetName = "Summary"
	exceldata.IsVertical = false
	exceldata.Headers = taskSummaryData.ReportHeader
	exceldata.Data = taskSummaryData.ReportData
	exceldatalist = append(exceldatalist, exceldata)

	exceldata = gmodels.ExcelData{}
	exceldata.SheetName = "Details"
	exceldata.IsVertical = true
	exceldata.Headers = taskDetailsData.ReportHeader
	exceldata.Data = taskDetailsData.ReportData
	exceldatalist = append(exceldatalist, exceldata)

	err, data := ghelper.CreateExcel(exceldatalist)
	if err != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while Creating Excel file.", err)
	}

	documentData := pcmodels.DocumentData{}
	documentData.ByteData = data
	documentData.ContentType = "attachment"

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully Created Report Excel File")

	return true, documentData

}

func (service ReportService) ViewReport(reqList lmodels.APIViewReportRequestModel) (bool, interface{}) {

	respList := []hktmodels.DBGetReportDataModel{}

	for _, req := range reqList.ReportRequest {

		dbErr, reportData := dbaccess.GetReportInfoByCode(service.ExeCtx.SessionInfo.Product.NodeDbConn, req.ReportCode)

		if dbErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting report info by report code.", dbErr)

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

		Query := strings.Replace(reportDataRecord[0].ReportQuery, "$WhereCpmIdValue$", strconv.FormatInt(service.ExeCtx.SessionInfo.Product.CustProdID, 10), 1)

		dberr, _, resultRows := dbaccess.GetReportQueryData(service.ExeCtx.SessionInfo.Product.NodeDbConn, Query, req.QueryParams...)
		if dberr != nil {
			logger.Context().WithField("DBErr", dberr.Error()).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting report query data.", dbErr)

			errModel := gmodels.APIResponseError{}
			errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
			return false, errModel
		}

		headerModel := hktmodels.ReportHeaderModel{}

		isJsonConvertSuccess := ghelper.ConvertFromJSONString(reportDataRecord[0].ReportHeader, &headerModel)

		if isJsonConvertSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Failed to convert from json string")
			return false, nil
		}

		reportDataModel := hktmodels.DBGetReportDataModel{}
		reportDataModel.ReportCode = reportDataRecord[0].ReportCode

		if req.Language == "en" {
			reportDataModel.ReportHeader = headerModel.En
		} else {
			reportDataModel.ReportHeader = headerModel.Hi
		}

		if resultRows == nil {
			reportDataModel.ReportData = [][]string{}
		} else {
			reportDataModel.ReportData = resultRows
		}

		respList = append(respList, reportDataModel)

	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched report data")

	return true, respList

}

func (service ReportService) ReportShortList() (bool, interface{}) {

	dbErr, listData := dbaccess.GetReportShortDataList(service.ExeCtx.SessionInfo.Product.NodeDbConn)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting report short data list.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched report short data list.")

	return true, listData

}
