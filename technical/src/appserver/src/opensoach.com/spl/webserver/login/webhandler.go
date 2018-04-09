package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ghelper "opensoach.com/core/helper"
	gmodels "opensoach.com/models"
	"opensoach.com/spl/constants"
	lmodels "opensoach.com/spl/models"
)

func registerRouters(router *gin.RouterGroup) {

	//logger.Instance.Debug("Registering log module")

	router.POST(constants.API_USER_LOGIN, commonHandler)

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

	case constants.API_USER_LOGIN:

		authReq := lmodels.AuthRequest{}

		err := pContext.Bind(&authReq)

		if err != nil {

			errModel := gmodels.APIResponseError{}
			errModel.Code = gmodels.MOD_OPER_ERR_INPUT_CLIENT_DATA
			resultData = errModel
			return false, resultData
		}

		isSuccess, resultData = AuthService.Auth(AuthService{}, authReq.UserName, authReq.Password, authReq.ProdCode)

		break

	}

	return isSuccess, resultData
}
