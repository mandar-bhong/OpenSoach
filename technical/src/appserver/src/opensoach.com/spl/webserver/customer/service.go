package customer

import (
	gmodels "opensoach.com/models"
)

type CustomerService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service CustomerService) UpdateCustomerDetails() (isSuccess bool, successErrorData interface{}) {
	return false, nil
}
