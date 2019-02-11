package endpoint

import (
	"github.com/gin-gonic/gin"
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	"opensoach.com/spl/api/constants"
	lhelper "opensoach.com/spl/api/helper"
	lmodels "opensoach.com/spl/api/models"
)

func registerRouters(router *gin.RouterGroup) {
	router.POST(constants.API_ENDPOINT_DEVICE_AUTH, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_ENDPOINT_DEVICE_USER_AUTH, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
}

func requestHandler(pContext *gin.Context) (bool, interface{}) {
	var resultData interface{}
	isSuccess := false

	logger.Context().WithField("Request: ", pContext.Request.URL.Path).LogDebug(SUB_MODULE_NAME, logger.Normal, "API request received.")

	switch pContext.Request.URL.Path {

	case constants.API_ENDPOINT_DEVICE_AUTH:

		devAuthReq := lmodels.APIDeviceAuthRequest{}

		err := pContext.Bind(&devAuthReq)

		if err != nil {
			errModel := gmodels.APIResponseError{}
			errModel.Code = gmodels.MOD_OPER_ERR_INPUT_CLIENT_DATA
			resultData = errModel
			return false, resultData
		}

		isSuccess, resultData = EndpointService.DeviceAuth(EndpointService{}, devAuthReq)

		break

	case constants.API_ENDPOINT_DEVICE_USER_AUTH:

		devAuthReq := lmodels.APIDeviceUserAuthRequest{}

		err := pContext.Bind(&devAuthReq)

		if err != nil {
			errModel := gmodels.APIResponseError{}
			errModel.Code = gmodels.MOD_OPER_ERR_INPUT_CLIENT_DATA
			resultData = errModel
			return false, resultData
		}

		isSuccess, resultData = EndpointService.DeviceUserAuth(EndpointService{}, devAuthReq)

		break

	}

	return isSuccess, resultData
}
