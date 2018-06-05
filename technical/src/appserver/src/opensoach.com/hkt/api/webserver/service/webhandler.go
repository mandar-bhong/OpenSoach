package service

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
	router.POST(constants.API_SERVICE_CONFIG_ADD, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_SERVICE_CONFIG_LIST, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_SERVICE_CONFIG_UPDATE, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_SERVICE_INSTANCE_ADD, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_SERVICE_INSTANCE_LIST, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_SERVICE_TXN_LIST, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_SERVICE_CONFIG_LIST_SHORT, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_SERVICE_CONFIG_COPY_TEMPLATE, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_SERVICE_CONFIG_INFO, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
}

func requestHandler(pContext *gin.Context) (bool, interface{}) {
	var resultData interface{}
	isSuccess := false

	logger.Context().WithField("Request: ", pContext.Request.URL.Path).LogDebug(SUB_MODULE_NAME, logger.Normal, "API request received.")

	switch pContext.Request.URL.Path {

	case constants.API_SERVICE_CONFIG_ADD:

		addReqData := lmodels.APIServiceConfAddRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &addReqData)

		if isPrepareExeSuccess == false {
			return false, successErrorData
		}

		isSuccess, resultData = ServiceConfigService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.ServiceConfigAdd(addReqData)

		break

	case constants.API_SERVICE_CONFIG_LIST:

		listReq := gmodels.APIDataListRequest{}
		listReq.Filter = &hktmodels.DBSearchServiceConfRequestFilterModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &listReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = ServiceConfigService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.ServiceConfigList(listReq)

		break

	case constants.API_SERVICE_CONFIG_UPDATE:

		reqData := &hktmodels.DBServiceConfUpdateRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = ServiceConfigService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.ServiceConnfigUpdate(reqData)

		break

	case constants.API_SERVICE_INSTANCE_ADD:

		addReqData := lmodels.APIServiceInstanceAddRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &addReqData)

		if isPrepareExeSuccess == false {
			return false, successErrorData
		}

		isSuccess, resultData = ServiceConfigService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.ServiceInstanceAdd(addReqData)

		break

	case constants.API_SERVICE_INSTANCE_LIST:

		listReq := gmodels.APIDataListRequest{}
		listReq.Filter = &hktmodels.DBSearchServiceInstanceRequestFilterModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &listReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = ServiceConfigService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.ServiceinstanceList(listReq)

		break

	case constants.API_SERVICE_TXN_LIST:

		req := lmodels.APIServiceInstnaceTxnRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &req)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = ServiceConfigService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetServiceInstanceTxn(req.SPID, req.StartDate, req.EndDate)

		break

	case constants.API_SERVICE_CONFIG_LIST_SHORT:

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionData(repo.Instance().Context, pContext)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = ServiceConfigService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.ServiceConfShortDataList()

		break

	case constants.API_SERVICE_CONFIG_COPY_TEMPLATE:

		reqData := hktmodels.DBServiceConfTemplateInsertDataModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			return false, successErrorData
		}

		isSuccess, resultData = ServiceConfigService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.ServiceConfigCopyTemplate(reqData)

		break

	case constants.API_SERVICE_CONFIG_INFO:

		recReq := gmodels.APIRecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = ServiceConfigService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.ServiceConfInfo(recReq.RecId)

		break

	}

	return isSuccess, resultData
}
