package master

import (
	"github.com/gin-gonic/gin"
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	"opensoach.com/vst/api/constants"
	lhelper "opensoach.com/vst/api/helper"
	repo "opensoach.com/vst/api/repository"
)

func registerRouters(router *gin.RouterGroup) {
	router.GET(constants.API_MASTER_SPC_TASK_INFO, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_MASTER_TASK_LIB_INFO, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_MASTER_SERV_CONF_TYPE_INFO, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_MASTER_SP_CATEGORY_INFO, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
}

func requestHandler(pContext *gin.Context) (bool, interface{}) {
	var resultData interface{}
	isSuccess := false

	logger.Context().WithField("Request: ", pContext.Request.URL.Path).LogDebug(SUB_MODULE_NAME, logger.Normal, "API request received.")

	switch pContext.Request.URL.Path {

	case constants.API_MASTER_SPC_TASK_INFO:

		recReq := gmodels.APIRecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = MasterService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetSpcTaskLib(recReq.RecId)

		break

	case constants.API_MASTER_TASK_LIB_INFO:

		recReq := gmodels.APIRecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = MasterService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetTaskLib(recReq.RecId)

		break

	case constants.API_MASTER_SERV_CONF_TYPE_INFO:

		recReq := gmodels.APIRecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = MasterService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetServConfType(recReq.RecId)

		break

	case constants.API_MASTER_SP_CATEGORY_INFO:

		recReq := gmodels.APIRecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = MasterService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetSpCategory(recReq.RecId)

		break
	}

	return isSuccess, resultData
}
