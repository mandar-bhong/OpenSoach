package helper

import (
	"time"

	"github.com/gin-gonic/gin"
	gcore "opensoach.com/core"
	ghelper "opensoach.com/core/helper"

	//gcache "opensoach.com/core/manager/cache"
	gmodels "opensoach.com/models"
)

var sessionTimeOutMin int

func init() {
	sessionTimeOutMin = 20
}

func SessionSetTimeOut(timeoutMin int) {
	sessionTimeOutMin = timeoutMin
}

func SessionCreate(osContext *gcore.Context, pSessionData *gmodels.UserSessionInfo) (bool, string) {
	sessionToken := ghelper.GenerateAPIToken()

	isJsonSuccess, jsonData := ghelper.ConvertToJSON(pSessionData)

	if !isJsonSuccess {
		//logger.Log(MODULENAME, logger.ERROR, "createSession : Unable to convert session data to JSON")
		return false, ""
	}

	isSetSuccess := osContext.Master.Cache.Set(sessionToken, jsonData, time.Minute*time.Duration(sessionTimeOutMin))

	return isSetSuccess, sessionToken
}

func SessionGet(osContext *gcore.Context, ginContext *gin.Context) (bool, *gmodels.UserSessionInfo) {

	userInfo := &gmodels.UserSessionInfo{}
	token := ginContext.GetHeader(gmodels.SESSION_CLIENT_HEADER_KEY)

	isSuccess, jsonData := osContext.Master.Cache.Get(token)

	if !isSuccess {
		return false, nil
	}

	isJsonConvSuccess := ghelper.ConvertFromJSONString(jsonData, userInfo)

	if !isJsonConvSuccess {
		return false, nil
	}

	return true, userInfo
}

func SessionUpdate(osContext *gcore.Context, ginContext *gin.Context) bool {
	token := ginContext.GetHeader(gmodels.SESSION_CLIENT_HEADER_KEY)
	return osContext.Master.Cache.Update(token, time.Minute*time.Duration(sessionTimeOutMin))
}

func SessionDelete(osContext *gcore.Context, ginContext *gin.Context) bool {
	token := ginContext.GetHeader(gmodels.SESSION_CLIENT_HEADER_KEY)
	return osContext.Master.Cache.Remove(token)
}

func DeviceSessionGet(osContext *gcore.Context, ginContext *gin.Context) (bool, *gmodels.DeviceTokenModel) {

	deviceInfo := &gmodels.DeviceTokenModel{}
	token := ginContext.GetHeader(gmodels.SESSION_CLIENT_HEADER_KEY)

	isSuccess, jsonData := osContext.Master.Cache.Get(token)

	if !isSuccess {
		return false, nil
	}

	isJsonConvSuccess := ghelper.ConvertFromJSONString(jsonData, deviceInfo)

	if !isJsonConvSuccess {
		return false, nil
	}

	return true, deviceInfo
}

func DeviceUserSessionCreate(osContext *gcore.Context, pSessionData *gmodels.DeviceUserSessionInfo) (bool, string) {
	sessionToken := ghelper.GenerateDeviceUserToken()

	isJsonSuccess, jsonData := ghelper.ConvertToJSON(pSessionData)

	if !isJsonSuccess {
		//logger.Log(MODULENAME, logger.ERROR, "createSession : Unable to convert session data to JSON")
		return false, ""
	}

	isSetSuccess := osContext.Master.Cache.Set(sessionToken, jsonData, time.Minute*time.Duration(sessionTimeOutMin))

	return isSetSuccess, sessionToken
}

func DeviceUserSessionGet(osContext *gcore.Context, ginContext *gin.Context) (bool, *gmodels.DeviceUserSessionInfo) {

	deviceUserInfo := &gmodels.DeviceUserSessionInfo{}
	token := ginContext.GetHeader(gmodels.SESSION_CLIENT_HEADER_KEY)

	isSuccess, jsonData := osContext.Master.Cache.Get(token)

	if !isSuccess {
		return false, nil
	}

	isJsonConvSuccess := ghelper.ConvertFromJSONString(jsonData, deviceUserInfo)

	if !isJsonConvSuccess {
		return false, nil
	}

	return true, deviceUserInfo
}
