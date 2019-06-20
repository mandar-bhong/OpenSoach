package manager

import (
	gcore "opensoach.com/core"
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	pchelper "opensoach.com/prodcore/helper"
)

var SUB_MODULE_NAME = "ProdCore.Server.Manager"

func GetOnlineDevices(cacheCtx gcore.CacheContext, tokenlistjsonstring string, cpmid int64) (bool, []gmodels.DeviceTokenModel) {

	tokens := []string{}

	isJsonSuccess := ghelper.ConvertFromJSONString(tokenlistjsonstring, &tokens)

	if isJsonSuccess == false {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Unable to convert json device packet ", nil)
		return false, nil
	}

	deviceTokenModelList := []gmodels.DeviceTokenModel{}

	for _, token := range tokens {

		isSuccess, _, jsonData := pchelper.CacheGetDeviceInfo(cacheCtx, token)

		if isSuccess == false {
			logger.Context().WithField("Token", token).Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Unable to get information for provided token")
			return false, nil
		}

		deviceTokenModel := gmodels.DeviceTokenModel{}
		isJsonSuccess := ghelper.ConvertFromJSONString(jsonData, &deviceTokenModel)

		if isJsonSuccess == false {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Unable to convert json device packet ", nil)
			return false, nil
		}

		if deviceTokenModel.CpmID == cpmid {
			deviceTokenModelList = append(deviceTokenModelList, deviceTokenModel)
		}

	}

	return true, deviceTokenModelList

}
