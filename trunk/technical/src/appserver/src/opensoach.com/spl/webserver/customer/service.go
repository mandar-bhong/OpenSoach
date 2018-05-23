package customer

import (
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	"opensoach.com/spl/constants"
	lmodels "opensoach.com/spl/models"
	repo "opensoach.com/spl/repository"
	dbaccess "opensoach.com/spl/webserver/customer/dbaccess"
)

var SUB_MODULE_NAME = "SPL.Customer"

type CustomerService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service CustomerService) Add(req lmodels.CustomerAddRequest) (isSuccess bool, successErrorData interface{}) {

	dbSplMasterCustomerTableRowModel := lmodels.DBSplMasterCustomerTableRowModel{}

	dbSplMasterCustomerTableRowModel.CorpId = req.CorporationID
	dbSplMasterCustomerTableRowModel.CustName = req.CustomerName
	dbSplMasterCustomerTableRowModel.CustState = req.CustomerState
	dbSplMasterCustomerTableRowModel.CustStateSince = ghelper.GetCurrentTime()

	dbErr, dbData := dbaccess.AddCustomer(repo.Instance().Context.Master.DBConn, dbSplMasterCustomerTableRowModel)

	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while adding customer.", dbErr)
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	response := gmodels.APIRecordIdResponse{}
	response.RecId = dbData
	return true, response
}

func (service CustomerService) UpdateCustomerDetails(customerData lmodels.DBSplMasterCustDetailsTableRowModel) (isSuccess bool, successErrorData interface{}) {

	dbErr, customerDetailsData := dbaccess.GetCustomerDetailsById(repo.Instance().Context.Master.DBConn, customerData.CustId)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while Get customer details by id.", dbErr)
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbCustomerDetailsRecord := *customerDetailsData

	if len(dbCustomerDetailsRecord) < 1 {
		dbErr, customerInsertedId := dbaccess.CustomerDetailsTableInsert(repo.Instance().Context.Master.DBConn, customerData)
		if dbErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating customer.", dbErr)

			errModel := gmodels.APIResponseError{}
			errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
			return false, errModel
		}

		response := gmodels.APIRecordIdResponse{}
		response.RecId = customerInsertedId

		logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Customer details inserted successfully.")

		return true, response

	} else {
		dbErr, customerAffectedRow := dbaccess.CustomerDetailsTableUpdate(repo.Instance().Context.Master.DBConn, customerData)
		if dbErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating customer.", dbErr)

			errModel := gmodels.APIResponseError{}
			errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
			return false, errModel
		}

		response := gmodels.APIRecordIdResponse{}
		response.RecId = customerAffectedRow

		logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Customer details updated Successfully.")

		return true, nil
	}

}

