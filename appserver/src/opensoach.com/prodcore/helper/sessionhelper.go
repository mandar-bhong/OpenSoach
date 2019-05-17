package helper

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	gcore "opensoach.com/core"
	ghelper "opensoach.com/core/helper"

	//gcache "opensoach.com/core/manager/cache"
	gmodels "opensoach.com/models"
	pcconst "opensoach.com/prodcore/constants"
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

func DeviceSessionGet(osContext *gcore.Context, ginContext *gin.Context) (bool, int, *gmodels.DeviceUserSessionInfo, *gmodels.DeviceTokenModel) {

	userDeviceInfo := &gmodels.DeviceUserSessionInfo{}
	sharedDeviceInfo := &gmodels.DeviceTokenModel{}
	var contextType int

	token := ginContext.GetHeader(gmodels.SESSION_CLIENT_HEADER_KEY)

	isSuccess, jsonData := osContext.Master.Cache.Get(token)
	if !isSuccess {
		return false, contextType, userDeviceInfo, sharedDeviceInfo
	}

	if strings.HasPrefix(token, pcconst.SHARED_DEVICE_TOKEN_PREFIX) {
		contextType = pcconst.DEVICE_TYPE_SHARED_DEVICE
		isJsonConvSuccess := ghelper.ConvertFromJSONString(jsonData, sharedDeviceInfo)
		if !isJsonConvSuccess {
			return false, contextType, userDeviceInfo, sharedDeviceInfo
		}
	} else if strings.HasPrefix(token, pcconst.USER_DEVICE_TOKEN_PREFIX) {
		contextType = pcconst.DEVICE_TYPE_USER_DEVICE
		isJsonConvSuccess := ghelper.ConvertFromJSONString(jsonData, userDeviceInfo)
		if !isJsonConvSuccess {
			return false, contextType, userDeviceInfo, sharedDeviceInfo
		}
	}

	return true, contextType, userDeviceInfo, sharedDeviceInfo
}

func DeviceUserSessionCreate(osContext *gcore.Context, pSessionData *gmodels.DeviceUserSessionInfo) (bool, string) {
	sessionToken := ghelper.GenerateDeviceUserToken()

	isJsonSuccess, jsonData := ghelper.ConvertToJSON(pSessionData)

	if !isJsonSuccess {
		//logger.Log(MODULENAME, logger.ERROR, "createSession : Unable to convert session data to JSON")
		return false, ""
	}

	isSetSuccess := osContext.Master.Cache.Set(sessionToken, jsonData, 0)

	return isSetSuccess, sessionToken
}
