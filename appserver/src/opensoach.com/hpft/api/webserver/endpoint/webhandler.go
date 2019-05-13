package endpoint

import (
	"github.com/gin-gonic/gin"
	"opensoach.com/core/logger"
	"opensoach.com/hpft/api/constants"
	lhelper "opensoach.com/hpft/api/helper"
	lmodels "opensoach.com/hpft/api/models"
	repo "opensoach.com/hpft/api/repository"
	hpftmodels "opensoach.com/hpft/models"
	gmodels "opensoach.com/models"
)

func registerRouters(router *gin.RouterGroup) {
	router.GET(constants.API_ENDPOINT_PATIENT_LIST, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_ENDPOINT_USER_PATIENT_ASSOCIATION, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_ENDPOINT_USER_PATIENT_ASSOCIATION_REMOVE, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
}

func requestHandler(pContext *gin.Context) (bool, interface{}) {
	var resultData interface{}
	isSuccess := false

	logger.Context().WithField("Request: ", pContext.Request.URL.Path).LogDebug(SUB_MODULE_NAME, logger.Normal, "API request received.")

	switch pContext.Request.URL.Path {

	case constants.API_ENDPOINT_PATIENT_LIST:

		listReq := gmodels.APIDataListRequest{}
		listReq.Filter = &hpftmodels.DBDeviceSearchPatientRequestFilterDataModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareDeviceUserExecutionReqData(repo.Instance().Context, pContext, &listReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = EndpointService{
			ExeCtx: successErrorData.(*gmodels.DeviceUserExecutionContext),
		}.GetPatientAdmissionList(listReq)

		break

	case constants.API_ENDPOINT_USER_PATIENT_ASSOCIATION:

		reqData := lmodels.APIUserPatientAsscociationRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareDeviceUserExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			return false, successErrorData
		}

		isSuccess, resultData = EndpointService{
			ExeCtx: successErrorData.(*gmodels.DeviceUserExecutionContext),
		}.UserPateintAssociate(reqData)

		break

	case constants.API_ENDPOINT_USER_PATIENT_ASSOCIATION_REMOVE:

		reqData := &hpftmodels.DBPatientMonitorMappingDeleteRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareDeviceUserExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = EndpointService{
			ExeCtx: successErrorData.(*gmodels.DeviceUserExecutionContext),
		}.UserPatientAsscociationRemove(reqData)

		break

	}

	return isSuccess, resultData
}
