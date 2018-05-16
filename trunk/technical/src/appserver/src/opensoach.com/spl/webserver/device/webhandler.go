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
	router.GET(constants.API_DEVICE_OSU_LIST, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_DEVICE_CU_LIST, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_DEVICE_OSU_ASSOCIATE_DEV_WITH_CUST, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_DEVICE_OSU_ASSOCIATE_DEV_WITH_CUSTPRODUCT, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_DEVICE_OSU_INFO_DETAILS, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_DEVICE_CU_INFO_DETAILS, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_DEVICE_PRODUCT_ASSCOCIATION_LIST, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
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

		reqData := &lmodels.DBDevStateRowModel{}

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

		userType := "OSU"

		detailsReqData := &lmodels.DBSplMasterDevDetailsRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &detailsReqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = DeviceService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.UpdateDeviceDetails(detailsReqData, userType)

		break

	case constants.API_DEVICE_CU_UPDATE_DETAILS:

		userType := "CU"

		detailsReqData := &lmodels.DBSplMasterDevDetailsRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &detailsReqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = DeviceService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.UpdateDeviceDetails(detailsReqData, userType)

		break

	case constants.API_DEVICE_OSU_LIST:

		listReq := gmodels.APIDataListRequest{}
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

		listReq := gmodels.APIDataListRequest{}
		listReq.Filter = &lmodels.DBSearchDeviceRequestFilterDataModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &listReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		listReq.Filter.(*lmodels.DBSearchDeviceRequestFilterDataModel).CpmId = &successErrorData.(*gmodels.ExecutionContext).SessionInfo.Product.CustProdID

		isSuccess, resultData = DeviceService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetDeviceDataList(listReq)

		break

	case constants.API_DEVICE_OSU_ASSOCIATE_DEV_WITH_CUST:

		reqData := &lmodels.DBDevCustRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = DeviceService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.AssociateDevWithCust(reqData)

		break

	case constants.API_DEVICE_OSU_ASSOCIATE_DEV_WITH_CUSTPRODUCT:

		reqData := &lmodels.DBSplCpmDevRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = DeviceService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.AssociateDevWithCustProduct(reqData)

		break

	case constants.API_DEVICE_OSU_INFO_DETAILS:

		userType := "OSU"

		recReq := gmodels.APIRecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = DeviceService.GetDeviceDetailsInfo(DeviceService{}, recReq.RecId, userType)

		break

	case constants.API_DEVICE_CU_INFO_DETAILS:

		userType := "CU"

		recReq := gmodels.APIRecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = DeviceService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetDeviceDetailsInfo(recReq.RecId, userType)

		break

	case constants.API_DEVICE_PRODUCT_ASSCOCIATION_LIST:

		recReq := gmodels.APIRecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = DeviceService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetDeviceProdAssociation(recReq.RecId)

		break

	}

	return isSuccess, resultData
}
