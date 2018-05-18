package helper

import (
	"fmt"

	gcore "opensoach.com/core"
	ghelper "opensoach.com/core/helper"
	gmodels "opensoach.com/models"
)

func CacheGetDeviceInfo(cacheCtx gcore.CacheContext, cacheKey string) (bool, *gmodels.DeviceTokenModel, string) {

	deviceTokenModel := &gmodels.DeviceTokenModel{}

	isSuccess, jsonData := cacheCtx.Get(cacheKey)

	if isSuccess == false {
		return false, nil, ""
	}

	isJsonSuccess := ghelper.ConvertFromJSONString(jsonData, deviceTokenModel)

	if isJsonSuccess == false {
		//log
		return false, nil, ""
	}

	return true, deviceTokenModel, jsonData
}

func CacheSetDeviceConnectionStatus(cacheCtx gcore.CacheContext, deviceID int64, isconnected bool) bool {

	cacheKey := fmt.Sprintf("%s%+v", gmodels.CACHE_KEY_ENTITY_CONNECTION_STATUS, deviceID)

	status := gmodels.ENTITY_CONNECTION_STATUS_DISCONNECTED

	if isconnected == true {
		status = gmodels.ENTITY_CONNECTION_STATUS_CONNECTED
	}

	return cacheCtx.Set(cacheKey, status, 0)
}
