package helper

import (
	"opensoach.com/core"
	ghelper "opensoach.com/core/helper"
	lmodels "opensoach.com/spl/models"
)

func CacheGetCPMDetails(rediscache core.CacheContext, key string) (bool, *lmodels.DBProductBriefRowModel) {

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

func CacheSetCPMDetails(key string, data *lmodels.DBProductBriefRowModel, cacheCtx core.CacheContext) bool {

	isSuccess, jsonData := ghelper.ConvertToJSON(data)

	if !isSuccess {
		return false
	}

	isSuccess = cacheCtx.Set(key, jsonData, -1) //Infinite storage

	return isSuccess
}
