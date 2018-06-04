package endpoint

import (
	"fmt"

	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	hktmodels "opensoach.com/hkt/models"
	"opensoach.com/hkt/server/dbaccess"
	repo "opensoach.com/hkt/server/repository"
	gmodels "opensoach.com/models"
	pchelper "opensoach.com/prodcore/helper"
)

func ProcessDeviceConnected(token string) error {

	logger.Context().WithField("Token", token).LogDebug(SUB_MODULE_NAME, logger.Normal, "Device disconnect task is handled by HKT server")

	isSuccess, _, jsonData := pchelper.CacheGetDeviceInfo(repo.Instance().Context.Master.Cache, token)

	if isSuccess == false {
		logger.Context().WithField("Token", token).Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Unable to get information for provided token")
		return fmt.Errorf("Unable to get information for provided token. Token: %s", token)
	}

	deviceTokenModel := &gmodels.DeviceTokenModel{}

	if isSuccess := ghelper.ConvertFromJSONString(jsonData, deviceTokenModel); isSuccess == false {
		logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Unable to convert json data to deviceTokenModel")
		return nil
	}

	dbErr, instDBConn := dbaccess.GetInstanceDBConn(repo.Instance().Context.Master.DBConn, deviceTokenModel.CpmID)

	if dbErr != nil {
		logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "ProcessDeviceConnected:Unable to get device token. Device is offline. Skipping creation of packet")
		return dbErr
	}

	dbDevStatusConnectionStateUpdateDataModel := hktmodels.DBDevStatusConnectionStateUpdateDataModel{}
	dbDevStatusConnectionStateUpdateDataModel.DevId = deviceTokenModel.DevID
	dbDevStatusConnectionStateUpdateDataModel.ConnectionState = gmodels.ENTITY_CONNECTION_STATUS_CONNECTED
	dbDevStatusConnectionStateUpdateDataModel.ConnectionStateSince = ghelper.GetCurrentTime()

	dberr := dbaccess.EPUpdateDeviceConnectionStatusData(instDBConn, dbDevStatusConnectionStateUpdateDataModel)

	if dberr != nil {
		logger.Context().WithField("Token", token).
			WithField("DeviceConnectionStatusData", dbDevStatusConnectionStateUpdateDataModel).LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while updating device connection state.", dbErr)
	}

	return nil
}

func ProcessDeviceDisConnected(token string) error {

	logger.Context().WithField("Token", token).LogDebug(SUB_MODULE_NAME, logger.Normal, "Device disconnect task is handled by HKT server")

	isSuccess, _, jsonData := pchelper.CacheGetDeviceInfo(repo.Instance().Context.Master.Cache, token)

	if isSuccess == false {
		logger.Context().WithField("Token", token).Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Unable to get information for provided token")
		return fmt.Errorf("Unable to get information for provided token. Token: %s", token)
	}

	deviceTokenModel := &gmodels.DeviceTokenModel{}

	if isSuccess := ghelper.ConvertFromJSONString(jsonData, deviceTokenModel); isSuccess == false {
		logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Unable to convert json data to deviceTokenModel")
		return nil
	}

	dbErr, instDBConn := dbaccess.GetInstanceDBConn(repo.Instance().Context.Master.DBConn, deviceTokenModel.CpmID)

	if dbErr != nil {
		logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "ProcessDeviceConnected:Unable to get device token. Device is offline. Skipping creation of packet")
		return dbErr
	}

	dbDevStatusConnectionStateUpdateDataModel := hktmodels.DBDevStatusConnectionStateUpdateDataModel{}
	dbDevStatusConnectionStateUpdateDataModel.DevId = deviceTokenModel.DevID
	dbDevStatusConnectionStateUpdateDataModel.ConnectionState = gmodels.ENTITY_CONNECTION_STATUS_DISCONNECTED
	dbDevStatusConnectionStateUpdateDataModel.ConnectionStateSince = ghelper.GetCurrentTime()

	dberr := dbaccess.EPUpdateDeviceConnectionStatusData(instDBConn, dbDevStatusConnectionStateUpdateDataModel)

	if dberr != nil {
		logger.Context().WithField("Token", token).
			WithField("DeviceConnectionStatusData", dbDevStatusConnectionStateUpdateDataModel).LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while updating device connection state.", dbErr)
	}

	return nil
}
