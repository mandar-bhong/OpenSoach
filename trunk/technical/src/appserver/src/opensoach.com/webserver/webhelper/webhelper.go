package webhelper

import (
	"github.com/gin-gonic/gin"
	"opensoach.com/utility/logger"
	wmodels "opensoach.com/webserver/webmodels"
)

func PrepareExecutionData(pContext *gin.Context, pClientReqModel interface{}) (bool, interface{}) {
	dataModel := wmodels.ExecutionContext{}

	err := pContext.Bind(pClientReqModel)
	if err != nil {
		logger.Error("Client data binding error: ", err.Error())
		errorData := wmodels.ResponseError{}
		errorData.Code = MOD_OPER_ERR_INPUT_CLIENT_DATA
		return false, errorData
	}

	dataModel.Request = pClientReqModel

	isSuccess, sessionData := SessionGetData(pContext)
	if !isSuccess {
		errorData := wmodels.ResponseError{}
		errorData.Code = MOD_OPER_UNAUTHORIZED
		return false, errorData
	}
	dataModel.SessionInfo = sessionData

	return true, &dataModel
}

func PrepareSessionData(pContext *gin.Context) (bool, interface{}) {
	dataModel := wmodels.ExecutionContext{}

	isSuccess, sessionData := SessionGetData(pContext)

	if !isSuccess {

		errorData := wmodels.ResponseError{}
		errorData.Code = MOD_OPER_UNAUTHORIZED
		return false, errorData
	}

	dataModel.SessionInfo = sessionData

	return true, &dataModel
}
