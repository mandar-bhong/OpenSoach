package patient

import (
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	lmodels "opensoach.com/hpft/api/models"
	repo "opensoach.com/hpft/api/repository"
	"opensoach.com/hpft/api/webserver/patient/dbaccess"
	"opensoach.com/hpft/constants"
	hktmodels "opensoach.com/hpft/models"
	gmodels "opensoach.com/models"
)

var SUB_MODULE_NAME = "HPFT.API.Patient"

type PatientService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service PatientService) PatientAdd(req lmodels.APIPatientAddRequest) (isSuccess bool, successErrorData interface{}) {

	dbRowModel := &hktmodels.DBPatientMasterInsertRowModel{}
	dbRowModel.DBPatientMasterDataModel = req.DBPatientMasterDataModel
	dbRowModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	dbRowModel.UpdatedBy = service.ExeCtx.SessionInfo.UserID
	dbRowModel.Uuid = ghelper.GenerateUUID()

	dbErr, insertedId := dbaccess.Insert(service.ExeCtx.SessionInfo.Product.NodeDbConn, dbRowModel)
	if dbErr != nil {

		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while adding new patient.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	addResponse := gmodels.APIRecordAddResponse{}
	addResponse.RecordID = insertedId

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "New patient added succesfully")

	return true, addResponse
}

func (service PatientService) GetPatientAdmissionList(listReqData gmodels.APIDataListRequest) (bool, interface{}) {

	dataListResponse := gmodels.APIDataListResponse{}

	filterModel := listReqData.Filter.(*hktmodels.DBSearchPatientRequestFilterDataModel)
	filterModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	CurrentPage := listReqData.CurrentPage
	startingRecord := ((CurrentPage - 1) * listReqData.Limit)

	dbErr, listData := dbaccess.GetPatientList(service.ExeCtx.SessionInfo.Product.NodeDbConn, filterModel, listReqData, startingRecord)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting patient admission list.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbListDataRecord := *listData

	dataListResponse.FilteredRecords = dbListDataRecord.RecordCount
	dataListResponse.Records = dbListDataRecord.RecordList

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched patient admission list data.")

	return true, dataListResponse

}

