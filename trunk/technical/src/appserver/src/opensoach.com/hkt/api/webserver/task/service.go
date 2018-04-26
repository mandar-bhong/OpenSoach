package task

import (
	lmodels "opensoach.com/hkt/api/models"
	gmodels "opensoach.com/models"
)

var SUB_MODULE_NAME = "SPL.Task"

type TaskService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service TaskService) Add(req lmodels.APITaskAddRequest) (isSuccess bool, successErrorData interface{}) {
	return false, nil
}
