package helper

import (
	"fmt"
	"strings"

	gcore "opensoach.com/core"
	ghelper "opensoach.com/core/helper"
	gmodels "opensoach.com/models"
	pcconst "opensoach.com/prodcore/constants"
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

func CacheGetDeviceInfoData(cacheCtx gcore.CacheContext, cacheKey string) (bool, int, *gmodels.DeviceTokenModel, *gmodels.DeviceUserSessionInfo, string) {

	deviceTokenModel := &gmodels.DeviceTokenModel{}
	userTokenInfo := &gmodels.DeviceUserSessionInfo{}
	var contextType int

	isSuccess, jsonData := cacheCtx.Get(cacheKey)

	if isSuccess == false {
		return false, pcconst.DEVICE_TYPE_NONE, nil, nil, ""
	}

	if strings.HasPrefix(cacheKey, pcconst.SHARED_DEVICE_TOKEN_PREFIX) {
		contextType = pcconst.DEVICE_TYPE_SHARED_DEVICE
		if isJsonSuccess := ghelper.ConvertFromJSONString(jsonData, deviceTokenModel); isJsonSuccess == false {
			return false, contextType, nil, nil, ""
		}
	} else if strings.HasPrefix(cacheKey, pcconst.USER_DEVICE_TOKEN_PREFIX) {
		contextType = pcconst.DEVICE_TYPE_USER_DEVICE
		if isJsonSuccess := ghelper.ConvertFromJSONString(jsonData, userTokenInfo); isJsonSuccess == false {
			return false, contextType, nil, nil, ""
		}
	}

	return true, contextType, deviceTokenModel, userTokenInfo, jsonData
}

func CacheSetDeviceConnectionStatus(cacheCtx gcore.CacheContext, deviceID int64, isconnected bool) bool {

	cacheKey := fmt.Sprintf("%s%+v", gmodels.CACHE_KEY_ENTITY_CONNECTION_STATUS, deviceID)

	status := pcconst.DB_DEVICE_CONNECTION_STATE_DISCONNECTED

	if isconnected == true {
		status = pcconst.DB_DEVICE_CONNECTION_STATE_CONNECTED
	}

	return cacheCtx.Set(cacheKey, status, 0)
}
