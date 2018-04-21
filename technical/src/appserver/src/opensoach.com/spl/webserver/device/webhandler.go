package device

import (
	"github.com/gin-gonic/gin"
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	"opensoach.com/spl/constants"
	lhelper "opensoach.com/spl/helper"
	lmodels "opensoach.com/spl/models"
	repo "opensoach.com/spl/repository"
)

func registerRouters(router *gin.RouterGroup) {
	router.POST(constants.API_DEVICE_OSU_ADD, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_DEVICE_OSU_UPDATE_STATE, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_DEVICE_OSU_UPDATE_DETAILS, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_DEVICE_CU_UPDATE_DETAILS, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_DEVICE_OSU_LIST, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_DEVICE_CU_LIST, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
}

func requestHandler(pContext *gin.Context) (bool, interface{}) {
	var resultData interface{}
	isSuccess := false

	logger.Context().WithField("Request: ", pContext.Request.URL.Path).LogDebug(SUB_MODULE_NAME, logger.Normal, "API request received.")

	switch pContext.Request.URL.Path {

	case constants.API_DEVICE_OSU_ADD:

		reqData := &lmodels.DBSplMasterDeviceRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = DeviceService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.AddDevice(reqData)

		break

	case constants.API_DEVICE_OSU_UPDATE_STATE:

		reqData := &lmodels.DevStateRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = DeviceService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.UpdateDevState(reqData)

		break

	case constants.API_DEVICE_OSU_UPDATE_DETAILS:

		detailsReqData := &lmodels.DBSplMasterDevDetailsRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &detailsReqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = DeviceService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.UpdateDeviceDetails(detailsReqData)

		break

	case constants.API_DEVICE_CU_UPDATE_DETAILS:

		detailsReqData := &lmodels.DBSplMasterDevDetailsRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &detailsReqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = DeviceService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.UpdateDeviceDetails(detailsReqData)

		break

	case constants.API_DEVICE_OSU_LIST:

		listReq := lmodels.DataListRequest{}
		listReq.Filter = &lmodels.DBSearchDeviceRequestFilterDataModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &listReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = DeviceService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetDeviceDataList(listReq)

		break

	case constants.API_DEVICE_CU_LIST:

		listReq := lmodels.DataListRequest{}
		listReq.Filter = &lmodels.DBSearchDeviceRequestFilterDataModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &listReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		listReq.Filter.(*lmodels.DBSearchDeviceRequestFilterDataModel).CustId = &successErrorData.(*gmodels.ExecutionContext).SessionInfo.CustomerID

		isSuccess, resultData = DeviceService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetDeviceDataList(listReq)

		break

	}

	return isSuccess, resultData
}
