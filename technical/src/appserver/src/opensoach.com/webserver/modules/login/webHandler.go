package login

import (
	"github.com/gin-gonic/gin"
	//"opensoach.com/utility/logger"
	"net/http"

	ghelper "opensoach.com/utility/helper"
	whelper "opensoach.com/webserver/webhelper"
	wmodels "opensoach.com/webserver/webmodels"
)

func registerRouters(router *gin.RouterGroup) {

	//logger.Instance.Debug("Registering log module")

	router.POST("/login", commonHandler)

	return
}

func commonHandler(pContext *gin.Context) {
	var isSuccess bool
	var successErrorData interface{}
	ginRetStatus := http.StatusOK

	responsePayload := wmodels.PayloadResponse{}

	ghelper.Block{
		Try: func() {
			isSuccess, successErrorData = requestHandler(pContext)
		},

		Catch: func(e ghelper.Exception) {
			panic(e)
			//logger.Log(helper.MODULENAME, logger.ERROR, "Exception occured while processing websocket data: %#v", e)
			isSuccess = false
			errorData := wmodels.ResponseError{}
			errorData.Code = whelper.MOD_OPER_ERR_SERVER
			successErrorData = errorData
		},

		Finally: func() {
			//Do something if required
		},
	}.Do()

	responsePayload.Success = isSuccess
	if isSuccess {
		responsePayload.Data = successErrorData
	} else {
		responsePayload.Error = successErrorData
	}

	pContext.JSON(ginRetStatus, responsePayload)

	return
}

func requestHandler(pContext *gin.Context) (bool, interface{}) {
	var resultData interface{}
	isSuccess := false

	switch pContext.Request.RequestURI {
	case "/login":

		loginReq := wmodels.LoginRequest{}

		if err := pContext.Bind(&loginReq); err != nil {
			return false, resultData
		}

		isSuccess, resultData = LoginService.Login(LoginService{}, loginReq.UserName, loginReq.Password)

		if isSuccess {

		}

		//		if retData.(gModels.UserLoginData).BankID != nil {
		//			sessionData.BankID = *retData.(gModels.UserLoginData).BankID
		//		}

		//		isServerSessionSetSuccess, token := ghelper.SessionCreate(pContext, sessionData)
		//		if !isServerSessionSetSuccess {
		//			logger.Log(helper.MODULENAME, logger.ERROR, "Unable to set session data")
		//			errorData := gModels.ResponseError{
		//				Code: ghelper.MOD_OPER_ERR_SERVER,
		//			}
		//			return false, errorData
		//		}

		//		jsonResponse.Token = token

		//		isSuccess = true
		//		resultData = jsonResponse

		break

	}

	return isSuccess, resultData
}
