package device

import (
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	"opensoach.com/spl/constants"
	lmodels "opensoach.com/spl/models"
	repo "opensoach.com/spl/repository"
	"opensoach.com/spl/webserver/device/dbaccess"
)

var SUB_MODULE_NAME = "SPL.Device"

type DeviceService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service DeviceService) AddDevice(reqData *lmodels.DBSplMasterDeviceRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.DevStateSince = ghelper.GetCurrentTime()

	dbErr, insertedId := dbaccess.SplMasterDeviceTableInsert(repo.Instance().Context.Master.DBConn, reqData)

	if dbErr != nil {

		errModel := gmodels.APIResponseError{}
		errHandledIsSuccess, errorCode := ghelper.GetApplicationErrorCodeFromDBError(dbErr)

		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		if errHandledIsSuccess == true {
			errModel.Code = errorCode
			return false, errModel
		}

		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	response := gmodels.APIRecordIdResponse{}
	response.RecId = insertedId

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Device data added successfully.")

	return true, response
}

func (service DeviceService) UpdateDevState(reqData *lmodels.DBDevStateRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.DevStateSince = ghelper.GetCurrentTime()

	dbErr, _ := dbaccess.UpdateDeviceState(repo.Instance().Context.Master.DBConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Device state updated successfully.")

	return true, nil
}

func (service DeviceService) UpdateDeviceDetails(reqData *lmodels.DBSplMasterDevDetailsRowModel, userType string) (isSuccess bool, successErrorData interface{}) {

	dbErr, deviceData := dbaccess.GetDeviceById(repo.Instance().Context.Master.DBConn, reqData.DevId)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbDevRecord := *deviceData

	if len(dbDevRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	if userType == "CU" {
		dbErr, rsltData := dbaccess.GetDeviceId(repo.Instance().Context.Master.DBConn, service.ExeCtx.SessionInfo.Product.CustProdID, reqData.DevId)
		if dbErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "DB Error occured while login.", dbErr)
			errModel := gmodels.APIResponseError{}
			errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
			return false, errModel
		}

		dbRecordData := *rsltData

		if len(dbRecordData) < 1 {
			errModel := gmodels.APIResponseError{}
			errModel.Code = constants.MOD_ERR_DEVICE_CUSTOMER_PRODUCT_MAPPING
			return false, errModel
		}
	}

	dbErr, detailsData := dbaccess.GetSplMasterDeviceDetailsTableById(repo.Instance().Context.Master.DBConn, reqData.DevId)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "DB Error occured while login.", dbErr)
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbDetailsRecord := *detailsData

	if len(dbDetailsRecord) < 1 {
		dbErr, insertedId := dbaccess.SplMasterDeviceDetailsTableInsert(repo.Instance().Context.Master.DBConn, reqData)
		if dbErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

			errModel := gmodels.APIResponseError{}
			errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
			return false, errModel
		}

		response := gmodels.APIRecordIdResponse{}
		response.RecId = insertedId

		logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Device details inserted successfully.")

		return true, response

	} else {
		dbErr, affectedRow := dbaccess.SplMasterDeviceDetailsTableUpdate(repo.Instance().Context.Master.DBConn, reqData)
		if dbErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

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

		logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Device details updated Successfully.")

		return true, nil
	}

}

func (DeviceService) GetDeviceDataList(listReqData gmodels.APIDataListRequest) (bool, interface{}) {

	dataListResponse := gmodels.APIDataListResponse{}

	filterModel := listReqData.Filter.(*lmodels.DBSearchDeviceRequestFilterDataModel)

	CurrentPage := listReqData.CurrentPage
	startingRecord := ((CurrentPage - 1) * listReqData.Limit)

	dbErr, listData := dbaccess.GetDeviceListData(repo.Instance().Context.Master.DBConn, filterModel, listReqData, startingRecord)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbListDataRecord := *listData

	dataListResponse.FilteredRecords = dbListDataRecord.RecordCount
	dataListResponse.Records = dbListDataRecord.RecordList

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched device list data.")

	return true, dataListResponse

}

