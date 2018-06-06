package dashboard

import (
	"github.com/gin-gonic/gin"

	"opensoach.com/core/logger"
	"opensoach.com/hkt/api/constants"
	lhelper "opensoach.com/hkt/api/helper"
	lmodels "opensoach.com/hkt/api/models"
	repo "opensoach.com/hkt/api/repository"
	gmodels "opensoach.com/models"
)

func registerRouters(router *gin.RouterGroup) {
	router.GET(constants.API_DASHBOARD_DEVICE_SUMMARY, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_DASHBOARD_LOCATION_SUMMARY, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_DASHBOARD_FEEDBACK_SUMMARY, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
}

func requestHandler(pContext *gin.Context) (bool, interface{}) {
	var resultData interface{}
	isSuccess := false

	logger.Context().WithField("Request: ", pContext.Request.URL.Path).LogDebug(SUB_MODULE_NAME, logger.Normal, "API request received.")

	switch pContext.Request.URL.Path {

	case constants.API_DASHBOARD_DEVICE_SUMMARY:

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionData(repo.Instance().Context, pContext)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = DashboardService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetDeviceSummary()

		break

	case constants.API_DASHBOARD_LOCATION_SUMMARY:
		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionData(repo.Instance().Context, pContext)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = DashboardService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetLocationSummary()
		break

	case constants.API_DASHBOARD_FEEDBACK_SUMMARY:

		feedbackReq := lmodels.APIDashboardFeedbackRequest{}
		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &feedbackReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = DashboardService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetFeedbackSummary(feedbackReq)
		break
	}

	return isSuccess, resultData
}
