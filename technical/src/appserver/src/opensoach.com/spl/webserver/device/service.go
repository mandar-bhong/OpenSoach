package device

import (
	"time"

	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	lmodels "opensoach.com/spl/models"
	repo "opensoach.com/spl/repository"
	"opensoach.com/spl/webserver/device/dbaccess"
)

var SUB_MODULE_NAME = "SPL.Device"

type DeviceService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service DeviceService) AddDevice(reqData *lmodels.DBSplMasterDeviceRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.DevStateSince = time.Now()

	dbErr, insertedId := dbaccess.SplMasterDeviceTableInsert(repo.Instance().Context.Master.DBConn, reqData)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	response := lmodels.RecordIdResponse{}
	response.RecId = insertedId

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Device data added successfully.")

	return true, response
}

func (service DeviceService) UpdateDevState(reqData *lmodels.DevStateRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.DevStateSince = time.Now()

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

func (service DeviceService) UpdateDeviceDetails(reqData *lmodels.DBSplMasterDevDetailsRowModel) (isSuccess bool, successErrorData interface{}) {

	if reqData.DevId == 0 {
		dbErr, rsltData := dbaccess.GetDeviceId(repo.Instance().Context.Master.DBConn, service.ExeCtx.SessionInfo.CpmID)
		if dbErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "DB Error occured while login.", dbErr)
			errModel := gmodels.APIResponseError{}
			errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
			return false, errModel
		}
		reqData.DevId = rsltData.DevId
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

		response := lmodels.RecordIdResponse{}
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

func (DeviceService) GetDeviceDataList(listReqData lmodels.DataListRequest) (bool, interface{}) {

	listResData := lmodels.DataListResponse{}

	filterModel := listReqData.Filter.(*lmodels.DBSearchDeviceRequestFilterDataModel)

	dbErr, filterRecordsCount := dbaccess.GetDeviceFilterRecordsCount(repo.Instance().Context.Master.DBConn, filterModel)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}
	dbFilterRecordsCount := *filterRecordsCount
	listResData.FilteredRecords = dbFilterRecordsCount.TotalRecords

	CurrentPage := listReqData.CurrentPage
	startingRecord := ((CurrentPage - 1) * listReqData.Limit)

	dbErr, filterData := dbaccess.GetDeviceListData(repo.Instance().Context.Master.DBConn, listReqData, filterModel, startingRecord)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbfilterData := *filterData
	listResData.Records = dbfilterData

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched device list data.")

	return true, listResData

}
