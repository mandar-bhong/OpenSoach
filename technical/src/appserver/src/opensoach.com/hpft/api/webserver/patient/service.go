package patient

import (
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

func (service PatientService) Add(req lmodels.APIPatientAddRequest) (isSuccess bool, successErrorData interface{}) {

	dbRowModel := &hktmodels.DBPatientInsertRowModel{}
	dbRowModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	dbRowModel.MedicalDetails = req.MedicalDetails
	dbRowModel.PatientDetails = req.PatientDetails
	dbRowModel.PatientFileTemplate = req.PatientFileTemplate

	dbErr, insertedId := dbaccess.Insert(service.ExeCtx.SessionInfo.Product.NodeDbConn, dbRowModel)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	addResponse := gmodels.APIRecordAddResponse{}
	addResponse.RecordID = insertedId

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "New patient added succesfully")

	return true, addResponse
}
