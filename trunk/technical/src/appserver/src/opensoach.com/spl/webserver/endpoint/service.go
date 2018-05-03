package endpoint

import (
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	"opensoach.com/spl/constants"
	lhelper "opensoach.com/spl/helper"
	lmodels "opensoach.com/spl/models"
	repo "opensoach.com/spl/repository"
	"opensoach.com/spl/webserver/endpoint/dbaccess"
)

var SUB_MODULE_NAME = "SPL.Endpoint"

type EndpointService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (EndpointService) DeviceAuth(serialno string) (bool, interface{}) {

	dbErr, dbData := dbaccess.ValidateDevice(repo.Instance().Context.Master.DBConn, serialno)

	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *dbData

	if len(dbRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = constants.MOD_ERR_DEVICE_NOT_FOUND
		return false, errModel
	}

	deviceRecordItem := dbRecord[0]

	if deviceRecordItem.DevState != constants.DB_DEVICE_STATE_ACTIVE {
		errModel := gmodels.APIResponseError{}
		errModel.Code = constants.MOD_ERR_INACTIVE_DEVICE_STATE
		return false, errModel
	}

	dbErr, deviceAuthData := dbaccess.GetDeviceAuthInfo(repo.Instance().Context.Master.DBConn, deviceRecordItem.DevId)

	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "DB Error occured while login.", dbErr)
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbdeviceAuthRecord := *deviceAuthData

	if len(dbdeviceAuthRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = constants.MOD_ERR_DEVICE_CUSTOMER_PRODUCT_MAPPING
		return false, errModel
	}

	deviceAuthRecordItem := dbdeviceAuthRecord[0]

	deviceInfoModel := gmodels.DeviceInfoModel{}
	deviceInfoModel.CpmID = deviceAuthRecordItem.CpmID
	deviceInfoModel.DevID = deviceRecordItem.DevId
	deviceInfoModel.LocationUrl = deviceAuthRecordItem.ServerAddress

	isSuccess, token := lhelper.CacheSetDeviceInfo(repo.Instance().Context, &deviceInfoModel)
	if isSuccess == false {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_SERVER
		return false, errModel
	}

	devAuthResponse := lmodels.APIDeviceAuthResponse{}
	devAuthResponse.Token = token
	devAuthResponse.LocationUrl = deviceAuthRecordItem.ServerAddress

	return true, devAuthResponse
}