func (service DeviceService) AssociateDevWithCust(reqData *lmodels.DBDevCustRowModel) (isSuccess bool, successErrorData interface{}) {

	dbErr, _ := dbaccess.SetDeviceCustId(repo.Instance().Context.Master.DBConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Device associated with customer, successfully.")

	return true, nil
}

func (service DeviceService) AssociateDevWithCustProduct(reqData *lmodels.DBSplCpmDevRowModel) (isSuccess bool, successErrorData interface{}) {

	dbErr, devData := dbaccess.GetDeviceById(repo.Instance().Context.Master.DBConn, reqData.DevId)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbDevRecord := *devData

	if len(dbDevRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	dbErr, data := dbaccess.GetCustIdByCpmId(repo.Instance().Context.Master.DBConn, reqData.CpmId)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbrecorddata := *data

	if dbDevRecord[0].CustId == nil {

		custID := dbrecorddata[0].CustId

		devicecustmap := &lmodels.DBDevCustRowModel{}
		devicecustmap.CustId = custID
		devicecustmap.DevId = reqData.DevId

		dberr, _ := dbaccess.SetDeviceCustId(repo.Instance().Context.Master.DBConn, devicecustmap)
		if dberr != nil {
			logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

			errModel := gmodels.APIResponseError{}
			errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
			return false, errModel
		}
	} else if *dbDevRecord[0].CustId != dbrecorddata[0].CustId {
		errModel := gmodels.APIResponseError{}
		errModel.Code = constants.MOD_ERR_DEVICE_UNAVAILABLE_FOR_CUSTOMER_PRODUCT_MAPPING
		return false, errModel
	}

	dbErr, insertedId := dbaccess.CpmDevTableInsert(repo.Instance().Context.Master.DBConn, reqData)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	taskDevProdAsscociatedModel := gmodels.TaskDevProdAsscociatedModel{}
	taskDevProdAsscociatedModel.CpmId = reqData.CpmId
	taskDevProdAsscociatedModel.DevId = reqData.DevId

	if isSuccess := repo.Instance().SendTaskToServer(gmodels.TASK_API_DEV_PROD_ASSOCIATED, service.ExeCtx.SessionToken, taskDevProdAsscociatedModel); isSuccess == false {
		logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while submiting task for cust prod assoc")
	}

	response := gmodels.APIRecordIdResponse{}
	response.RecId = insertedId

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Device associated with customer product, successfully.")

	return true, nil
}

func (service DeviceService) GetDeviceDetailsInfo(deviceID int64, userType string) (bool, interface{}) {

	dbErr, deviceData := dbaccess.GetDeviceById(repo.Instance().Context.Master.DBConn, deviceID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbDevRecord := *deviceData

	if len(dbDevRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	if userType == "CU" {
		dbErr, rsltData := dbaccess.GetDeviceId(repo.Instance().Context.Master.DBConn, service.ExeCtx.SessionInfo.Product.CustProdID, deviceID)
		if dbErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "DB Error occured while login.", dbErr)
			errModel := gmodels.APIResponseError{}
			errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
			return false, errModel
		}

		dbRecordData := *rsltData

		if len(dbRecordData) < 1 {
			errModel := gmodels.APIResponseError{}
			errModel.Code = constants.MOD_ERR_DEVICE_CUSTOMER_PRODUCT_MAPPING
			return false, errModel
		}
	}

	dbErr, deviceDetails := dbaccess.GetSplMasterDeviceDetailsTableById(repo.Instance().Context.Master.DBConn, deviceID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *deviceDetails

	if len(dbRecord) < 1 {
		return true, nil
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched device details")
	return true, dbRecord[0]
}

func (service DeviceService) GetDeviceProdAssociation(devID int64) (bool, interface{}) {

	dbErr, data := dbaccess.GetDeviceAssociationByDevId(repo.Instance().Context.Master.DBConn, devID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecords := *data

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched Device Product association list")
	return true, dbRecords
}

func (service DeviceService) DeviceShortDataList() (bool, interface{}) {

	dbErr, listData := dbaccess.GetDeviceShortDataList(repo.Instance().Context.Master.DBConn)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched device short data list.")

	return true, listData

}
