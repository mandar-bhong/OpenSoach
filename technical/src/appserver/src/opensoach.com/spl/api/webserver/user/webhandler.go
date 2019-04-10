package user

import (
	"github.com/gin-gonic/gin"
	"opensoach.com/core/logger"

	gmodels "opensoach.com/models"
	"opensoach.com/spl/api/constants"
	lhelper "opensoach.com/spl/api/helper"
	lmodels "opensoach.com/spl/api/models"
	repo "opensoach.com/spl/api/repository"
)

func registerRouters(router *gin.RouterGroup) {
	router.POST(constants.API_USER_OSU_ADD, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_USER_OSU_UPDATE_DETAILS, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_USER_CU_UPDATE_DETAILS, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_USER_OSU_UPDATE_STATE, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_USER_CU_UPDATE_STATE, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_USER_UPDATE_PASSWORD, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_USER_OSU_LIST, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_USER_CU_LIST, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_USER_OSU_ASSOCIATE_USER_WITH_CUST, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_USER_CU_ASSOCIATE_USER_WITH_CUST, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_USER_OSU_INFO_DETAILS, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_USER_CU_INFO_DETAILS, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_UROLE_OSU_LIST, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_UROLE_LIST, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_USER_PRODUCT_ASSCOCIATION_OSU_LIST, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_USER_PRODUCT_ASSCOCIATION_OSU_UPDATE, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_USER_OSU_UPDATE, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_USER_OSU_INFO_MASTER, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_USER_CU_INFO_MASTER, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_USER_CU_ROLE_ADD, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_USER_CU_ROLE_INFO_MASTER, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_USER_CU_ROLE_INFO_DETAILS, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_USER_CU_ROLE_UPDATE, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_USER_CU_ROLE_UPDATE_DETAILS, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_USER_ACTIVATION, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
}

func requestHandler(pContext *gin.Context) (bool, interface{}) {
	var resultData interface{}
	isSuccess := false

	logger.Context().WithField("Request: ", pContext.Request.URL.Path).LogDebug(SUB_MODULE_NAME, logger.Normal, "API request received.")

	switch pContext.Request.URL.Path {

	case constants.API_USER_OSU_ADD:

		userReqData := lmodels.DBSplMasterUserRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &userReqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = UserService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.AddUser(userReqData)

		break

	case constants.API_USER_CU_ROLE_ADD:

		userReqData := lmodels.DBSplMasterUserRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &userReqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = UserService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.AddCUUser(userReqData)

		break

	case constants.API_USER_OSU_UPDATE_DETAILS:

		usrDetailsReqData := lmodels.DBSplMasterUsrDetailsRowModel{}

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

		usrDetailsReqData := lmodels.DBSplMasterUsrDetailsRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &usrDetailsReqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		usrDetailsReqData.UsrId = successErrorData.(*gmodels.ExecutionContext).SessionInfo.UserID

		isSuccess, resultData = UserService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.UpdateUserDetails(usrDetailsReqData)

		break

	case constants.API_USER_OSU_UPDATE_STATE:

		userReqData := lmodels.DBSplMasterUserRowModel{}

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

		userReqData := lmodels.DBSplMasterUserRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &userReqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		userReqData.UsrId = successErrorData.(*gmodels.ExecutionContext).SessionInfo.UserID

		isSuccess, resultData = UserService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.UpdateUserState(userReqData)

		break

	case constants.API_USER_UPDATE_PASSWORD:

		userReqData := lmodels.APIUpdatePasswordRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &userReqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = UserService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.ChangeUserPassword(userReqData, successErrorData.(*gmodels.ExecutionContext).SessionInfo.UserID)

		break

	case constants.API_USER_OSU_INFO_DETAILS:

		recReq := gmodels.APIRecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = UserService.GetUserDetailsInfo(UserService{}, recReq.RecId)

		break

	case constants.API_USER_CU_INFO_DETAILS:

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionData(repo.Instance().Context, pContext)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = UserService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetUserDetailsInfo(successErrorData.(*gmodels.ExecutionContext).SessionInfo.UserID)

		break

	case constants.API_USER_OSU_LIST:

		userListReq := gmodels.APIDataListRequest{}
		userListReq.Filter = &lmodels.DBSearchUserRequestFilterDataModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &userListReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = UserService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetOSUDataList(userListReq)

		break

	case constants.API_USER_CU_LIST:

		userListReq := gmodels.APIDataListRequest{}
		userListReq.Filter = &lmodels.DBSearchUserRequestFilterDataModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &userListReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		userListReq.Filter.(*lmodels.DBSearchUserRequestFilterDataModel).CpmId = &successErrorData.(*gmodels.ExecutionContext).SessionInfo.Product.CustProdID

		isSuccess, resultData = UserService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetCUDataList(userListReq)

		break

	case constants.API_USER_OSU_ASSOCIATE_USER_WITH_CUST:

		reqData := &lmodels.APICustomerAssociateUserRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = UserService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.AssociateUserWithCust(reqData)

		break

	case constants.API_USER_CU_ASSOCIATE_USER_WITH_CUST:

		reqData := &lmodels.APICustomerAssociateUserRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		reqData.CpmId = successErrorData.(*gmodels.ExecutionContext).SessionInfo.Product.CustProdID

		isSuccess, resultData = UserService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.AssociateUserWithCust(reqData)

		break

	case constants.API_UROLE_OSU_LIST:

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionData(repo.Instance().Context, pContext)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = UserService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetUserRoleListOSU()

		break

	case constants.API_UROLE_LIST:

		uroleReq := lmodels.APIUroleRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &uroleReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = UserService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetUserRoleList(uroleReq.Prodcode)

		break

	case constants.API_USER_PRODUCT_ASSCOCIATION_OSU_LIST:

		recReq := gmodels.APIRecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = UserService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetUserProdAssociation(recReq.RecId)

		break

	case constants.API_USER_PRODUCT_ASSCOCIATION_OSU_UPDATE:

		reqData := &lmodels.DBUsrCpmStateUpdateRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = UserService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.UpdateUcpmState(reqData)

		break

	case constants.API_USER_OSU_UPDATE:

		reqData := &lmodels.DBUserUpdateRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = UserService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.UpdateUser(reqData)

		break

	case constants.API_USER_OSU_INFO_MASTER:

		recReq := gmodels.APIRecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = UserService.GetUserInfo(UserService{}, recReq.RecId)

		break

	case constants.API_USER_CU_INFO_MASTER:

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionData(repo.Instance().Context, pContext)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = UserService.GetCUUserInfo(UserService{}, successErrorData.(*gmodels.ExecutionContext).SessionInfo.UserID)

		break

	case constants.API_USER_CU_ROLE_INFO_MASTER:

		recReq := gmodels.APIRecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = UserService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetCUUserInfo(recReq.RecId)

		break

	case constants.API_USER_CU_ROLE_INFO_DETAILS:

		recReq := gmodels.APIRecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = UserService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetUserDetailsInfo(recReq.RecId)

		break

	case constants.API_USER_CU_ROLE_UPDATE:

		reqData := &lmodels.APICUUserUpdateRequestModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = UserService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.UpdateCUUser(reqData)

		break

	case constants.API_USER_CU_ROLE_UPDATE_DETAILS:

		usrDetailsReqData := lmodels.DBSplMasterUsrDetailsRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &usrDetailsReqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = UserService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.UpdateUserDetails(usrDetailsReqData)

		break

	case constants.API_USER_ACTIVATION:

		reqData := lmodels.APIUserActivateRequestModel{}

		err := pContext.Bind(&reqData)

		if err != nil {
			errModel := gmodels.APIResponseError{}
			errModel.Code = gmodels.MOD_OPER_ERR_INPUT_CLIENT_DATA
			resultData = errModel
			return false, resultData
		}

		isSuccess, resultData = UserService.UserActivation(UserService{}, reqData)

		break

	}

	return isSuccess, resultData
}
