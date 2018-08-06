package patient

import (
	"opensoach.com/core/logger"
	lmodels "opensoach.com/hpft/api/models"
	repo "opensoach.com/hpft/api/repository"
	"opensoach.com/hpft/api/webserver/patient/dbaccess"
	hktconst "opensoach.com/hpft/constants"
	hktmodels "opensoach.com/hpft/models"
	gmodels "opensoach.com/models"
)

var SUB_MODULE_NAME = "HPFT.API.Patient"

type PatientService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service PatientService) Add(req lmodels.APIPatientAddRequest) (isSuccess bool, successErrorData interface{}) {

	dbServiceInstanceInsertRowModel := &hktmodels.DBServiceInstanceInsertRowModel{}
	dbServiceInstanceInsertRowModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	dbServiceInstanceInsertRowModel.ServConfId = req.PatientFileTemplateID
	dbServiceInstanceInsertRowModel.SpId = req.SpId

	dbTxErr, tx := dbaccess.GetDBTransaction(service.ExeCtx.SessionInfo.Product.NodeDbConn)

	if dbTxErr != nil {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbErr, insertedServInId := dbaccess.ServiceInstanceInsert(tx, dbServiceInstanceInsertRowModel)
	if dbErr != nil {

		txErr := tx.Rollback()

		if txErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to rollback transaction", txErr)
		}

		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while adding service instance data.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRowModel := &hktmodels.DBPatientInsertRowModel{}
	dbRowModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	dbRowModel.MedicalDetails = req.MedicalDetails
	dbRowModel.PatientDetails = req.PatientDetails
	dbRowModel.PatientFileTemplateID = req.PatientFileTemplateID
	dbRowModel.SpId = req.SpId
	dbRowModel.ServInId = insertedServInId
	dbRowModel.Status = req.Status

	dbErr, insertedId := dbaccess.Insert(tx, dbRowModel)
	if dbErr != nil {
		txErr := tx.Rollback()

		if txErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to rollback transaction", txErr)
		}

		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while adding new patient.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	txErr := tx.Commit()

	if txErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to commit transaction", txErr)
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	addResponse := gmodels.APIRecordAddResponse{}
	addResponse.RecordID = insertedId

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "New patient added succesfully")

	return true, addResponse
}

func (service PatientService) GetPatientList() (bool, interface{}) {

	cpmID := service.ExeCtx.SessionInfo.Product.CustProdID

	dbErr, listData := dbaccess.GetPatientList(service.ExeCtx.SessionInfo.Product.NodeDbConn, cpmID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched patient data list.")

	return true, listData

}

func (service PatientService) Update(reqData *hktmodels.DBPatientUpdateRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID

	dbErr, affectedRow := dbaccess.UpdateByFilter(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	if affectedRow == 0 {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient data updated successfully.")

	return true, nil
}

func (service PatientService) UpdateStatus(reqData *hktmodels.DBPatientUpdateStatusRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID

	dbErr, affectedRow := dbaccess.UpdatePatientStatus(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	if affectedRow == 0 {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	taskPatientStatusUpdated := &hktmodels.TaskPatientStatusUpdated{}
	taskPatientStatusUpdated.PatientId = reqData.PatientId
	taskPatientStatusUpdated.CpmId = reqData.CpmId

	isSendSuccess := repo.Instance().
		SendTaskToServer(hktconst.TASK_HKT_API_PATIENT_STATUS_UPDATED,
			service.ExeCtx.SessionToken, taskPatientStatusUpdated)

	if isSendSuccess == false {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Unable to submit task to server.", nil)
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "patient status updated successfully.")

	return true, nil
}
