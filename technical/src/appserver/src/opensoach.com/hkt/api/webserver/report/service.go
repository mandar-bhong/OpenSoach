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
)

var SUB_MODULE_NAME = "HKT.API.Report"

type ReportService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service ReportService) GenerateReport(req hktmodels.DBReportRequestDataModel) (bool, interface{}) {

	dbErr, reportData := dbaccess.GetReportInfoByCode(service.ExeCtx.SessionInfo.Product.NodeDbConn, req.ReportCode)

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

func (service ReportService) ViewReport(reqList lmodels.APIViewReportRequestModel) (bool, interface{}) {

	respList := []hktmodels.DBGetReportDataModel{}

	for _, req := range reqList.ReportRequest {

		dbErr, reportData := dbaccess.GetReportInfoByCode(service.ExeCtx.SessionInfo.Product.NodeDbConn, req.ReportCode)

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

		Query := strings.Replace(reportDataRecord[0].ReportQuery, "$WhereCpmIdValue$", strconv.FormatInt(service.ExeCtx.SessionInfo.Product.CustProdID, 10), 1)

		dberr, _, resultRows := dbaccess.GetReportQueryData(service.ExeCtx.SessionInfo.Product.NodeDbConn, Query, req.QueryParams...)
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

		reportDataModel := hktmodels.DBGetReportDataModel{}
		reportDataModel.ReportCode = reportDataRecord[0].ReportCode

		if req.Language == "en" {
			reportDataModel.ReportHeader = headerModel.En
		} else {
			reportDataModel.ReportHeader = headerModel.Hi
		}

		reportDataModel.ReportData = resultRows

		respList = append(respList, reportDataModel)

	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched report data")

	return true, respList

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
