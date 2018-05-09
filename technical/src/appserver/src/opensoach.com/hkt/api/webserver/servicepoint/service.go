package servicepoint

import (
	"opensoach.com/core/logger"
	lmodels "opensoach.com/hkt/api/models"
	"opensoach.com/hkt/api/webserver/servicepoint/dbaccess"
	hktmodels "opensoach.com/hkt/models"
	gmodels "opensoach.com/models"
)

var SUB_MODULE_NAME = "HKT.API.Servicepoint"

type ServicePointService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service ServicePointService) SpCategoryAdd(req lmodels.APISpCategoryAddRequest) (isSuccess bool, successErrorData interface{}) {

	dbRowModel := &hktmodels.DBSpCategoryInsertRowModel{}
	dbRowModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	dbRowModel.SpcName = req.SpcName
	dbRowModel.ShortDesc = req.ShortDesc

	dbErr, insertedId := dbaccess.SpCategoryInsert(service.ExeCtx.SessionInfo.Product.NodeDbConn, dbRowModel)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	addResponse := gmodels.APIRecordAddResponse{}
	addResponse.RecordID = insertedId

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "New SpCategory Added succesfully")

	return true, addResponse
}

func (service ServicePointService) SpUpdate(reqData *hktmodels.DBSpUpdateRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID

	dbErr, affectedRow := dbaccess.SpUpdateByFilter(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqData)
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

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Service Point updated successfully.")

	return true, nil
}

func (service ServicePointService) FopSpAdd(req lmodels.APIFopSpAddRequest) (isSuccess bool, successErrorData interface{}) {

	dbRowModel := &hktmodels.DBFopSpInsertRowModel{}
	dbRowModel.FopId = req.FopId
	dbRowModel.SpId = req.SpId
	dbRowModel.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID

	dbErr, insertedId := dbaccess.FopSpInsert(service.ExeCtx.SessionInfo.Product.NodeDbConn, dbRowModel)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	addResponse := gmodels.APIRecordAddResponse{}
	addResponse.RecordID = insertedId

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Service Point associated with Field operater  succesfully")

	return true, addResponse
}

func (service ServicePointService) FopSpDelete(reqdata *lmodels.APIFopSpDeleteRequest) (isSuccess bool, successErrorData interface{}) {

	dbErr, affectedRow := dbaccess.FopSpDelete(service.ExeCtx.SessionInfo.Product.NodeDbConn, reqdata)
	if dbErr != nil {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	if affectedRow == 0 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Service Point association with Field operater deleted successfully.")

	return true, nil
}
