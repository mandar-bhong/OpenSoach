package user

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
	router.POST(constants.API_USER_OSU_ADD, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_USER_OSU_UPDATE_DETAILS, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_USER_CU_UPDATE_DETAILS, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_USER_OSU_UPDATE_STATE, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_USER_CU_UPDATE_STATE, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_USER_UPDATE_PASSWORD, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_USER_LIST, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
}

func requestHandler(pContext *gin.Context) (bool, interface{}) {
	var resultData interface{}
	isSuccess := false

	logger.Context().WithField("Request: ", pContext.Request.URL.Path).LogDebug(SUB_MODULE_NAME, logger.Normal, "API request received.")

	switch pContext.Request.URL.Path {

	case constants.API_USER_OSU_ADD:

		userReqData := lmodels.DBSplMasterUserTableRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &userReqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = UserService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.AddUser(userReqData)

		break

	case constants.API_USER_OSU_UPDATE_DETAILS:

		usrDetailsReqData := lmodels.DBSplMasterUsrDetailsTableRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &usrDetailsReqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = UserService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.UpdateUserDetails(usrDetailsReqData)

		break

	case constants.API_USER_CU_UPDATE_DETAILS:

		usrDetailsReqData := lmodels.DBSplMasterUsrDetailsTableRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &usrDetailsReqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		usrDetailsReqData.UsrIdFk = successErrorData.(*gmodels.ExecutionContext).SessionInfo.UserID

		isSuccess, resultData = UserService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.UpdateUserDetails(usrDetailsReqData)

		break

	case constants.API_USER_OSU_UPDATE_STATE:

		userReqData := lmodels.DBSplMasterUserTableRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &userReqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = UserService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.UpdateUserState(userReqData)

		break

	case constants.API_USER_CU_UPDATE_STATE:

		userReqData := lmodels.DBSplMasterUserTableRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &userReqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		userReqData.Id = successErrorData.(*gmodels.ExecutionContext).SessionInfo.UserID

		isSuccess, resultData = UserService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.UpdateUserState(userReqData)

		break

	case constants.API_USER_UPDATE_PASSWORD:

		userReqData := lmodels.UpdatePasswordRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &userReqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = UserService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.ChangeUserPassword(userReqData, successErrorData.(*gmodels.ExecutionContext).SessionInfo.UserID)

		break

	case constants.API_USER_LIST:

		userListReq := lmodels.DataListRequest{}
		userListReq.Filter = &lmodels.DBSearchUserRequestFilterDataModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &userListReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = UserService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetUserDataList(userListReq)

		break

	}

	return isSuccess, resultData
}
