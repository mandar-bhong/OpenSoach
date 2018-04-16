package customer

import (
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	lmodels "opensoach.com/spl/models"
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

func (service CustomerService) GetCustomerInfo(customerID int64) (bool, interface{}) {

	dbErr, customerDetails := dbaccess.GetCustomerById(repo.Instance().Context.Master.DBConn, customerID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}
	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched customer master details")
	return true, customerDetails
}

func (service CustomerService) GetCustomerDetailsInfo(customerID int64) (bool, interface{}) {

	dbErr, customerDetails := dbaccess.GetCustomerDetailsById(repo.Instance().Context.Master.DBConn, customerID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}
	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched customer details")
	return true, customerDetails
}

func (service CustomerService) GetCorpInfo(customerID int64) (bool, interface{}) {

	dbErr, customerDetails := dbaccess.GetCorpDetailsById(repo.Instance().Context.Master.DBConn, customerID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}
	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched customer corporate details.")
	return true, customerDetails
}

func (CustomerService) GetCustomerDataList(custListReqData lmodels.DataListRequest) (bool, interface{}) {

	custListResData := lmodels.DataListResponse{}

	filterModel := custListReqData.Filter.(*lmodels.DBSearchCustomerDataModel)

	dbErr, customerFilteredRecords := dbaccess.GetSplMasterCustomerTableTotalFilteredRecords(repo.Instance().Context.Master.DBConn, filterModel)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}
	dbCustomerFilteredRecords := *customerFilteredRecords
	custListResData.FilteredRecords = dbCustomerFilteredRecords.TotalRecords

	CurrentPage := custListReqData.CurrentPage
	startingRecord := ((CurrentPage - 1) * custListReqData.Limit)

	dbErr, customerFilterData := dbaccess.SplMasterCustomerTableSelectByFilter(repo.Instance().Context.Master.DBConn, custListReqData, filterModel, startingRecord)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbCustomerFilterRecord := *customerFilterData

	for i := 0; i < len(dbCustomerFilterRecord); i++ {
		custListResData.Records = append(custListResData.Records, dbCustomerFilterRecord[i])
	}

	return true, custListResData

}
