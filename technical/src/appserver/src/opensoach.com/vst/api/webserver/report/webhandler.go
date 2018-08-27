package report

import (
	"github.com/gin-gonic/gin"
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	"opensoach.com/vst/api/constants"
	lhelper "opensoach.com/vst/api/helper"
	lmodels "opensoach.com/vst/api/models"
	repo "opensoach.com/vst/api/repository"
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

		generateReportRequest := lmodels.APIGenerateReportRequestModel{}

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
