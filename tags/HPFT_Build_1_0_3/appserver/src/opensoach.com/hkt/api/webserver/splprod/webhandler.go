package splprod

import (
	"github.com/gin-gonic/gin"
	"opensoach.com/core/logger"
	"opensoach.com/hkt/api/constants"
	lhelper "opensoach.com/hkt/api/helper"
)

func registerRouters(router *gin.RouterGroup) {
	router.GET(constants.API_SPL_PROD_BASE_URL, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
}

func requestHandler(pContext *gin.Context) (bool, interface{}) {
	var resultData interface{}
	isSuccess := false

	logger.Context().WithField("Request: ", pContext.Request.URL.Path).LogDebug(SUB_MODULE_NAME, logger.Normal, "API request received.")

	switch pContext.Request.URL.Path {

	case constants.API_SPL_PROD_BASE_URL:

		isSuccess, resultData = SplprodService.GetBaseUrl(SplprodService{})

		break

	}

	return isSuccess, resultData
}
