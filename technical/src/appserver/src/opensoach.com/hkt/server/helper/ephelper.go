package helper

import (
	"opensoach.com/core"
	ghelper "opensoach.com/core/helper"
	gmodels "opensoach.com/models"
)

func GetEPTokenInfo(cache core.CacheContext, token string) (bool, *gmodels.DeviceTokenModel) {

	isSuccess, data := cache.Get(token)

	if isSuccess == false {
		return false, nil
	}
	tokenInfo := &gmodels.DeviceTokenModel{}

	if isJsonSuccess := ghelper.ConvertFromJSONString(data, tokenInfo); isJsonSuccess == false {
		return false, nil
	}

	return true, tokenInfo

}
