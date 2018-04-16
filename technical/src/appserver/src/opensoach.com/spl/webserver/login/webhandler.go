package login

import (
	"github.com/gin-gonic/gin"

	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	"opensoach.com/spl/constants"
	lhelper "opensoach.com/spl/helper"
	lmodels "opensoach.com/spl/models"
	repo "opensoach.com/spl/repository"
)

func registerRouters(router *gin.RouterGroup) {

	//logger.Instance.Debug("Registering log module")

	router.POST(constants.API_USER_LOGIN, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_USER_LOGIN_INFO, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_USER_LOGOUT, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_CUSTOMER_LOGIN_INFO, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	return
}

func requestHandler(pContext *gin.Context) (bool, interface{}) {
	var resultData interface{}
	isSuccess := false

	logger.Context().WithField("Request: ", pContext.Request.URL.Path).LogDebug(SUB_MODULE_NAME, logger.Normal, "API request received.")

	switch pContext.Request.URL.Path {

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

	case constants.API_USER_LOGIN_INFO:

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionData(repo.Instance().Context, pContext)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = AuthService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetUserLoginDetails()

		break

	case constants.API_USER_LOGOUT:
		isLogoutSuccess := AuthService.UserLogout(AuthService{}, pContext)
		return isLogoutSuccess, nil

	case constants.API_CUSTOMER_LOGIN_INFO:

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionData(repo.Instance().Context, pContext)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = AuthService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetCustomerLoginDetails()

		break

	}

	return isSuccess, resultData
}
