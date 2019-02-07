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
	router.GET(constants.API_PATIENT_FILTER_INFO_MASTER, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_PATIENT_ADMISSION_ADD, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_PATIENT_ADMISSION_UPDATE, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_PATIENT_ADMISSION_INFO_MASTER, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_PATIENT_PERSONAL_DETAILS_ADD, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_PATIENT_PERSONAL_DETAILS_UPDATE, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_PATIENT_PERSONAL_INFO_MASTER, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_PATIENT_MEDICAL_DETAILS_ADD, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_PATIENT_MEDICAL_DETAILS_UPDATE, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_PATIENT_MEDICAL_INFO_MASTER, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_PATIENT_CONFIG_LIST, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_PATIENT_CONFIG_INFO, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_PATIENT_CONFIG_UPDATE, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_PATIENT_LIST_MASTER, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
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
		}.PatientAdd(patientAddReq)

		break

	case constants.API_PATIENT_LIST:

		listReq := gmodels.APIDataListRequest{}
		listReq.Filter = &hktmodels.DBSearchPatientRequestFilterDataModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &listReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetPatientAdmissionList(listReq)

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
		}.PatientUpdate(reqData)

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
		}.SelectPatientById(recReq.RecId)

		break

	case constants.API_PATIENT_FILTER_INFO_MASTER:

		recReq := &hktmodels.DBPatientFilterModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.SelectPatientByFilter(recReq)

		break

	case constants.API_PATIENT_ADMISSION_ADD:

		addReq := lmodels.APIAdmissionAddRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &addReq)

		if isPrepareExeSuccess == false {
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.AdmissionAdd(addReq)

		break

	case constants.API_PATIENT_ADMISSION_UPDATE:

		reqData := &hktmodels.DBAdmissionTblUpdateRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.AdmissionUpdate(reqData)

		break

	case constants.API_PATIENT_ADMISSION_INFO_MASTER:

		recReq := gmodels.APIRecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.SelectAdmissionById(recReq.RecId)

		break

	case constants.API_PATIENT_PERSONAL_DETAILS_ADD:

		addReq := lmodels.APIPersonalDetailsAddRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &addReq)

		if isPrepareExeSuccess == false {
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.PersonalDetialsAdd(addReq)

		break

	case constants.API_PATIENT_PERSONAL_DETAILS_UPDATE:

		reqData := &hktmodels.DBPersonalDetailsUpdateRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.PersonalDetailsUpdate(reqData)

		break

	case constants.API_PATIENT_PERSONAL_INFO_MASTER:

		recReq := gmodels.APIRecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.SelectPersonalDetailsById(recReq.RecId)

		break

	case constants.API_PATIENT_MEDICAL_DETAILS_ADD:

		addReq := lmodels.APIMedicalDetailsAddRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &addReq)

		if isPrepareExeSuccess == false {
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.MedicalDetialsAdd(addReq)

		break

	case constants.API_PATIENT_MEDICAL_DETAILS_UPDATE:

		reqData := &hktmodels.DBMedicalDetailsUpdateRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.MedicalDetailsUpdate(reqData)

		break

	case constants.API_PATIENT_MEDICAL_INFO_MASTER:

		recReq := gmodels.APIRecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.SelectMedicalDetailsById(recReq.RecId)

		break

	case constants.API_PATIENT_CONFIG_LIST:

		recReq := gmodels.APIRecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.SelectAllPatientConf(recReq.RecId)

		break

	case constants.API_PATIENT_CONFIG_INFO:

		recReq := gmodels.APIRecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.SelectPatientConfById(recReq.RecId)

		break

	case constants.API_PATIENT_CONFIG_UPDATE:

		reqData := &hktmodels.DBPatientConfUpdateRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.PatientConfUpdate(reqData)

		break

	case constants.API_PATIENT_LIST_MASTER:

		listReq := gmodels.APIDataListRequest{}
		listReq.Filter = &hktmodels.DBSearchPatientMasterRequestFilterDataModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &listReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetPatientMasterList(listReq)

		break

	}

	return isSuccess, resultData
}
