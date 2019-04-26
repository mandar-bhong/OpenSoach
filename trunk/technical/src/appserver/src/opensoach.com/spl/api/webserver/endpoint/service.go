package endpoint

import (
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	pchelper "opensoach.com/prodcore/helper"
	"opensoach.com/spl/api/constants"
	lhelper "opensoach.com/spl/api/helper"
	lmodels "opensoach.com/spl/api/models"
	repo "opensoach.com/spl/api/repository"
	"opensoach.com/spl/api/webserver/endpoint/dbaccess"
)

var SUB_MODULE_NAME = "SPL.Endpoint"

type EndpointService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (EndpointService) DeviceAuth(req lmodels.APIDeviceAuthRequest) (bool, interface{}) {

	dbErr, dbData := dbaccess.ValidateDevice(repo.Instance().Context.Master.DBConn, req.SerialNo)

	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating device.", dbErr)

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

	dbErr, deviceAuthData := dbaccess.GetDeviceAuthInfo(repo.Instance().Context.Master.DBConn, deviceRecordItem.DevId, req.ProductCode)

	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "DB Error occured while getting device auth info.", dbErr)
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

	deviceTokenModel := gmodels.DeviceTokenModel{}
	deviceTokenModel.CpmID = deviceAuthRecordItem.CpmID
	deviceTokenModel.DevID = deviceRecordItem.DevId

	deviceTokenModel.Product = gmodels.ProductInfoModel{}
	deviceTokenModel.Product.NodeDbConn = deviceAuthRecordItem.ConnectionString

	isSuccess, token := lhelper.CacheMapDeviceInfo(repo.Instance().Context, &deviceTokenModel)
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

func (EndpointService) DeviceUserAuth(req lmodels.APIDeviceUserLoginRequest) (bool, interface{}) {

	dbErr, dbData := dbaccess.ValidateDeviceUser(repo.Instance().Context.Master.DBConn, req.UserName, req.Password)

	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating device user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *dbData

	if len(dbRecord) < 1 {
		logger.Context().WithField("request data:", req).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating device user.", nil)
		errModel := gmodels.APIResponseError{}
		errModel.Code = constants.MOD_ERR_LOGIN_INVALID_USER
		return false, errModel
	}

	userRecordItem := dbRecord[0]

	if userRecordItem.UsrState != constants.DB_USER_STATE_ACTIVE {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "User state inactive.", nil)
		errModel := gmodels.APIResponseError{}
		errModel.Code = constants.MOD_ERR_LOGIN_INACTIVE_USER_STATE
		return false, errModel
	}

	dbErr, deviceAuthData := dbaccess.GetDeviceUserAuthInfo(repo.Instance().Context.Master.DBConn, userRecordItem.UserId, req.ProdCode)

	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "DB Error occured while getting device auth info.", dbErr)
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbdeviceAuthRecord := *deviceAuthData

	deviceAuthRecordItem := dbdeviceAuthRecord[0]

	deviceUserSessionInfo := gmodels.DeviceUserSessionInfo{}
	deviceUserSessionInfo.UserID = userRecordItem.UserId
	deviceUserSessionInfo.UserRoleID = deviceAuthRecordItem.UserRoleId

	deviceUserSessionInfo.Product = gmodels.ProductInfoModel{}
	deviceUserSessionInfo.Product.CustProdID = deviceAuthRecordItem.CpmId
	deviceUserSessionInfo.Product.NodeDbConn = deviceAuthRecordItem.Connectionstring

	isSuccess, token := pchelper.DeviceUserSessionCreate(repo.Instance().Context, &deviceUserSessionInfo)
	if isSuccess == false {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error while creating session", nil)
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_SERVER
		return false, errModel
	}

	loginResp := lmodels.APIDeviceUserLoginResponse{}
	loginResp.Token = token
	loginResp.LocationUrl = deviceAuthRecordItem.ServerAddress
	loginResp.UserID = userRecordItem.UserId
	loginResp.UserRoleID = deviceAuthRecordItem.UserRoleId

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully login device user")

	return true, loginResp

}

func (EndpointService) DeviceUserList(req lmodels.APIDeviceUserListRequest) (bool, interface{}) {

	isSuccess, jsonData := repo.Instance().Context.Master.Cache.Get(req.DeviceToken)
	if isSuccess == false {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_USER_TOKEN_NOT_AVAILABLE
		return isSuccess, errModel
	}

	deviceTokenModel := gmodels.DeviceTokenModel{}
	isJsonSuccess := ghelper.ConvertFromJSONString(jsonData, &deviceTokenModel)

	if isJsonSuccess == false {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Unable to convert json device packet ", nil)
		return false, nil
	}

	dbErr, dbData := dbaccess.GetDeviceUserListData(repo.Instance().Context.Master.DBConn, deviceTokenModel.CpmID)

	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating device.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *dbData

	return true, dbRecord

}
