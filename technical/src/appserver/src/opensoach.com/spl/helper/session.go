package helper

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	gcore "opensoach.com/core"
	ghelper "opensoach.com/core/helper"
	//gcache "opensoach.com/core/manager/cache"
	gmodels "opensoach.com/models"
)

func SessionCreate(osContext *gcore.Context, pSessionData *gmodels.UserSessionInfo) (bool, string) {
	isTokenCreateSuccess, sessionToken := ghelper.CreateToken()
	if !isTokenCreateSuccess {
		//logger.Log(MODULENAME, logger.ERROR, "SessionCreate : Unable to create session token")
		fmt.Println("Error occured while creating token")
		return false, ""
	}

	isJsonSuccess, jsonData := ghelper.ConvertToJSON(pSessionData)

	if !isJsonSuccess {
		//logger.Log(MODULENAME, logger.ERROR, "createSession : Unable to convert session data to JSON")
		return false, ""
	}

	osContext.Master.Cache.Set(sessionToken, jsonData, time.Minute*20)

	return true, sessionToken
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
	return osContext.Master.Cache.Update(token, time.Minute*2)
}

func SessionDelete(osContext *gcore.Context, ginContext *gin.Context) bool {
	token := ginContext.GetHeader(gmodels.SESSION_CLIENT_HEADER_KEY)
	return osContext.Master.Cache.Remove(token)
}
