package vehicle

import (
	"github.com/gin-gonic/gin"
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	"opensoach.com/vst/api/constants"
	lhelper "opensoach.com/vst/api/helper"
	repo "opensoach.com/vst/api/repository"
)

func registerRouters(router *gin.RouterGroup) {
	router.GET(constants.API_VEHICLE_INFO_MASTER, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
}

func requestHandler(pContext *gin.Context) (bool, interface{}) {
	var resultData interface{}
	isSuccess := false

	logger.Context().WithField("Request: ", pContext.Request.URL.Path).LogDebug(SUB_MODULE_NAME, logger.Normal, "API request received.")

	switch pContext.Request.URL.Path {

	case constants.API_VEHICLE_INFO_MASTER:

		recReq := gmodels.APIRecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = VehicleService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.SelectById(recReq.RecId)

		break

	}

	return isSuccess, resultData
}
