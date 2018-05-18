package processor

import (
	"fmt"

	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	"opensoach.com/splserver/dbaccess"
	repo "opensoach.com/splserver/repository"
)

func EndPointHandlerOnConnection(msg string) error {

	deviceTokenModel := &gmodels.DeviceTokenModel{}

	if isSuccess := ghelper.ConvertFromJSONString(msg, deviceTokenModel); isSuccess == false {
		logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Unable to convert json data to deviceTokenModel")
		//Retries are skipped by returning nil
		return nil
	}

	dbErr := dbaccess.UpdateEPConnectionState(repo.Instance().Context.Master.DBConn, deviceTokenModel.DevID, gmodels.ENTITY_CONNECTION_STATUS_CONNECTED)

	if dbErr != nil {
		logger.Context().WithField("DeviceID", deviceTokenModel.DevID).LogError(SUB_MODULE_NAME, logger.Normal, "Unable to update connection status of device", dbErr)
		return fmt.Errorf("Unable to update connection status of device. DeviceID %+v", deviceTokenModel.DevID)
	}

	return nil
}

func EndPointHandlerOnDisConnection(msg string) error {

	deviceTokenModel := &gmodels.DeviceTokenModel{}

	if isSuccess := ghelper.ConvertFromJSONString(msg, deviceTokenModel); isSuccess == false {
		logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Unable to convert json data to deviceTokenModel")
		//Retries are skipped by returning nil
		return nil
	}

	dbErr := dbaccess.UpdateEPConnectionState(repo.Instance().Context.Master.DBConn, deviceTokenModel.DevID, gmodels.ENTITY_CONNECTION_STATUS_DISCONNECTED)

	if dbErr != nil {
		logger.Context().WithField("DeviceID", deviceTokenModel.DevID).LogError(SUB_MODULE_NAME, logger.Normal, "Unable to update connection status of device", dbErr)
		return fmt.Errorf("Unable to update connection status of device. DeviceID %+v", deviceTokenModel.DevID)
	}

	return nil
}
