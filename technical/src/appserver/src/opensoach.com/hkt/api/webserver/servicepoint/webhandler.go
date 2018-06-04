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
	router.POST(constants.API_SERVICE_POINT_ADD, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_SERVICE_POINT_CATEGORY_LIST_SHORT, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_SERVICE_POINT_ASSOCIATE_DEVICE, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_SERVICE_POINT_DEVICE_ASSOCIATION_REMOVE, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_SERVICE_POINT_LIST, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_SERVICE_POINT_ASSOCIATE_FOP_INFO, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_SERVICE_POINT_LIST_SHORT, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_SERVICE_POINT_INFO, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
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

	case constants.API_SERVICE_POINT_ADD:

		reqData := lmodels.APISpAddRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			return false, successErrorData
		}

		isSuccess, resultData = ServicePointService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.ServicePointAdd(reqData)

		break

	case constants.API_SERVICE_POINT_CATEGORY_LIST_SHORT:

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionData(repo.Instance().Context, pContext)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = ServicePointService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.SpCategoryShortDataList()

		break

	case constants.API_SERVICE_POINT_ASSOCIATE_DEVICE:

		reqData := lmodels.APIDevSpAsscociationRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			return false, successErrorData
		}

		isSuccess, resultData = ServicePointService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.DevSpAssociation(reqData)

		break

	case constants.API_SERVICE_POINT_DEVICE_ASSOCIATION_REMOVE:

		reqData := &lmodels.APIDevSpAsscociationRemoveRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = ServicePointService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.DevSpAsscociationRemove(reqData)

		break

	case constants.API_SERVICE_POINT_LIST:

		listReq := gmodels.APIDataListRequest{}
		listReq.Filter = &hktmodels.DBSearchServicePointRequestFilterDataModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &listReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		listReq.Filter.(*hktmodels.DBSearchServicePointRequestFilterDataModel).CpmId = &successErrorData.(*gmodels.ExecutionContext).SessionInfo.Product.CustProdID

		isSuccess, resultData = ServicePointService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetSPList(listReq)

		break

	case constants.API_SERVICE_POINT_ASSOCIATE_FOP_INFO:

		recReq := gmodels.APIRecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = ServicePointService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetFopSpAssociation(recReq.RecId)

		break

	case constants.API_SERVICE_POINT_LIST_SHORT:

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionData(repo.Instance().Context, pContext)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = ServicePointService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.ServicePointShortDataList()

		break

	case constants.API_SERVICE_POINT_INFO:

		recReq := gmodels.APIRecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = ServicePointService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetServicePointInfo(recReq.RecId)

	}

	return isSuccess, resultData
}
