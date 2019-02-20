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
	router.POST(constants.API_PATIENT_PERSONAL_DETAILS_UPDATE_PERSON_ACCOMPANYING, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_PATIENT_PERSONAL_INFO_MASTER, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_PATIENT_MEDICAL_DETAILS_ADD, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_PATIENT_MEDICAL_DETAILS_UPDATE, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_PATIENT_MEDICAL_DETAILS_UPDATE_PRESENT_COMPLAINTS, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_PATIENT_MEDICAL_DETAILS_UPDATE_REASON_FOR_ADMISSION, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_PATIENT_MEDICAL_DETAILS_UPDATE_HISTORY_PRESENT_ILLNESS, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_PATIENT_MEDICAL_DETAILS_UPDATE_PAST_HISTORY, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_PATIENT_MEDICAL_DETAILS_UPDATE_TREATMENT_BEFORE_ADMISSION, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_PATIENT_MEDICAL_DETAILS_UPDATE_INVESTIGATION_BEFORE_ADMISSION, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_PATIENT_MEDICAL_DETAILS_UPDATE_FAMILY_HISTORY, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_PATIENT_MEDICAL_DETAILS_UPDATE_ALLERGIES, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_PATIENT_MEDICAL_DETAILS_UPDATE_PERSONAL_HISTORY, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_PATIENT_MEDICAL_INFO_MASTER, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_PATIENT_CONFIG_LIST, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_PATIENT_CONFIG_INFO, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_PATIENT_CONFIG_UPDATE, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_PATIENT_LIST_MASTER, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_PATIENT_ADMISSION_INFO_STATUS, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_PATIENT_ADMISSION_INFO_DETAILS, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
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

	case constants.API_PATIENT_PERSONAL_DETAILS_UPDATE_PERSON_ACCOMPANYING:

		reqData := &hktmodels.DBPersonalDetailsUpdatePersonAccompanyingRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.PersonalDetailsUpdatePersonAccompanying(reqData)

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

	case constants.API_PATIENT_MEDICAL_DETAILS_UPDATE_PRESENT_COMPLAINTS:

		reqData := &hktmodels.DBMedicalDetailsUpdatePresentComplaintsRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.MedicalDetailsUpdatePresentComplaints(reqData)

		break

	case constants.API_PATIENT_MEDICAL_DETAILS_UPDATE_REASON_FOR_ADMISSION:

		reqData := &hktmodels.DBMedicalDetailsUpdateReasonForAdmissionRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.MedicalDetailsUpdateReasonForAdmission(reqData)

		break

	case constants.API_PATIENT_MEDICAL_DETAILS_UPDATE_HISTORY_PRESENT_ILLNESS:

		reqData := &hktmodels.DBMedicalDetailsUpdateHistoryPresentIllnessRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.MedicalDetailsUpdateHistoryPresentIllness(reqData)

		break

	case constants.API_PATIENT_MEDICAL_DETAILS_UPDATE_PAST_HISTORY:

		reqData := &hktmodels.DBMedicalDetailsUpdatePastHistoryRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.MedicalDetailsUpdatePastHistory(reqData)

		break

	case constants.API_PATIENT_MEDICAL_DETAILS_UPDATE_TREATMENT_BEFORE_ADMISSION:

		reqData := &hktmodels.DBMedicalDetailsUpdateTreatmentBeforeAdmissionRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.MedicalDetailsUpdateTreatmentBeforeAdmission(reqData)

		break

	case constants.API_PATIENT_MEDICAL_DETAILS_UPDATE_INVESTIGATION_BEFORE_ADMISSION:

		reqData := &hktmodels.DBMedicalDetailsUpdateInvestigationBeforeAdmissionRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.MedicalDetailsUpdateInvestigationBeforeAdmission(reqData)

		break

	case constants.API_PATIENT_MEDICAL_DETAILS_UPDATE_FAMILY_HISTORY:

		reqData := &hktmodels.DBMedicalDetailsUpdateFamilyHistoryRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.MedicalDetailsUpdateFamilyHistory(reqData)

		break

	case constants.API_PATIENT_MEDICAL_DETAILS_UPDATE_ALLERGIES:

		reqData := &hktmodels.DBMedicalDetailsUpdateAllergiesRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.MedicalDetailsUpdateAllergies(reqData)

		break

	case constants.API_PATIENT_MEDICAL_DETAILS_UPDATE_PERSONAL_HISTORY:

		reqData := &hktmodels.DBMedicalDetailsUpdatePersonalHistoryRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.MedicalDetailsUpdatePersonalHistory(reqData)

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

	case constants.API_PATIENT_ADMISSION_INFO_STATUS:

		recReq := gmodels.APIRecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetAdmissionStatusById(recReq.RecId)

		break

	case constants.API_PATIENT_ADMISSION_INFO_DETAILS:

		recReq := gmodels.APIRecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = PatientService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.SelectAdmissionDetailsById(recReq.RecId)

		break

	}

	return isSuccess, resultData
}
