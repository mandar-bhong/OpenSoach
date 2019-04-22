package task

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
	router.POST(constants.API_TASK_ADD, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_TASK_UPDATE, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_TASK_INFO_MASTER, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_TASK_LIST, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
}

func requestHandler(pContext *gin.Context) (bool, interface{}) {
	var resultData interface{}
	isSuccess := false

	logger.Context().WithField("Request: ", pContext.Request.URL.Path).LogDebug(SUB_MODULE_NAME, logger.Normal, "API request received.")

	switch pContext.Request.URL.Path {

	case constants.API_TASK_ADD:
		taskAddReq := lmodels.APITaskAddRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &taskAddReq)

		if isPrepareExeSuccess == false {
			return false, successErrorData
		}

		isSuccess, resultData = TaskService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.Add(taskAddReq)

		break

	case constants.API_TASK_INFO_MASTER:

		recReq := gmodels.APIRecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = TaskService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.SelectById(recReq.RecId)

		break

	case constants.API_TASK_UPDATE:

		reqData := &hktmodels.DBTaskLibUpdateRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = TaskService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.Update(reqData)

		break

	case constants.API_TASK_LIST:

		recReq := gmodels.APIRecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = TaskService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.SelectBySPCId(recReq.RecId)

		break

	}

	return isSuccess, resultData
}
