package helper

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	gcore "opensoach.com/core"
	gmodels "opensoach.com/models"
)

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
		errorData.Code = gmodels.MOD_OPER_ERR_INPUT_CLIENT_DATA
		return false, errorData
	}

	isUpdateSuccess := SessionUpdate(osContext, ginContext)

	if !isUpdateSuccess {
		errorData := gmodels.APIResponseError{}
		errorData.Code = gmodels.MOD_OPER_ERR_SERVER
		return false, errorData
	}

	dataModel.Request = pClientReq
	dataModel.SessionInfo = *userInfo

	return true, dataModel

}
