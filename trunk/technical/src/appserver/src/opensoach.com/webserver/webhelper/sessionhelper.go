package webhelper

import (
	"crypto/rand"
	"fmt"

	"github.com/gin-gonic/gin"
	ghelper "opensoach.com/utility/helper"
	memsession "opensoach.com/webserver/session"
	wmodels "opensoach.com/webserver/webmodels"
)

func SessionGetData(pContext *gin.Context) (bool, *wmodels.UserSessionInfo) {
	sessionData := wmodels.UserSessionInfo{}
	token := pContext.GetHeader(SESSION_KEY)
	//logger.Log(MODULENAME, logger.DEBUG, "header token: %s", token)

	isSessionGetSuccess, sessionDataJSON := memsession.Get(token)

	if !isSessionGetSuccess {
		//logger.Log(MODULENAME, logger.ERROR, "SessionGetData : Unable to get session for token: %s", token)
		return false, nil
	}

	isConvertFromJSONSuccess := ghelper.ConvertFromJSONString(sessionDataJSON, &sessionData)

	if !isConvertFromJSONSuccess {
		//logger.Log(MODULENAME, logger.ERROR, "SessionGetData : Unable to convert session data from JSON string")
	}

	return isSessionGetSuccess, &sessionData
}

func SessionCreate(pContext *gin.Context, pSessionData *wmodels.UserSessionInfo) (bool, string) {
	isTokenCreateSuccess, sessionToken := createSessionToken()
	if !isTokenCreateSuccess {
		//logger.Log(MODULENAME, logger.ERROR, "SessionCreate : Unable to create session token")
		return false, ""
	}

	//logger.Log(MODULENAME, logger.DEBUG, "session-token: %s", sessionToken)

	isSessionCreateSuccess := createSession(pContext, pSessionData, sessionToken)

	if !isSessionCreateSuccess {
		//logger.Log(MODULENAME, logger.ERROR, "SessionCreate : Unable to create session")
	}

	return true, sessionToken
}

func SessionUpdateExpiration(pContext *gin.Context) bool {

	sessionToken := pContext.GetHeader(SESSION_KEY)

	isSessionGetSuccess, sessionDataJSON := memsession.Get(sessionToken)

	if !isSessionGetSuccess {
		//logger.Log(MODULENAME, logger.ERROR, "SessionUpdateExpiration : Unable to get session data from cache")
		return false
	}

	isSessionUpdateSuccess := memsession.Replace(sessionToken, sessionDataJSON, SESSION_TIME_OUT)

	if !isSessionUpdateSuccess {
		//logger.Log(MODULENAME, logger.ERROR, "Unable to replace session data. Creating new session with same token")
		memsession.Set(sessionToken, sessionDataJSON, SESSION_TIME_OUT)
	}

	return true
}

func SessionDelete(pContext *gin.Context) bool {
	token := pContext.GetHeader(SESSION_KEY)
	memsession.DeleteKey(token)
	return true
}

func createSession(pContext *gin.Context, sessionData *wmodels.UserSessionInfo, sessionToken string) bool {

	isJsonSuccess, jsonData := ghelper.ConvertToJSON(sessionData)

	if !isJsonSuccess {
		//logger.Log(MODULENAME, logger.ERROR, "createSession : Unable to convert session data to JSON")
		return false
	}

	memsession.Set(sessionToken, jsonData, SESSION_TIME_OUT)

	return true
}

func createSessionToken() (bool, string) {
	b := make([]byte, 16)
	_, err := rand.Read(b)

	if err != nil {
		//logger.Log(MODULENAME, logger.ERROR, "createSessionToken:Unable to create session token. Error: %s", err.Error())
		return false, ""
	}

	uuid := fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return true, uuid
}
