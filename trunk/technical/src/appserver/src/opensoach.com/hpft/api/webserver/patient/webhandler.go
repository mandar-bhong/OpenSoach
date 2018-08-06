package patient

import (
	"github.com/gin-gonic/gin"

	"opensoach.com/core/logger"
	"opensoach.com/hpft/api/constants"
	lhelper "opensoach.com/hpft/api/helper"
	lmodels "opensoach.com/hpft/api/models"
	repo "opensoach.com/hpft/api/repository"
	hktmodels "opensoach.com/hpft/models"
	gmodels "opensoach.com/models"
)

func registerRouters(router *gin.RouterGroup) {
	router.POST(constants.API_PATIENT_ADD, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_PATIENT_LIST, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_PATIENT_UPDATE, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_PATIENT_UPDATE_STATUS, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_PATIENT_INFO_MASTER, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
}

func requestHandler(pContext *gin.Context) (bool, interface{}) {
	var resultData interface{}
	isSuccess := false

	logger.Context().WithField("Request: ", pContext.Request.URL.Path).LogDebug(SUB_MODULE_NAME, logger.Normal, "API request received.")

	switch pContext.Request.URL.Path {

	case constants.API_PATIENT_ADD:

		patientAddReq := lmodels.APIPatientAddRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &patientAddReq)

		if isPrepareExeSuccess == false {
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.Add(patientAddReq)

		break

	case constants.API_PATIENT_LIST:

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionData(repo.Instance().Context, pContext)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetPatientList()

		break

	case constants.API_PATIENT_UPDATE:

		reqData := &hktmodels.DBPatientUpdateRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.Update(reqData)

		break

	case constants.API_PATIENT_UPDATE_STATUS:

		reqData := &hktmodels.DBPatientUpdateStatusRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.UpdateStatus(reqData)

		break

	case constants.API_PATIENT_INFO_MASTER:

		recReq := gmodels.APIRecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.SelectById(recReq.RecId)

		break

	}

	return isSuccess, resultData
}
