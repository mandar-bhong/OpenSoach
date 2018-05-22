package task

import (
	"opensoach.com/core/logger"
	lmodels "opensoach.com/hkt/api/models"
	"opensoach.com/hkt/api/webserver/task/dbaccess"
	hktmodels "opensoach.com/hkt/models"
	gmodels "opensoach.com/models"
)

var SUB_MODULE_NAME = "HKT.API.Task"

type TaskService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service TaskService) Add(req lmodels.APITaskAddRequest) (isSuccess bool, successErrorData interface{}) {

	dbTaskLibRowModel := &hktmodels.DBTaskLibInsertRowModel{}
	dbTaskLibRowModel.CPMID = service.ExeCtx.SessionInfo.Product.CustProdID
	dbTaskLibRowModel.SpcId = req.SpcId
	dbTaskLibRowModel.TaskName = req.TaskName
	dbTaskLibRowModel.ShortDesc = req.ShortDesc

	dbErr, insertedId := dbaccess.Insert(service.ExeCtx.SessionInfo.Product.NodeDbConn, dbTaskLibRowModel)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	addResponse := gmodels.APIRecordAddResponse{}
	addResponse.RecordID = insertedId

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "New Task Added succesfully")

	return true, addResponse
}

func (service TaskService) Update(reqData *hktmodels.DBTaskLibUpdateRowModel) (isSuccess bool, successErrorData interface{}) {

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

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Task updated successfully.")

	return true, nil
}

func (service TaskService) SelectById(taskID int64) (bool, interface{}) {

	dbErr, taskData := dbaccess.TaskSelectByID(service.ExeCtx.SessionInfo.Product.NodeDbConn, taskID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *taskData

	if len(dbRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched task info")
	return true, dbRecord[0]
}

func (service TaskService) SelectByAll() (bool, interface{}) {

	dbErr, taskData := dbaccess.TaskSelectAll(service.ExeCtx.SessionInfo.Product.NodeDbConn)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *taskData

	if len(dbRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched all task")
	return true, dbRecord[0]
}

func (service TaskService) SelectBySPCId(taskID int64) (bool, interface{}) {

	dbErr, taskData := dbaccess.TaskSelectBySPCID(service.ExeCtx.SessionInfo.Product.NodeDbConn, taskID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched task info")
	return true, taskData
}
