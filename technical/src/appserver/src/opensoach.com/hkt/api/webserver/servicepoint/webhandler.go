package servicepoint

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
	router.POST(constants.API_SERVICE_POINT_CATEGORY_ADD, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_SERVICE_POINT_UPDATE, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_SERVICE_POINT_ASSOCIATE_FOP, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_SERVICE_POINT_FOP_ASSOCIATION_REMOVE, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
}

func requestHandler(pContext *gin.Context) (bool, interface{}) {
	var resultData interface{}
	isSuccess := false

	logger.Context().WithField("Request: ", pContext.Request.URL.Path).LogDebug(SUB_MODULE_NAME, logger.Normal, "API request received.")

	switch pContext.Request.URL.Path {

	case constants.API_SERVICE_POINT_CATEGORY_ADD:

		reqData := lmodels.APISpCategoryAddRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			return false, successErrorData
		}

		isSuccess, resultData = ServicePointService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.SpCategoryAdd(reqData)

		break

	case constants.API_SERVICE_POINT_UPDATE:

		reqData := &hktmodels.DBSpUpdateRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = ServicePointService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.SpUpdate(reqData)

		break

	case constants.API_SERVICE_POINT_ASSOCIATE_FOP:

		reqData := lmodels.APIFopSpAddRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			return false, successErrorData
		}

		isSuccess, resultData = ServicePointService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.FopSpAdd(reqData)

		break

	case constants.API_SERVICE_POINT_FOP_ASSOCIATION_REMOVE:

		reqData := &lmodels.APIFopSpDeleteRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = ServicePointService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.FopSpDelete(reqData)

		break

	}

	return isSuccess, resultData
}
