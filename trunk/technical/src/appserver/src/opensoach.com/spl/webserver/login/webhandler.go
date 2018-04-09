package login

import (
	
	"github.com/gin-gonic/gin"
	
	gmodels "opensoach.com/models"
	"opensoach.com/spl/constants"
	lhelper "opensoach.com/spl/helper"
	lmodels "opensoach.com/spl/models"
)

func registerRouters(router *gin.RouterGroup) {

	//logger.Instance.Debug("Registering log module")

	router.POST(constants.API_USER_LOGIN, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })

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
