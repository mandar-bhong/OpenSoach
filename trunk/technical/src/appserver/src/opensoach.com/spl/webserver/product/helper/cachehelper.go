package helper

import (
	"opensoach.com/core"
	ghelper "opensoach.com/core/helper"
	lmodels "opensoach.com/spl/models"
)

func CacheGetCPMDetails(rediscache core.RedisContext, key string) (bool, *lmodels.DBProductBriefRowModel) {

	isSuccess, data := rediscache.Get(key)
	datamodel := &lmodels.DBProductBriefRowModel{}

	if isSuccess {
		isJSONSuccess := ghelper.ConvertFromJSONString(data, datamodel)

		if isJSONSuccess {
			return true, datamodel
		}
	}

	return false, nil
}

func CacheSetCPMDetails(key string, data *lmodels.DBProductBriefRowModel, rediscache core.RedisContext) bool {

	isSuccess, jsonData := ghelper.ConvertToJSON(data)

	if !isSuccess {
		return false
	}

	isSuccess = rediscache.Set(key, jsonData, -1) //Infinite storage

	return isSuccess
}
