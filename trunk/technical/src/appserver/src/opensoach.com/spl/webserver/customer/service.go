package customer

import (
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	repo "opensoach.com/spl/repository"
	"opensoach.com/spl/webserver/customer/dbaccess"
)

var SUB_MODULE_NAME = "SPL.Customer"

type CustomerService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service CustomerService) UpdateCustomerDetails() (isSuccess bool, successErrorData interface{}) {
	return false, nil
}

func (service CustomerService) GetCustomerDetails(customerID int64) (bool, interface{}) {

	dbErr, customerDetails := dbaccess.GetCustomerById(repo.Instance().Context.Dynamic.DB, customerID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	return true, customerDetails
}
