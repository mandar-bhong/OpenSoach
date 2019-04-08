package helper

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	gcore "opensoach.com/core"
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	"opensoach.com/prodcore/models"
)

var SUB_MODULE_NAME = "ProdCore.Helper"

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

func PrepareDeviceExecutionData(osContext *gcore.Context, ginContext *gin.Context) (bool, interface{}) {
	dataModel := &gmodels.DeviceExecutionContext{}

	isSessionSuccess, deviceInfo := DeviceSessionGet(osContext, ginContext)

	if !isSessionSuccess {
		errorData := gmodels.APIResponseError{}
		errorData.Code = gmodels.MOD_OPER_ERR_INPUT_CLIENT_DATA
		return false, errorData
	}

	dataModel.DeviceSessionInfo = *deviceInfo

	return true, dataModel

}

func PrepareDeviceExecutionReqData(osContext *gcore.Context, ginContext *gin.Context, pClientReq interface{}) (bool, interface{}) {

	dataModel := &gmodels.DeviceExecutionContext{}

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

	isSessionSuccess, deviceInfo := DeviceSessionGet(osContext, ginContext)

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

	dataModel.DeviceSessionInfo = *deviceInfo
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
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while executing api request", fmt.Errorf("Error:%+v", e))
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

func FileDownloadHandler(pContext *gin.Context, requestHandlerFunc RequestHandler) {
	var isSuccess bool
	var successErrorData interface{}
	var successData models.DocumentData

	ghelper.Block{
		Try: func() {
			isSuccess, successErrorData = requestHandlerFunc(pContext)

			if isSuccess {
				successData = successErrorData.(models.DocumentData)
			}
		},
		Catch: func(e ghelper.Exception) {
			//TODO:
		},
		Finally: func() {
			//Do something if required
		},
	}.Do()

	if isSuccess {
		pContext.Data(http.StatusOK, successData.ContentType, successData.ByteData)

	} else {
		pContext.Data(http.StatusNotFound, "attachment", nil)
	}
}
