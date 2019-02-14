package patient

import (
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	lmodels "opensoach.com/hpft/api/models"
	"opensoach.com/hpft/api/webserver/patient/dbaccess"
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

	dbErr, affectedRow := dbaccess.UpdateByFilter(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating patient info.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	if affectedRow == 0 {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Update request has no updated data.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient data updated successfully.")

	return true, nil
}

func (service PatientService) UpdateStatus(reqData *hktmodels.DBPatientUpdateStatusRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	reqData.DischargedOn = ghelper.GetCurrentTime()

	dbErr, affectedRow := dbaccess.UpdatePatientStatus(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while udating patient status.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	if affectedRow == 0 {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Update request has no updated data.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

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

	if len(dbRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

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

	if len(dbRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched patient info")
	return true, dbRecord
}

func (service PatientService) AdmissionAdd(req lmodels.APIAdmissionAddRequest) (isSuccess bool, successErrorData interface{}) {

	dbRowModel := &hktmodels.DBAdmissionTblInsertRowModel{}
	dbRowModel.DBAdmissionTblDataModel = req.DBAdmissionTblDataModel
	dbRowModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	dbRowModel.UpdatedBy = service.ExeCtx.SessionInfo.UserID

	dbErr, insertedId := dbaccess.AdmissionTblInsert(service.ExeCtx.SessionInfo.Product.NodeDbConn, dbRowModel)
	if dbErr != nil {

		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while adding new patient admission info.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	addResponse := gmodels.APIRecordAddResponse{}
	addResponse.RecordID = insertedId

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient admission info added succesfully")

	return true, addResponse
}

func (service PatientService) AdmissionUpdate(reqData *hktmodels.DBAdmissionTblUpdateRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	reqData.UpdatedBy = service.ExeCtx.SessionInfo.UserID

	dbErr, affectedRow := dbaccess.AdmissionTblUpdateByFilter(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating patient admission info.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	if affectedRow == 0 {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Update request has no updated data.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
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

	if len(dbRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched patient admission info")
	return true, dbRecord[0]
}

func (service PatientService) PersonalDetialsAdd(req lmodels.APIPersonalDetailsAddRequest) (isSuccess bool, successErrorData interface{}) {

	dbRowModel := &hktmodels.DBPersonalDetailsInsertRowModel{}
	dbRowModel.DBPersonalDetailsDataModel = req.DBPersonalDetailsDataModel
	dbRowModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	dbRowModel.UpdatedBy = service.ExeCtx.SessionInfo.UserID

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

	dbErr, affectedRow := dbaccess.PersonalDetailsUpdateByFilter(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating patient personal details.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	if affectedRow == 0 {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Update request has no updated data.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient personal details updated successfully.")

	return true, nil
}

func (service PatientService) PersonalDetailsUpdatePersonAccompanying(reqData *hktmodels.DBPersonalDetailsUpdatePersonAccompanyingRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	reqData.UpdatedBy = service.ExeCtx.SessionInfo.UserID

	dbErr, affectedRow := dbaccess.PersonalDetailsUpdatePersonAccompanying(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating patient personal details.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	if affectedRow == 0 {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Update request has no updated data.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
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

	if len(dbRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched patient personal details.")
	return true, dbRecord[0]
}

func (service PatientService) MedicalDetialsAdd(req lmodels.APIMedicalDetailsAddRequest) (isSuccess bool, successErrorData interface{}) {

	dbRowModel := &hktmodels.DBMedicalDetailsInsertRowModel{}
	dbRowModel.DBMedicalDetailsDataModel = req.DBMedicalDetailsDataModel
	dbRowModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	dbRowModel.UpdatedBy = service.ExeCtx.SessionInfo.UserID

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

	dbErr, affectedRow := dbaccess.MedicalDetailsUpdateByFilter(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating patient medical details.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	if affectedRow == 0 {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Update request has no updated data.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient medical details updated successfully.")

	return true, nil
}

func (service PatientService) MedicalDetailsUpdatePresentComplaints(reqData *hktmodels.DBMedicalDetailsUpdatePresentComplaintsRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	reqData.UpdatedBy = service.ExeCtx.SessionInfo.UserID

	dbErr, affectedRow := dbaccess.MedicalDetailsUpdatePresentComplaints(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating patient medical details.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	if affectedRow == 0 {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Update request has no updated data.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient present complaints details updated successfully.")

	return true, nil
}

func (service PatientService) MedicalDetailsUpdateReasonForAdmission(reqData *hktmodels.DBMedicalDetailsUpdateReasonForAdmissionRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	reqData.UpdatedBy = service.ExeCtx.SessionInfo.UserID

	dbErr, affectedRow := dbaccess.MedicalDetailsUpdateReasonForAdmission(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating patient medical details.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	if affectedRow == 0 {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Update request has no updated data.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient reason for admission details updated successfully.")

	return true, nil
}

func (service PatientService) MedicalDetailsUpdateHistoryPresentIllness(reqData *hktmodels.DBMedicalDetailsUpdateHistoryPresentIllnessRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	reqData.UpdatedBy = service.ExeCtx.SessionInfo.UserID

	dbErr, affectedRow := dbaccess.MedicalDetailsUpdateHistoryPresentIllness(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating patient medical details.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	if affectedRow == 0 {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Update request has no updated data.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient history present illness details updated successfully.")

	return true, nil
}

func (service PatientService) MedicalDetailsUpdatePastHistory(reqData *hktmodels.DBMedicalDetailsUpdatePastHistoryRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	reqData.UpdatedBy = service.ExeCtx.SessionInfo.UserID

	dbErr, affectedRow := dbaccess.MedicalDetailsUpdatePastHistory(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating patient medical details.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	if affectedRow == 0 {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Update request has no updated data.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient past history details updated successfully.")

	return true, nil
}

func (service PatientService) MedicalDetailsUpdateTreatmentBeforeAdmission(reqData *hktmodels.DBMedicalDetailsUpdateTreatmentBeforeAdmissionRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	reqData.UpdatedBy = service.ExeCtx.SessionInfo.UserID

	dbErr, affectedRow := dbaccess.MedicalDetailsUpdateTreatmentBeforeAdmission(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating patient medical details.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	if affectedRow == 0 {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Update request has no updated data.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient treatment before admission details updated successfully.")

	return true, nil
}

func (service PatientService) MedicalDetailsUpdateInvestigationBeforeAdmission(reqData *hktmodels.DBMedicalDetailsUpdateInvestigationBeforeAdmissionRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	reqData.UpdatedBy = service.ExeCtx.SessionInfo.UserID

	dbErr, affectedRow := dbaccess.MedicalDetailsUpdateInvestigationBeforeAdmission(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating patient medical details.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	if affectedRow == 0 {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Update request has no updated data.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient Investigation details before admission details updated successfully.")

	return true, nil
}

func (service PatientService) MedicalDetailsUpdateFamilyHistory(reqData *hktmodels.DBMedicalDetailsUpdateFamilyHistoryRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	reqData.UpdatedBy = service.ExeCtx.SessionInfo.UserID

	dbErr, affectedRow := dbaccess.MedicalDetailsUpdateFamilyHistory(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating patient medical details.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	if affectedRow == 0 {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Update request has no updated data.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient family history details updated successfully.")

	return true, nil
}

func (service PatientService) MedicalDetailsUpdateAllergies(reqData *hktmodels.DBMedicalDetailsUpdateAllergiesRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	reqData.UpdatedBy = service.ExeCtx.SessionInfo.UserID

	dbErr, affectedRow := dbaccess.MedicalDetailsUpdateAllergies(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating patient medical details.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	if affectedRow == 0 {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Update request has no updated data.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient allegies details updated successfully.")

	return true, nil
}

func (service PatientService) MedicalDetailsUpdatePersonalHistory(reqData *hktmodels.DBMedicalDetailsUpdatePersonalHistoryRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	reqData.UpdatedBy = service.ExeCtx.SessionInfo.UserID

	dbErr, affectedRow := dbaccess.MedicalDetailsUpdatePersonalHistory(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating patient medical details.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	if affectedRow == 0 {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Update request has no updated data.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
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

	if len(dbRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched patient medical details.")
	return true, dbRecord[0]
}

func (service PatientService) SelectAllPatientConf(patientID int64) (bool, interface{}) {

	cpmID := service.ExeCtx.SessionInfo.Product.CustProdID

	dbErr, data := dbaccess.GetPatientConfList(service.ExeCtx.SessionInfo.Product.NodeDbConn, cpmID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting patient conf list.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *data

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched patient conf list")
	return true, dbRecord
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

	if len(dbRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched patient config.")
	return true, dbRecord[0]
}

func (service PatientService) PatientConfUpdate(reqData *hktmodels.DBPatientConfUpdateRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID

	dbErr, affectedRow := dbaccess.PatientConfUpdateByFilter(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating patient config.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	if affectedRow == 0 {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Update request has no updated data.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
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

	if len(dbRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched patient admission status info")
	return true, dbRecord[0]
}