func (service CustomerService) GetCustomerInfo(customerID int64) (bool, interface{}) {

	dbErr, customerDetails := dbaccess.GetCustomerById(repo.Instance().Context.Master.DBConn, customerID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *customerDetails

	if len(dbRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched customer master details")
	return true, dbRecord[0]
}

func (service CustomerService) GetCustomerDetailsInfo(customerID int64) (bool, interface{}) {

	dbErr, custData := dbaccess.GetCustomerById(repo.Instance().Context.Master.DBConn, customerID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbCustRecord := *custData

	if len(dbCustRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	dbErr, customerDetails := dbaccess.GetCustomerDetailsById(repo.Instance().Context.Master.DBConn, customerID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *customerDetails

	if len(dbRecord) < 1 {
		return true, nil
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched customer details")
	return true, dbRecord[0]
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

func (CustomerService) GetCustomerDataList(custListReqData gmodels.APIDataListRequest) (bool, interface{}) {

	dataListResponse := gmodels.APIDataListResponse{}

	filterModel := custListReqData.Filter.(*lmodels.DBSearchCustomerRequestFilterDataModel)

	CurrentPage := custListReqData.CurrentPage
	startingRecord := ((CurrentPage - 1) * custListReqData.Limit)

	dbErr, listData := dbaccess.GetCustList(repo.Instance().Context.Master.DBConn, filterModel, custListReqData, startingRecord)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbListDataRecord := *listData

	dataListResponse.FilteredRecords = dbListDataRecord.RecordCount
	dataListResponse.Records = dbListDataRecord.RecordList

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched customer list data.")

	return true, dataListResponse

}

func (service CustomerService) AssociateCustWithProduct(reqData *lmodels.DBCustProdMappingInsertRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmStateSince = ghelper.GetCurrentTime()

	dbErr, insertedId := dbaccess.CpmTableInsert(repo.Instance().Context.Master.DBConn, reqData)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	taskAPICustProdAssociatedModel := gmodels.TaskAPICustProdAssociatedModel{}
	taskAPICustProdAssociatedModel.CustId = reqData.CustId
	taskAPICustProdAssociatedModel.ProdId = reqData.ProdId
	taskAPICustProdAssociatedModel.DbiId = reqData.DbiId
	taskAPICustProdAssociatedModel.CpmId = insertedId

	if isSuccess := repo.Instance().SendTaskToServer(gmodels.TASK_API_CUST_PROD_ASSOCIATED, service.ExeCtx.SessionToken, taskAPICustProdAssociatedModel); isSuccess == false {
		logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while submiting task for cust prod assoc")
	}

	response := gmodels.APIRecordIdResponse{}
	response.RecId = insertedId

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Customer associated with product, successfully.")

	return true, response
}

func (service CustomerService) GetCustProdAssociation(customerID int64) (bool, interface{}) {

	dbErr, data := dbaccess.GetProdAssociationByCustId(repo.Instance().Context.Master.DBConn, customerID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecords := *data

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched Customer Product association list")
	return true, dbRecords
}

func (service CustomerService) UpdateCPMState(reqData *lmodels.DBCpmStateUpdateRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CpmStateSince = ghelper.GetCurrentTime()

	dbErr, _ := dbaccess.CpmStateUpdate(repo.Instance().Context.Master.DBConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "CPM state updated successfully.")

	return true, nil
}

func (service CustomerService) UpdateCust(reqData *lmodels.DBCustomerUpdateRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.CustStateSince = ghelper.GetCurrentTime()

	dbErr, affectedRow := dbaccess.CustomerUpdate(repo.Instance().Context.Master.DBConn, reqData)
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

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Customer data updated successfully.")

	return true, nil
}

func (CustomerService) CustShortDataList() (bool, interface{}) {

	dbErr, listData := dbaccess.GetCustShortDataList(repo.Instance().Context.Master.DBConn)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched customer short data list.")

	return true, listData

}

func (CustomerService) GetCustServicePoint(customerId int64) (bool, interface{}) {

	dbErr, customerSpModels := dbaccess.GetCustServicePoints(repo.Instance().Context.Master.DBConn, customerId)

	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	return true, customerSpModels

}

func (service CustomerService) CustServicePointAssociationCountUpdate(reqData lmodels.APICustSpCountUpdateRequest) (bool, interface{}) {

	servicepointrowmodel := &lmodels.DBServicepointInsertRowModel{}
	servicepointrowmodel.CpmId = reqData.CpmId
	servicepointrowmodel.SpState = constants.DB_SERVICE_POINT_STATE_ACTIVE
	servicepointrowmodel.SpStateSince = ghelper.GetCurrentTime()

	insertedIdList := []int64{}

	for i := 0; i < reqData.UpdateCount; i++ {
		dbErr, insertedId := dbaccess.SpInsert(repo.Instance().Context.Master.DBConn, servicepointrowmodel)
		if dbErr != nil {
			logger.Context().WithField("InputRequest", servicepointrowmodel).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

			errModel := gmodels.APIResponseError{}
			errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
			return false, errModel
		}

		insertedIdList = append(insertedIdList, insertedId)
	}

	taskCustServicePointAssociatedModel := gmodels.TaskCustServicePointAssociatedModel{}
	taskCustServicePointAssociatedModel.CpmId = reqData.CpmId
	taskCustServicePointAssociatedModel.SpIdList = insertedIdList
	taskCustServicePointAssociatedModel.SpState = servicepointrowmodel.SpState
	taskCustServicePointAssociatedModel.SpStateSince = servicepointrowmodel.SpStateSince

	if isSuccess := repo.Instance().SendTaskToServer(gmodels.TASK_API_CUST_SERVICE_POINT_ASSOCIATED, service.ExeCtx.SessionToken, taskCustServicePointAssociatedModel); isSuccess == false {
		logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while submiting task for cust prod assoc")
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Service point added successfully.")

	return true, nil

}
