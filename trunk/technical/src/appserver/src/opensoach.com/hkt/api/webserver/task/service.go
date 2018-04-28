package task

import (
	lmodels "opensoach.com/hkt/api/models"
	"opensoach.com/hkt/api/webserver/task/dbaccess"
	hktmodels "opensoach.com/hkt/models"
	gmodels "opensoach.com/models"
)

var SUB_MODULE_NAME = "SPL.Task"

type TaskService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service TaskService) Add(req lmodels.APITaskAddRequest) (isSuccess bool, successErrorData interface{}) {

	dbTaskLibRowModel := &hktmodels.DBTaskLibRowModel{}
	dbTaskLibRowModel.CPMID = service.ExeCtx.SessionInfo.Product.CustProdID
	dbTaskLibRowModel.SPCategoryID = req.SPCategoryID
	dbTaskLibRowModel.TaskName = req.Name
	dbTaskLibRowModel.Description = req.Discription

	dbaccess.Insert(service.ExeCtx.SessionInfo.Product.NodeDbConn, dbTaskLibRowModel)
	return false, nil
}
