package report

import (
	"github.com/gin-gonic/gin"
	"opensoach.com/core/logger"
	"opensoach.com/hkt/api/constants"
	lhelper "opensoach.com/hkt/api/helper"
	lmodels "opensoach.com/hkt/api/models"
	repo "opensoach.com/hkt/api/repository"
	hktmodels "opensoach.com/hkt/models"
	gmodels "opensoach.com/models"
)

func registerRouters(router *gin.RouterGroup) {
	router.GET(constants.API_REPORT_GENERATE, func(c *gin.Context) { lhelper.FileDownloadHandler(c, requestHandler) })
	router.GET(constants.API_REPORT_VIEW, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_REPORT_LIST_SHORT, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
}

func requestHandler(pContext *gin.Context) (bool, interface{}) {
	var resultData interface{}
	isSuccess := false

	logger.Context().WithField("Request: ", pContext.Request.URL.Path).LogDebug(SUB_MODULE_NAME, logger.Normal, "API request received.")

	switch pContext.Request.URL.Path {

	case constants.API_REPORT_GENERATE:

		generateReportRequest := hktmodels.DBReportRequestDataModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &generateReportRequest)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = ReportService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GenerateReport(generateReportRequest)

		break

	case constants.API_REPORT_VIEW:

		viewReportRequest := lmodels.APIViewReportRequestModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &viewReportRequest)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = ReportService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.ViewReport(viewReportRequest)

		break

	case constants.API_REPORT_LIST_SHORT:

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionData(repo.Instance().Context, pContext)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = ReportService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.ReportShortList()

		break

	}

	return isSuccess, resultData
}
