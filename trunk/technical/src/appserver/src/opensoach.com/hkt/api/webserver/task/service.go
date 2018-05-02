package task

import (
	"opensoach.com/core/logger"
	lmodels "opensoach.com/hkt/api/models"
	"opensoach.com/hkt/api/webserver/task/dbaccess"
	hktmodels "opensoach.com/hkt/models"
	gmodels "opensoach.com/models"
)

var SUB_MODULE_NAME = "HKT.Task"

type TaskService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service TaskService) Add(req lmodels.APITaskAddRequest) (isSuccess bool, successErrorData interface{}) {

	dbTaskLibRowModel := &hktmodels.DBTaskLibRowModel{}
	dbTaskLibRowModel.CPMID = service.ExeCtx.SessionInfo.Product.CustProdID
	dbTaskLibRowModel.SPCategoryID = req.SPCategoryID
	dbTaskLibRowModel.TaskName = req.Name
	dbTaskLibRowModel.Description = req.Discription

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
