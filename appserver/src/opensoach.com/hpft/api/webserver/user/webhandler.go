package user

import (
	"github.com/gin-gonic/gin"
	"opensoach.com/core/logger"
	"opensoach.com/hpft/api/constants"
	lhelper "opensoach.com/hpft/api/helper"
	repo "opensoach.com/hpft/api/repository"
	gmodels "opensoach.com/models"
)

func registerRouters(router *gin.RouterGroup) {
	router.GET(constants.API_USER_DOCTOR_USERS_LIST, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
}

func requestHandler(pContext *gin.Context) (bool, interface{}) {
	var resultData interface{}
	isSuccess := false

	logger.Context().WithField("Request: ", pContext.Request.URL.Path).LogDebug(SUB_MODULE_NAME, logger.Normal, "API request received.")

	switch pContext.Request.URL.Path {

	case constants.API_USER_DOCTOR_USERS_LIST:

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionData(repo.Instance().Context, pContext)

		if isPrepareExeSuccess == false {
			return false, successErrorData
		}

		isSuccess, resultData = UserService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.SelectDoctorUsers()

		break

	}

	return isSuccess, resultData
}
