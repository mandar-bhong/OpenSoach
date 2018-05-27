package helper

import (
	"fmt"
	"strconv"

	gcore "opensoach.com/core"
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	pcconst "opensoach.com/prodcore/constants"
)

func CacheSetCPMKey(osContext *gcore.Context, cpmid int64, productInfoModel *gmodels.ProductInfoModel) bool {

	isJsonSuccess, jsonData := ghelper.ConvertToJSON(productInfoModel)

	if !isJsonSuccess {
		return false
	}

	key := gmodels.CACHE_KEY_PREFIX_CPM_ID + strconv.FormatInt(cpmid, 10)

	isSetSuccess := osContext.Master.Cache.Set(key, jsonData, 0)

	return isSetSuccess
}

func CacheGetCPMKey(osContext *gcore.Context, cpmid int64) (bool, *gmodels.ProductInfoModel) {

	productInfoModel := &gmodels.ProductInfoModel{}

	key := gmodels.CACHE_KEY_PREFIX_CPM_ID + strconv.FormatInt(cpmid, 10)

	isSuccess, jsonData := osContext.Master.Cache.Get(key)

	if !isSuccess {
		return false, nil
	}

	isJsonConvSuccess := ghelper.ConvertFromJSONString(jsonData, productInfoModel)

	if !isJsonConvSuccess {
		return false, nil
	}

	return true, productInfoModel

}

func CacheMapDeviceInfo(osContext *gcore.Context, DevInfoModel *gmodels.DeviceTokenModel) (bool, string) {

	DevAuthCacheToken := ghelper.GenerateDeviceToken()

	isJsonSuccess, jsonData := ghelper.ConvertToJSON(DevInfoModel)

	if !isJsonSuccess {
		logger.Context().WithField("Method", "CacheSetDeviceInfo").LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while converting to json data", nil)
		return false, ""
	}

	deviceTokenKey := fmt.Sprintf("%s%d", pcconst.CACHE_DEVICE_TOKEN_MAPPING_KEY_PREFIX, DevInfoModel.DevID)

	isSetSuccess := osContext.Master.Cache.Set(DevAuthCacheToken, jsonData, 0)

	if isSetSuccess == false {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Unable to set Device auth token key into cache", nil)
		return false, ""
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Device Auth token vs auth model set successfully")

	isSetSuccess = osContext.Master.Cache.Set(deviceTokenKey, DevAuthCacheToken)

	if isSetSuccess == false {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "", nil)
		return false, ""
	}

	return isSetSuccess, DevAuthCacheToken
}

func CacheGetDeviceInfo(osContext *gcore.Context, token string) (bool, *gmodels.DeviceTokenModel) {

	DeviceInfoModel := &gmodels.DeviceTokenModel{}

	isSuccess, jsonData := osContext.Master.Cache.Get(token)

	if !isSuccess {
		return false, nil
	}

	isJsonConvSuccess := ghelper.ConvertFromJSONString(jsonData, DeviceInfoModel)

	if !isJsonConvSuccess {
		return false, nil
	}

	return true, DeviceInfoModel

}
