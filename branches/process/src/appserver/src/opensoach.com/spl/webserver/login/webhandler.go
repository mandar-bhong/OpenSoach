package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ghelper "opensoach.com/core/helper"
	gmodels "opensoach.com/models"
	lhelper "opensoach.com/spl/helper"
	lmodels "opensoach.com/spl/models"
	repo "opensoach.com/spl/repository"
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

	responsePayload := gmodels.APIPayloadResponse{}

	ghelper.Block{
		Try: func() {
			isSuccess, successErrorData = requestHandler(pContext)
		},

		Catch: func(e ghelper.Exception) {
			panic(e)
			//logger.Log(helper.MODULENAME, logger.ERROR, "Exception occured while processing websocket data: %#v", e)
			isSuccess = false
			errorData := gmodels.APIResponseError{}
			errorData.Code = gmodels.MOD_OPER_ERR_SERVER
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

	//logger.Debug(helper.MODULE_NAME, "API Request Received: %s", pContext.Request.RequestURI)

	switch pContext.Request.RequestURI {
	case "/login":

		loginReq := lmodels.LoginRequest{}

		if err := pContext.Bind(&loginReq); err != nil {
			return false, resultData
		}

		isSuccess, resultData = LoginService.Login(LoginService{}, loginReq.UserName, loginReq.Password)

		if isSuccess {
			sessionInfo := gmodels.UserSessionInfo{}
			loginResp := resultData.(*lmodels.LoginResponse)
			sessionInfo.UserID = loginResp.UserID
			sessionInfo.UserRoleID = loginResp.Category
			sessionInfo.UserType = loginResp.Category
			isSessionSuccess, token := lhelper.SessionCreate(repo.Instance().Context, &sessionInfo)

			if isSessionSuccess {
				loginResp.Token = token
			} else {
				errModel := gmodels.APIResponseError{}
				errModel.Code = gmodels.MOD_OPER_ERR_SERVER
				resultData = errModel
				isSuccess = false
				return isSuccess, resultData
			}
		}

		break
	}

	return isSuccess, resultData
}