func (service PatientService) PatientUpdate(reqData *hktmodels.DBPatientUpdateRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	reqData.UpdatedBy = service.ExeCtx.SessionInfo.UserID

	dbErr, _ := dbaccess.UpdateByFilter(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating patient info.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient data updated successfully.")

	return true, nil
}

func (service PatientService) UpdateStatus(reqData *hktmodels.DBPatientUpdateStatusRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID

	dbErr, _ := dbaccess.UpdatePatientStatus(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while udating patient status.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	// remove doctor patient association
	dbErr, admsnData := dbaccess.GetAdmissionById(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData.AdmissionId)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting admission info by id.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	admsnDataDBRecord := *admsnData
	deassociateReq := &hktmodels.DBPatientMonitorMappingDeleteRowModel{}
	deassociateReq.UsrId = admsnDataDBRecord[0].DrIncharge
	deassociateReq.SpId = &admsnDataDBRecord[0].SpId
	deassociateReq.PatientId = &admsnDataDBRecord[0].PatientId
	service.UserPatientAsscociationRemove(deassociateReq)

	// taskPatientStatusUpdated := &hktmodels.TaskPatientStatusUpdated{}
	// taskPatientStatusUpdated.PatientId = reqData.PatientId
	// taskPatientStatusUpdated.CpmId = reqData.CpmId

	// isSendSuccess := repo.Instance().
	// 	SendTaskToServer(hktconst.TASK_HKT_API_PATIENT_STATUS_UPDATED,
	// 		service.ExeCtx.SessionToken, taskPatientStatusUpdated)

	// if isSendSuccess == false {
	// 	logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Unable to submit task to server.", nil)
	// }

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "patient status updated successfully.")

	return true, nil
}

func (service PatientService) SelectPatientById(patientID int64) (bool, interface{}) {

	dbErr, data := dbaccess.GetPatientById(service.ExeCtx.SessionInfo.Product.NodeDbConn, patientID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting patient info by id.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *data

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched patient info")
	return true, dbRecord[0]
}

func (service PatientService) SelectPatientByFilter(req *hktmodels.DBPatientFilterModel) (bool, interface{}) {

	dbErr, data := dbaccess.GetPatientByFilter(service.ExeCtx.SessionInfo.Product.NodeDbConn, req)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting patient info by id.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *data

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched patient info")
	return true, dbRecord
}

func (service PatientService) AdmissionAdd(req lmodels.APIAdmissionAddRequest) (isSuccess bool, successErrorData interface{}) {

	dbRowModel := &hktmodels.DBAdmissionTblInsertRowModel{}
	dbRowModel.DBAdmissionTblDataModel = req.DBAdmissionTblDataModel
	dbRowModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	dbRowModel.UpdatedBy = service.ExeCtx.SessionInfo.UserID
	dbRowModel.Uuid = ghelper.GenerateUUID()

	dbErr, insertedId := dbaccess.AdmissionTblInsert(service.ExeCtx.SessionInfo.Product.NodeDbConn, dbRowModel)
	if dbErr != nil {

		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while adding new patient admission info.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	addResponse := lmodels.APIAdmissionAddResponse{}
	addResponse.AdmissionId = insertedId

	//get patient master data
	dbErr, data := dbaccess.GetPatientById(service.ExeCtx.SessionInfo.Product.NodeDbConn, req.PatientId)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting patient info by id.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *data

	// insert personal details
	personalDetailsAddRequest := lmodels.APIPersonalDetailsAddRequest{}
	personalDetailsAddRequest.PatientId = req.PatientId
	personalDetailsAddRequest.AdmissionId = insertedId
	personalDetailsAddRequest.Age = dbRecord[0].Age
	personalDetailsAddRequest.Uuid = ghelper.GenerateUUID()

	isSuccess, personalDetailsAddResponse := service.PersonalDetialsAdd(personalDetailsAddRequest)
	if isSuccess == false {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while adding personal details info.", dbErr)
	}
	addResponse.PersonalDetailsId = personalDetailsAddResponse.(gmodels.APIRecordAddResponse).RecordID

	// insert medical details
	medicalDetailsAddRequest := lmodels.APIMedicalDetailsAddRequest{}
	medicalDetailsAddRequest.PatientId = req.PatientId
	medicalDetailsAddRequest.AdmissionId = insertedId
	medicalDetailsAddRequest.Uuid = ghelper.GenerateUUID()

	isSuccess, medicalDetailsAddResponse := service.MedicalDetialsAdd(medicalDetailsAddRequest)
	if isSuccess == false {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while adding medical details info.", dbErr)
	}

	addResponse.MedicalDetailsId = medicalDetailsAddResponse.(gmodels.APIRecordAddResponse).RecordID

	// associate patient to doctor
	dbErr, upmmData := dbaccess.GetUserPatientassociationByUsrIdSpId(service.ExeCtx.SessionInfo.Product.NodeDbConn, req.DrIncharge, req.SpId)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting patient info by id.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	upmmDbRecord := *upmmData

	if len(upmmDbRecord) == 0 || len(upmmDbRecord) > 1 {
		associateReq := lmodels.APIUserPatientAsscociationRequest{}
		associateReq.UsrId = req.DrIncharge
		associateReq.SpId = &req.SpId
		associateReq.PatientId = &req.PatientId
		service.UserPateintAssociate(associateReq)
	} else if len(upmmDbRecord) == 1 {
		if upmmDbRecord[0].PatientId != nil {
			associateReq := lmodels.APIUserPatientAsscociationRequest{}
			associateReq.UsrId = req.DrIncharge
			*associateReq.SpId = req.SpId
			*associateReq.PatientId = req.PatientId
			service.UserPateintAssociate(associateReq)
		}
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient admission info added succesfully")

	return true, addResponse
}

func (service PatientService) AdmissionUpdate(reqData *hktmodels.DBAdmissionTblUpdateRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	reqData.UpdatedBy = service.ExeCtx.SessionInfo.UserID

	dbErr, _ := dbaccess.AdmissionTblUpdateByFilter(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating patient admission info.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient admission data updated successfully.")

	return true, nil
}

func (service PatientService) SelectAdmissionById(admissionID int64) (bool, interface{}) {

	dbErr, data := dbaccess.GetAdmissionById(service.ExeCtx.SessionInfo.Product.NodeDbConn, admissionID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting admission info by id.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *data

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched patient admission info")
	return true, dbRecord[0]
}

func (service PatientService) PersonalDetialsAdd(req lmodels.APIPersonalDetailsAddRequest) (isSuccess bool, successErrorData interface{}) {

	dbRowModel := &hktmodels.DBPersonalDetailsInsertRowModel{}
	dbRowModel.DBPersonalDetailsDataModel = req.DBPersonalDetailsDataModel
	dbRowModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	dbRowModel.UpdatedBy = service.ExeCtx.SessionInfo.UserID
	dbRowModel.Uuid = ghelper.GenerateUUID()

	dbErr, insertedId := dbaccess.PersonalDetailsInsert(service.ExeCtx.SessionInfo.Product.NodeDbConn, dbRowModel)
	if dbErr != nil {

		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while adding new patient personal details info.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	addResponse := gmodels.APIRecordAddResponse{}
	addResponse.RecordID = insertedId

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient admission info added succesfully")

	return true, addResponse
}

func (service PatientService) PersonalDetailsUpdate(reqData *hktmodels.DBPersonalDetailsUpdateRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	reqData.UpdatedBy = service.ExeCtx.SessionInfo.UserID

	dbErr, _ := dbaccess.PersonalDetailsUpdateByFilter(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating patient personal details.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient personal details updated successfully.")

	return true, nil
}

func (service PatientService) PersonalDetailsUpdatePersonAccompanying(reqData *hktmodels.DBPersonalDetailsUpdatePersonAccompanyingRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	reqData.UpdatedBy = service.ExeCtx.SessionInfo.UserID

	dbErr, _ := dbaccess.PersonalDetailsUpdatePersonAccompanying(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating patient personal details.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient person accompanying details updated successfully.")

	return true, nil
}

func (service PatientService) SelectPersonalDetailsById(personalDetailsID int64) (bool, interface{}) {

	dbErr, data := dbaccess.GetPersonalDetailsById(service.ExeCtx.SessionInfo.Product.NodeDbConn, personalDetailsID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting patient personal details by id.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *data

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched patient personal details.")
	return true, dbRecord[0]
}

func (service PatientService) MedicalDetialsAdd(req lmodels.APIMedicalDetailsAddRequest) (isSuccess bool, successErrorData interface{}) {

	dbRowModel := &hktmodels.DBMedicalDetailsInsertRowModel{}
	dbRowModel.DBMedicalDetailsDataModel = req.DBMedicalDetailsDataModel
	dbRowModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	dbRowModel.UpdatedBy = service.ExeCtx.SessionInfo.UserID
	dbRowModel.Uuid = ghelper.GenerateUUID()

	dbErr, insertedId := dbaccess.MedicalDetailsInsert(service.ExeCtx.SessionInfo.Product.NodeDbConn, dbRowModel)
	if dbErr != nil {

		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while adding new patient medical details.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	addResponse := gmodels.APIRecordAddResponse{}
	addResponse.RecordID = insertedId

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient medical details added succesfully")

	return true, addResponse
}

func (service PatientService) MedicalDetailsUpdate(reqData *hktmodels.DBMedicalDetailsUpdateRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	reqData.UpdatedBy = service.ExeCtx.SessionInfo.UserID

	dbErr, _ := dbaccess.MedicalDetailsUpdateByFilter(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating patient medical details.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient medical details updated successfully.")

	return true, nil
}

func (service PatientService) MedicalDetailsUpdatePresentComplaints(reqData *hktmodels.DBMedicalDetailsUpdatePresentComplaintsRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	reqData.UpdatedBy = service.ExeCtx.SessionInfo.UserID

	dbErr, _ := dbaccess.MedicalDetailsUpdatePresentComplaints(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating patient medical details.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient present complaints details updated successfully.")

	return true, nil
}

func (service PatientService) MedicalDetailsUpdateReasonForAdmission(reqData *hktmodels.DBMedicalDetailsUpdateReasonForAdmissionRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	reqData.UpdatedBy = service.ExeCtx.SessionInfo.UserID

	dbErr, _ := dbaccess.MedicalDetailsUpdateReasonForAdmission(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating patient medical details.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient reason for admission details updated successfully.")

	return true, nil
}

func (service PatientService) MedicalDetailsUpdateHistoryPresentIllness(reqData *hktmodels.DBMedicalDetailsUpdateHistoryPresentIllnessRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	reqData.UpdatedBy = service.ExeCtx.SessionInfo.UserID

	dbErr, _ := dbaccess.MedicalDetailsUpdateHistoryPresentIllness(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating patient medical details.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient history present illness details updated successfully.")

	return true, nil
}

func (service PatientService) MedicalDetailsUpdatePastHistory(reqData *hktmodels.DBMedicalDetailsUpdatePastHistoryRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	reqData.UpdatedBy = service.ExeCtx.SessionInfo.UserID

	dbErr, _ := dbaccess.MedicalDetailsUpdatePastHistory(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating patient medical details.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient past history details updated successfully.")

	return true, nil
}

func (service PatientService) MedicalDetailsUpdateTreatmentBeforeAdmission(reqData *hktmodels.DBMedicalDetailsUpdateTreatmentBeforeAdmissionRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	reqData.UpdatedBy = service.ExeCtx.SessionInfo.UserID

	dbErr, _ := dbaccess.MedicalDetailsUpdateTreatmentBeforeAdmission(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating patient medical details.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient treatment before admission details updated successfully.")

	return true, nil
}

func (service PatientService) MedicalDetailsUpdateInvestigationBeforeAdmission(reqData *hktmodels.DBMedicalDetailsUpdateInvestigationBeforeAdmissionRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	reqData.UpdatedBy = service.ExeCtx.SessionInfo.UserID

	dbErr, _ := dbaccess.MedicalDetailsUpdateInvestigationBeforeAdmission(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating patient medical details.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient Investigation details before admission details updated successfully.")

	return true, nil
}

func (service PatientService) MedicalDetailsUpdateFamilyHistory(reqData *hktmodels.DBMedicalDetailsUpdateFamilyHistoryRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	reqData.UpdatedBy = service.ExeCtx.SessionInfo.UserID

	dbErr, _ := dbaccess.MedicalDetailsUpdateFamilyHistory(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating patient medical details.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient family history details updated successfully.")

	return true, nil
}

func (service PatientService) MedicalDetailsUpdateAllergies(reqData *hktmodels.DBMedicalDetailsUpdateAllergiesRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	reqData.UpdatedBy = service.ExeCtx.SessionInfo.UserID

	dbErr, _ := dbaccess.MedicalDetailsUpdateAllergies(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating patient medical details.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient allegies details updated successfully.")

	return true, nil
}

func (service PatientService) MedicalDetailsUpdatePersonalHistory(reqData *hktmodels.DBMedicalDetailsUpdatePersonalHistoryRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	reqData.UpdatedBy = service.ExeCtx.SessionInfo.UserID

	dbErr, _ := dbaccess.MedicalDetailsUpdatePersonalHistory(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating patient medical details.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient personal history details updated successfully.")

	return true, nil
}

func (service PatientService) SelectMedicalDetailsById(medicalDetailsID int64) (bool, interface{}) {

	dbErr, data := dbaccess.GetMedicalDetailsById(service.ExeCtx.SessionInfo.Product.NodeDbConn, medicalDetailsID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting patient medical details by id.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *data

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched patient medical details.")
	return true, dbRecord[0]
}

func (service PatientService) GetPatientConfigList(listReqData gmodels.APIDataListRequest) (bool, interface{}) {

	dataListResponse := gmodels.APIDataListResponse{}

	filterModel := listReqData.Filter.(*hktmodels.DBSearchPatientConfRequestFilterDataModel)
	filterModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	CurrentPage := listReqData.CurrentPage
	startingRecord := ((CurrentPage - 1) * listReqData.Limit)

	dbErr, listData := dbaccess.GetPatientConfList(service.ExeCtx.SessionInfo.Product.NodeDbConn, filterModel, listReqData, startingRecord)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting patient config list.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbListDataRecord := *listData

	dataListResponse.FilteredRecords = dbListDataRecord.RecordCount
	dataListResponse.Records = dbListDataRecord.RecordList

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched patient config list data.")

	return true, dataListResponse

}

func (service PatientService) SelectPatientConfById(confID int64) (bool, interface{}) {

	dbErr, data := dbaccess.GetPatientConfById(service.ExeCtx.SessionInfo.Product.NodeDbConn, confID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting patient config by id.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *data

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched patient config.")
	return true, dbRecord[0]
}

func (service PatientService) PatientConfUpdate(reqData *hktmodels.DBPatientConfUpdateRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID

	dbErr, _ := dbaccess.PatientConfUpdateByFilter(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating patient config.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient config updated successfully.")

	return true, nil
}

func (service PatientService) GetPatientMasterList(listReqData gmodels.APIDataListRequest) (bool, interface{}) {

	dataListResponse := gmodels.APIDataListResponse{}

	filterModel := listReqData.Filter.(*hktmodels.DBSearchPatientMasterRequestFilterDataModel)
	filterModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	CurrentPage := listReqData.CurrentPage
	startingRecord := ((CurrentPage - 1) * listReqData.Limit)

	dbErr, listData := dbaccess.GetPatientMasterList(service.ExeCtx.SessionInfo.Product.NodeDbConn, filterModel, listReqData, startingRecord)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting patient master list.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbListDataRecord := *listData

	dataListResponse.FilteredRecords = dbListDataRecord.RecordCount
	dataListResponse.Records = dbListDataRecord.RecordList

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched patient master list data.")

	return true, dataListResponse

}

func (service PatientService) GetAdmissionStatusById(admissionID int64) (bool, interface{}) {

	dbErr, data := dbaccess.GetAdmissionStatusById(service.ExeCtx.SessionInfo.Product.NodeDbConn, admissionID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting admission info by id.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *data

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched patient admission status info")
	return true, dbRecord[0]
}

func (service PatientService) SelectAdmissionDetailsById(admissionID int64) (bool, interface{}) {

	dbErr, data := dbaccess.GetPersonalDetailsByAdmissionId(service.ExeCtx.SessionInfo.Product.NodeDbConn, admissionID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting personal details info by admissionid.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	personalDetailsData := *data

	dbErr, dbdata := dbaccess.GetMedicalDetailsDetailsByAdmissionId(service.ExeCtx.SessionInfo.Product.NodeDbConn, admissionID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting medical details info by admissionid.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	medicalDetailsData := *dbdata

	admissionDetailsResponse := lmodels.APIAdmissionDetailsResponse{}
	admissionDetailsResponse.PersonalDetails = personalDetailsData[0]
	admissionDetailsResponse.MedicalDetails = medicalDetailsData[0]

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched patient admission details info")
	return true, admissionDetailsResponse
}

func (service PatientService) GetPatientActionTxnList(listReqData gmodels.APIDataListRequest) (bool, interface{}) {

	dataListResponse := gmodels.APIDataListResponse{}

	filterModel := listReqData.Filter.(*hktmodels.DBSearchPatientActionTxnRequestFilterDataModel)
	filterModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	CurrentPage := listReqData.CurrentPage
	startingRecord := ((CurrentPage - 1) * listReqData.Limit)

	dbErr, listData := dbaccess.GetPatientActionTxnList(service.ExeCtx.SessionInfo.Product.NodeDbConn, filterModel, listReqData, startingRecord)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting patient action txn list.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbListDataRecord := *listData
	dbListDataRecord.RecordList = *dbListDataRecord.RecordList.(*[]hktmodels.DBSearchPatientActionTxnResponseFilterDataModel)

	for i, each := range dbListDataRecord.RecordList.([]hktmodels.DBSearchPatientActionTxnResponseFilterDataModel) {
		dbErr, userdata := dbaccess.GetUserInfoById(repo.Instance().Context.Master.DBConn, each.UpdatedBy)
		if dbErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting patient info by id.", dbErr)

			errModel := gmodels.APIResponseError{}
			errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
			return false, errModel
		}

		userRecord := *userdata

		dbListDataRecord.RecordList.([]hktmodels.DBSearchPatientActionTxnResponseFilterDataModel)[i].FirstName = userRecord[0].Firstname
		dbListDataRecord.RecordList.([]hktmodels.DBSearchPatientActionTxnResponseFilterDataModel)[i].LastName = userRecord[0].LastName
	}

	dataListResponse.FilteredRecords = dbListDataRecord.RecordCount
	dataListResponse.Records = dbListDataRecord.RecordList

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched patient action txn list data.")

	return true, dataListResponse

}

func (service PatientService) SelectPatientDoctorsOrdersById(doctorsordersID int64) (bool, interface{}) {

	dbErr, data := dbaccess.GetPatientDoctorsOrdersById(service.ExeCtx.SessionInfo.Product.NodeDbConn, doctorsordersID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting patient doctord orders info by id.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *data

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched patient doctor orders info.")
	return true, dbRecord[0]
}

func (service PatientService) SelectPatientPathologyRecordsById(pathologyrecordID int64) (bool, interface{}) {

	dbErr, data := dbaccess.GetPatientPathologyRecordsById(service.ExeCtx.SessionInfo.Product.NodeDbConn, pathologyrecordID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting patient pathology record info by id.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *data

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched patient pathology record info.")
	return true, dbRecord[0]
}

func (service PatientService) SelectPatientTreatmentById(pathologyrecordID int64) (bool, interface{}) {

	dbErr, data := dbaccess.GetPatientTreatmentById(service.ExeCtx.SessionInfo.Product.NodeDbConn, pathologyrecordID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting patient treatment info by id.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *data

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched patient treatment info.")
	return true, dbRecord[0]
}

func (service PatientService) GetPatientDoctorOrdersList(listReqData gmodels.APIDataListRequest) (bool, interface{}) {

	dataListResponse := gmodels.APIDataListResponse{}

	filterModel := listReqData.Filter.(*hktmodels.DBSearchPatientDoctorOrdersRequestFilterDataModel)
	filterModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	CurrentPage := listReqData.CurrentPage
	startingRecord := ((CurrentPage - 1) * listReqData.Limit)

	dbErr, listData := dbaccess.GetPatientDoctorsOrdersList(service.ExeCtx.SessionInfo.Product.NodeDbConn, filterModel, listReqData, startingRecord)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting patient doctor orders list.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbListDataRecord := *listData
	dbListDataRecord.RecordList = *dbListDataRecord.RecordList.(*[]hktmodels.DBSearchPatientDoctorOrdersResponseFilterDataModel)

	for i, each := range dbListDataRecord.RecordList.([]hktmodels.DBSearchPatientDoctorOrdersResponseFilterDataModel) {

		dbErr, userdata := dbaccess.GetUserInfoById(repo.Instance().Context.Master.DBConn, each.DoctorId)
		if dbErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting doctor info by id.", dbErr)

			errModel := gmodels.APIResponseError{}
			errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
			return false, errModel
		}

		userRecord := *userdata

		dbListDataRecord.RecordList.([]hktmodels.DBSearchPatientDoctorOrdersResponseFilterDataModel)[i].DoctorFirstName = &userRecord[0].Firstname
		dbListDataRecord.RecordList.([]hktmodels.DBSearchPatientDoctorOrdersResponseFilterDataModel)[i].DoctorLastName = &userRecord[0].LastName

	}

	for i, each := range dbListDataRecord.RecordList.([]hktmodels.DBSearchPatientDoctorOrdersResponseFilterDataModel) {
		if each.AckBy != nil {
			dbErr, userdata := dbaccess.GetUserInfoById(repo.Instance().Context.Master.DBConn, *each.AckBy)
			if dbErr != nil {
				logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting ack by info by id.", dbErr)

				errModel := gmodels.APIResponseError{}
				errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
				return false, errModel
			}

			userRecord := *userdata

			dbListDataRecord.RecordList.([]hktmodels.DBSearchPatientDoctorOrdersResponseFilterDataModel)[i].AckByFirstName = &userRecord[0].Firstname
			dbListDataRecord.RecordList.([]hktmodels.DBSearchPatientDoctorOrdersResponseFilterDataModel)[i].AckByLastName = &userRecord[0].LastName
		}
	}

	dataListResponse.FilteredRecords = dbListDataRecord.RecordCount
	dataListResponse.Records = dbListDataRecord.RecordList

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched patient doctor orders list data.")

	return true, dataListResponse

}

func (service PatientService) GetPatientTreatmentList(listReqData gmodels.APIDataListRequest) (bool, interface{}) {

	dataListResponse := gmodels.APIDataListResponse{}

	filterModel := listReqData.Filter.(*hktmodels.DBSearchPatientTreatmentRequestFilterDataModel)
	filterModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	CurrentPage := listReqData.CurrentPage
	startingRecord := ((CurrentPage - 1) * listReqData.Limit)

	dbErr, listData := dbaccess.GetPatientTreatmentsList(service.ExeCtx.SessionInfo.Product.NodeDbConn, filterModel, listReqData, startingRecord)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting patient treatment data list.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbListDataRecord := *listData
	dbListDataRecord.RecordList = *dbListDataRecord.RecordList.(*[]hktmodels.DBSearchPatientTreatmentResponseFilterDataModel)

	for i, each := range dbListDataRecord.RecordList.([]hktmodels.DBSearchPatientTreatmentResponseFilterDataModel) {

		dbErr, data := dbaccess.GetPatientTreatmentDocumentsById(service.ExeCtx.SessionInfo.Product.NodeDbConn, each.TreatmentId)
		if dbErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting treatment documents by id.", dbErr)

			errModel := gmodels.APIResponseError{}
			errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
			return false, errModel
		}

		dbRecord := *data

		for j := 0; j < len(dbRecord); j++ {
			dbListDataRecord.RecordList.([]hktmodels.DBSearchPatientTreatmentResponseFilterDataModel)[i].DocumentList = append(dbListDataRecord.RecordList.([]hktmodels.DBSearchPatientTreatmentResponseFilterDataModel)[i].DocumentList, dbRecord[j])
		}

	}

	dataListResponse.FilteredRecords = dbListDataRecord.RecordCount
	dataListResponse.Records = dbListDataRecord.RecordList

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched patient treatment list data.")

	return true, dataListResponse

}

func (service PatientService) GetPatientPathologyRecordsList(listReqData gmodels.APIDataListRequest) (bool, interface{}) {

	dataListResponse := gmodels.APIDataListResponse{}

	filterModel := listReqData.Filter.(*hktmodels.DBSearchPatientPathologyRecordRequestFilterDataModel)
	filterModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	CurrentPage := listReqData.CurrentPage
	startingRecord := ((CurrentPage - 1) * listReqData.Limit)

	dbErr, listData := dbaccess.GetPatientPathologyRecordList(service.ExeCtx.SessionInfo.Product.NodeDbConn, filterModel, listReqData, startingRecord)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting patient pathology records list.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbListDataRecord := *listData
	dbListDataRecord.RecordList = *dbListDataRecord.RecordList.(*[]hktmodels.DBSearchPatientPathologyRecordResponseFilterDataModel)

	for i, each := range dbListDataRecord.RecordList.([]hktmodels.DBSearchPatientPathologyRecordResponseFilterDataModel) {

		dbErr, data := dbaccess.GetPatientPathologyRecordsDocumentsById(service.ExeCtx.SessionInfo.Product.NodeDbConn, each.PathologyId)
		if dbErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting treatment documents by id.", dbErr)

			errModel := gmodels.APIResponseError{}
			errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
			return false, errModel
		}

		dbRecord := *data

		for j := 0; j < len(dbRecord); j++ {
			dbListDataRecord.RecordList.([]hktmodels.DBSearchPatientPathologyRecordResponseFilterDataModel)[i].DocumentList = append(dbListDataRecord.RecordList.([]hktmodels.DBSearchPatientPathologyRecordResponseFilterDataModel)[i].DocumentList, dbRecord[j])
		}

	}

	dataListResponse.FilteredRecords = dbListDataRecord.RecordCount
	dataListResponse.Records = dbListDataRecord.RecordList

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched patient pathology records list data.")

	return true, dataListResponse

}

func (service PatientService) SelectPatientInfoByAdmissionId(req lmodels.APIPatientInfoRequest) (bool, interface{}) {

	dbErr, data := dbaccess.GetPatientById(service.ExeCtx.SessionInfo.Product.NodeDbConn, req.PatientId)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting patient info by id.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *data

	if req.AdmissionId != nil {

		dbErr, personaldetailsdata := dbaccess.GetPersonalDetailsByAdmissionId(service.ExeCtx.SessionInfo.Product.NodeDbConn, *req.AdmissionId)
		if dbErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting patient info by id.", dbErr)

			errModel := gmodels.APIResponseError{}
			errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
			return false, errModel
		}

		personaldetailsdbRecord := *personaldetailsdata

		if len(personaldetailsdbRecord) > 0 {
			dbRecord[0].Age = personaldetailsdbRecord[0].Age
		}
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched patient info")
	return true, dbRecord[0]
}

func (service PatientService) PatientTreatmentAdd(req lmodels.APIPatientTreatmentAddRequest) (isSuccess bool, successErrorData interface{}) {

	dbRowModel := &hktmodels.DBPatientTreatmentInsertRowModel{}
	dbRowModel.DBPatientTreatmentDataModel = req.DBPatientTreatmentDataModel
	dbRowModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	dbRowModel.UpdatedBy = service.ExeCtx.SessionInfo.UserID
	dbRowModel.Uuid = ghelper.GenerateUUID()

	dbTxErr, tx := dbaccess.GetDBTransaction(service.ExeCtx.SessionInfo.Product.NodeDbConn)

	if dbTxErr != nil {

		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Transaction Error.", dbTxErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbErr, insertedId := dbaccess.SplHpftTreatmentTblInsert(tx, dbRowModel)
	if dbErr != nil {

		txErr := tx.Rollback()

		if txErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to rollback transaction", txErr)
		}

		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while adding patient treatment data.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	// insert treatment doc tbl data

	if len(req.DocumentUUIDList) != 0 {

		for _, documentUUID := range req.DocumentUUIDList {

			dbErr, docdata := dbaccess.GetDocumentDataByDocumentUUID(service.ExeCtx.SessionInfo.Product.NodeDbConn, documentUUID)
			if dbErr != nil {
				logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting patient info by uuid.", dbErr)

				errModel := gmodels.APIResponseError{}
				errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
				return false, errModel
			}

			docdbRecord := *docdata

			dBPatientTreatmentDocInsertRowModel := &hktmodels.DBPatientTreatmentDocInsertRowModel{}
			dBPatientTreatmentDocInsertRowModel.TreatmentId = insertedId
			dBPatientTreatmentDocInsertRowModel.DocumentId = docdbRecord[0].DocId

			dbErr, _ = dbaccess.SplHpftTreatmentDocTblInsert(tx, dBPatientTreatmentDocInsertRowModel)
			if dbErr != nil {

				txErr := tx.Rollback()

				if txErr != nil {
					logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to rollback transaction", txErr)
				}

				logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while adding patient treatment doc data.", dbErr)

				errModel := gmodels.APIResponseError{}
				errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
				return false, errModel
			}
		}

	}

	txErr := tx.Commit()

	if txErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to commit transaction", txErr)
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	// handler for - notify db changes
	if dbmgr.DefaultPostDataChangeHandler != nil {

		dataChangeHandlerConfigModel := gmodels.DataChangeHandlerConfigModel{}
		dataChangeHandlerConfigModel.ChangedData = dbRowModel
		dataChangeHandlerConfigModel.ChangeType = gmodels.DB_OPERATION_INSERT_UPDATE

		dbmgr.DefaultPostDataChangeHandler(constants.DB_SPL_HPFT_TREATMENT_TBL, dataChangeHandlerConfigModel)
		dbmgr.DefaultPostDataChangeHandler(constants.DB_SPL_HPFT_TREATMENT_DOC_TBL, dataChangeHandlerConfigModel)
	}

	addResponse := gmodels.APIRecordAddResponse{}
	addResponse.RecordID = insertedId

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient treament data added succesfully")

	return true, addResponse
}

func (service PatientService) PatientPathologyRecordAdd(req lmodels.APIPatientPathologyRecordAddRequest) (isSuccess bool, successErrorData interface{}) {

	dbRowModel := &hktmodels.DBPatientPathologyRecordInsertRowModel{}
	dbRowModel.DBPatientPathologyRecordDataModel = req.DBPatientPathologyRecordDataModel
	dbRowModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	dbRowModel.UpdatedBy = service.ExeCtx.SessionInfo.UserID
	dbRowModel.Uuid = ghelper.GenerateUUID()

	dbTxErr, tx := dbaccess.GetDBTransaction(service.ExeCtx.SessionInfo.Product.NodeDbConn)

	if dbTxErr != nil {

		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Transaction Error.", dbTxErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbErr, insertedId := dbaccess.SplHpftPathologyRecordTblInsert(tx, dbRowModel)
	if dbErr != nil {

		txErr := tx.Rollback()

		if txErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to rollback transaction", txErr)
		}

		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while adding patient pathology record data.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	// insert pathology record doc tbl data

	if len(req.DocumentUUIDList) != 0 {

		for _, documentUUID := range req.DocumentUUIDList {

			dbErr, docdata := dbaccess.GetDocumentDataByDocumentUUID(service.ExeCtx.SessionInfo.Product.NodeDbConn, documentUUID)
			if dbErr != nil {
				logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting patient info by id.", dbErr)

				errModel := gmodels.APIResponseError{}
				errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
				return false, errModel
			}

			docdbRecord := *docdata

			dBPatientPathologyRecordDocInsertRowModel := &hktmodels.DBPatientPathologyRecordDocInsertRowModel{}
			dBPatientPathologyRecordDocInsertRowModel.PathologyId = insertedId
			dBPatientPathologyRecordDocInsertRowModel.DocumentId = docdbRecord[0].DocId

			dbErr, _ = dbaccess.SplHpftPathologyRecordDocTblInsert(tx, dBPatientPathologyRecordDocInsertRowModel)
			if dbErr != nil {

				txErr := tx.Rollback()

				if txErr != nil {
					logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to rollback transaction", txErr)
				}

				logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while adding patient pathology record doc tbl data.", dbErr)

				errModel := gmodels.APIResponseError{}
				errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
				return false, errModel
			}
		}

	}

	txErr := tx.Commit()

	if txErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to commit transaction", txErr)
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	// handler for - notify db changes
	if dbmgr.DefaultPostDataChangeHandler != nil {

		dataChangeHandlerConfigModel := gmodels.DataChangeHandlerConfigModel{}
		dataChangeHandlerConfigModel.ChangedData = dbRowModel
		dataChangeHandlerConfigModel.ChangeType = gmodels.DB_OPERATION_INSERT_UPDATE

		dbmgr.DefaultPostDataChangeHandler(constants.DB_SPL_HPFT_PATHOLOGY_RECORD_TBL, dataChangeHandlerConfigModel)
		dbmgr.DefaultPostDataChangeHandler(constants.DB_SPL_HPFT_PATHOLOGY_RECORD_DOC_TBL, dataChangeHandlerConfigModel)
	}

	addResponse := gmodels.APIRecordAddResponse{}
	addResponse.RecordID = insertedId

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient pathology record data added succesfully")

	return true, addResponse
}

func (service PatientService) UserPateintAssociate(req lmodels.APIUserPatientAsscociationRequest) (isSuccess bool, successErrorData interface{}) {

	dbRowModel := &hktmodels.DBPatientMonitorMappingInsertRowModel{}
	dbRowModel.UsrId = req.UsrId
	dbRowModel.SpId = req.SpId
	dbRowModel.PatientId = req.PatientId
	dbRowModel.CpmId = service.ExeCtx.GetCPMID()
	dbRowModel.UpdatedBy = service.ExeCtx.SessionInfo.UserID
	dbRowModel.Uuid = ghelper.GenerateUUID()

	if req.PatientId == nil {
		dbPatientMonitorMappingDeleteRowModel := &hktmodels.DBPatientMonitorMappingDeleteRowModel{}
		dbPatientMonitorMappingDeleteRowModel.CpmId = dbRowModel.CpmId
		dbPatientMonitorMappingDeleteRowModel.UsrId = dbRowModel.UsrId
		dbPatientMonitorMappingDeleteRowModel.SpId = dbRowModel.SpId
		service.UserPatientAsscociationRemove(dbPatientMonitorMappingDeleteRowModel)
	} else if req.SpId == nil {
		dbPatientMonitorMappingDeleteRowModel := &hktmodels.DBPatientMonitorMappingDeleteRowModel{}
		dbPatientMonitorMappingDeleteRowModel.CpmId = dbRowModel.CpmId
		dbPatientMonitorMappingDeleteRowModel.UsrId = dbRowModel.UsrId
		service.UserPatientAsscociationRemove(dbPatientMonitorMappingDeleteRowModel)
	}

	dbErr, insertedId := dbaccess.PatientUserAssociation(service.ExeCtx.GetNodeDBConnection(), dbRowModel)
	if dbErr != nil {

		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while associated user and patient.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	addResponse := gmodels.APIRecordAddResponse{}
	addResponse.RecordID = insertedId

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "New user and patient associated succesfully")

	return true, addResponse
}

func (service PatientService) UserPatientAsscociationRemove(reqdata *hktmodels.DBPatientMonitorMappingDeleteRowModel) (isSuccess bool, successErrorData interface{}) {

	reqdata.CpmId = service.ExeCtx.GetCPMID()

	dbErr, affectedRow := dbaccess.PatientUserDeAssociation(service.ExeCtx.GetNodeDBConnection(), reqdata)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while deassociating user patient.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	if affectedRow == 0 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "User patient association removed successfully.")

	return true, nil
}
