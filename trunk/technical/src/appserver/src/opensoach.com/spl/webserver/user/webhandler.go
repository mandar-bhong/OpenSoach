package user

import (
	"github.com/gin-gonic/gin"

	gmodels "opensoach.com/models"
	"opensoach.com/spl/constants"
	lhelper "opensoach.com/spl/helper"
	repo "opensoach.com/spl/repository"
)

func registerRouters(router *gin.RouterGroup) {
	router.POST(constants.API_USER_OSU_ADD, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
}

func requestHandler(pContext *gin.Context) (bool, interface{}) {
	var resultData interface{}
	isSuccess := false

	switch pContext.Request.URL.Path {

	case constants.API_USER_OSU_UPDATE_DETAILS:

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionData(repo.Instance().Context, pContext)

		if isPrepareExeSuccess == false {
			return false, successErrorData
		}

		isSuccess, resultData = UserService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.UpdateUserDetails()

		break

	case constants.API_USER_CU_UPDATE_DETAILS:

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionData(repo.Instance().Context, pContext)

		if isPrepareExeSuccess == false {
			return false, successErrorData
		}

		isSuccess, resultData = UserService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.UpdateUserDetails()

		break

	}

	return isSuccess, resultData
}
