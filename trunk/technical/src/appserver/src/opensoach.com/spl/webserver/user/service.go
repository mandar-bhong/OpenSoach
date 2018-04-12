package user

import (
	gmodels "opensoach.com/models"
)

var SUB_MODULE_NAME = "SPL.User"

type UserService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service UserService) UpdateUserDetails() (isSuccess bool, successErrorData interface{}) {

	return false, nil
}
