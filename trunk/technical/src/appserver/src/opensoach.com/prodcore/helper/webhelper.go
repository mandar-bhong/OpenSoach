package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	gcore "opensoach.com/core"
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	pcconst "opensoach.com/prodcore/constants"
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

	var dataModel interface{}

	isSessionSuccess, contextType, userDeviceInfo, sharedDeviceInfo := DeviceSessionGet(osContext, ginContext)
	if !isSessionSuccess {
		errorData := gmodels.APIResponseError{}
		errorData.Code = gmodels.MOD_OPER_ERR_INPUT_CLIENT_DATA
		return false, errorData
	}

	switch contextType {
	case pcconst.DEVICE_TYPE_SHARED_DEVICE:
		dataModel = &gmodels.DeviceExecutionContext{}
		dataModel.(*gmodels.DeviceExecutionContext).DeviceSessionInfo = *sharedDeviceInfo
		break
	case pcconst.DEVICE_TYPE_USER_DEVICE:
		dataModel = &gmodels.DeviceUserExecutionContext{}
		dataModel.(*gmodels.DeviceUserExecutionContext).DeviceUserSessionInfo = *userDeviceInfo
		break
	}

	return true, dataModel

}

func PrepareDeviceExecutionReqData(osContext *gcore.Context, ginContext *gin.Context, pClientReq interface{}) (bool, interface{}) {

	var dataModel interface{}

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

	isSessionSuccess, contextType, userDeviceInfo, sharedDeviceInfo := DeviceSessionGet(osContext, ginContext)
	if !isSessionSuccess {
		errorData := gmodels.APIResponseError{}
		errorData.Code = gmodels.MOD_OPER_ERR_INPUT_CLIENT_DATA
		return false, errorData
	}

	isUpdateSuccess := SessionUpdate(osContext, ginContext)

	if !isUpdateSuccess {
		errorData := gmodels.APIResponseError{}
		errorData.Code = gmodels.MOD_OPER_ERR_SERVER
		return false, errorData
	}

	switch contextType {
	case pcconst.DEVICE_TYPE_SHARED_DEVICE:
		dataModel = &gmodels.DeviceExecutionContext{}
		dataModel.(*gmodels.DeviceExecutionContext).DeviceSessionInfo = *sharedDeviceInfo
		dataModel.(*gmodels.DeviceExecutionContext).SessionToken = ginContext.GetHeader(gmodels.SESSION_CLIENT_HEADER_KEY)
		break
	case pcconst.DEVICE_TYPE_USER_DEVICE:
		dataModel = &gmodels.DeviceUserExecutionContext{}
		dataModel.(*gmodels.DeviceUserExecutionContext).DeviceUserSessionInfo = *userDeviceInfo
		dataModel.(*gmodels.DeviceUserExecutionContext).SessionToken = ginContext.GetHeader(gmodels.SESSION_CLIENT_HEADER_KEY)
		break
	}

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
		pContext.DataFromReader(http.StatusOK, int64(len(successData.ByteData)), successData.ContentType, bytes.NewReader(successData.ByteData), map[string]string{})

	} else {
		pContext.Data(http.StatusNotFound, "attachment", nil)
	}
}
