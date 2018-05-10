package helper

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	gcore "opensoach.com/core"
	ghelper "opensoach.com/core/helper"
	gmodels "opensoach.com/models"
)

type RequestHandler func(pContext *gin.Context) (isSuccess bool, successErrorData interface{})

func PrepareExecutionData(osContext *gcore.Context, ginContext *gin.Context) (bool, interface{}) {
	dataModel := &gmodels.ExecutionContext{}

	isSessionSuccess, userInfo := SessionGet(osContext, ginContext)

	if !isSessionSuccess {
		errorData := gmodels.APIResponseError{}
		errorData.Code = gmodels.MOD_OPER_ERR_INPUT_CLIENT_DATA
		return false, errorData
	}

	dataModel.SessionInfo = *userInfo

	return true, dataModel

}

func PrepareExecutionReqData(osContext *gcore.Context, ginContext *gin.Context, pClientReq interface{}) (bool, interface{}) {

	dataModel := &gmodels.ExecutionContext{}

	if ginContext.Request.Method == http.MethodGet {

		jsonData := ginContext.Query("params")

		if jsonData == "" { // Expected Data but no data received
			errorData := gmodels.APIResponseError{}
			errorData.Code = gmodels.MOD_OPER_ERR_INPUT_CLIENT_DATA
			return false, errorData
		}

		jsonDecodeErr := json.Unmarshal([]byte(jsonData), pClientReq)

		if jsonDecodeErr != nil {
			errorData := gmodels.APIResponseError{}
			errorData.Code = gmodels.MOD_OPER_ERR_INPUT_CLIENT_DATA
			return false, errorData
		}

	} else {

		err := ginContext.Bind(pClientReq)

		if err != nil {
			//logger.Log(MODULENAME, logger.ERROR, "Client data binding error: ", err.Error())
			errorData := gmodels.APIResponseError{}
			errorData.Code = gmodels.MOD_OPER_ERR_INPUT_CLIENT_DATA
			return false, errorData
		}
	}

	isSessionSuccess, userInfo := SessionGet(osContext, ginContext)

	if !isSessionSuccess {
		errorData := gmodels.APIResponseError{}
		errorData.Code = gmodels.MOD_OPER_ERR_USER_SESSION_NOT_AVAILABLE
		return false, errorData
	}

	isUpdateSuccess := SessionUpdate(osContext, ginContext)

	if !isUpdateSuccess {
		errorData := gmodels.APIResponseError{}
		errorData.Code = gmodels.MOD_OPER_ERR_SERVER
		return false, errorData
	}

	dataModel.SessionInfo = *userInfo

	dataModel.SessionToken = ginContext.GetHeader(gmodels.SESSION_CLIENT_HEADER_KEY)

	return true, dataModel

}

func CommonWebRequestHandler(pContext *gin.Context, requestHandlerFunc RequestHandler) {
	var isSuccess bool
	var successErrorData interface{}
	ginRetStatus := http.StatusOK

	responsePayload := gmodels.APIPayloadResponse{}

	ghelper.Block{
		Try: func() {
			isSuccess, successErrorData = requestHandlerFunc(pContext)
		},

		Catch: func(e ghelper.Exception) {
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
