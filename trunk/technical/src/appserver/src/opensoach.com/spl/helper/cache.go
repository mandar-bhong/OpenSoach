package helper

import (
	"fmt"
	"strconv"

	gcore "opensoach.com/core"
	ghelper "opensoach.com/core/helper"
	gmodels "opensoach.com/models"
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

func CacheSetDeviceInfo(osContext *gcore.Context, DevInfoModel *gmodels.DeviceTokenModel) (bool, string) {

	isTokenCreateSuccess, DevAuthCacheToken := ghelper.CreateToken()
	if !isTokenCreateSuccess {
		fmt.Println("Error occured while creating token")
		return false, ""
	}

	isJsonSuccess, jsonData := ghelper.ConvertToJSON(DevInfoModel)

	if !isJsonSuccess {
		return false, ""
	}

	isSetSuccess := osContext.Master.Cache.Set(DevAuthCacheToken, jsonData, 0)

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
