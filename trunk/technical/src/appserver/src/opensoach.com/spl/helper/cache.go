package helper

import (
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
