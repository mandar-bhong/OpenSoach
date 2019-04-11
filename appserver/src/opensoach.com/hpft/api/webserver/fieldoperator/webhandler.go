package fieldoperator

import (
	"github.com/gin-gonic/gin"

	"opensoach.com/core/logger"
	"opensoach.com/hpft/api/constants"
	lhelper "opensoach.com/hpft/api/helper"
	lmodels "opensoach.com/hpft/api/models"
	repo "opensoach.com/hpft/api/repository"
	hktmodels "opensoach.com/hpft/models"
	gmodels "opensoach.com/models"
)

func registerRouters(router *gin.RouterGroup) {
	router.POST(constants.API_FIELD_OPERATOR_ADD, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_FIELD_OPERATOR_INFO_MASTER, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_FIELD_OPERATOR_LIST, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_FIELD_OPERATOR_UPDATE, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_FIELD_OPERATOR_LIST_SHORT, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_FIELD_OPERATOR_ASSOCIATE_SP, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_FIELD_OPERATOR_SP_ASSOCIATION_REMOVE, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_FIELD_OPERATOR_ASSOCIATE_SP_INFO, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
}

func requestHandler(pContext *gin.Context) (bool, interface{}) {
	var resultData interface{}
	isSuccess := false

	logger.Context().WithField("Request: ", pContext.Request.URL.Path).LogDebug(SUB_MODULE_NAME, logger.Normal, "API request received.")

	switch pContext.Request.URL.Path {

	case constants.API_FIELD_OPERATOR_ADD:

		fielOperatorAddReq := lmodels.APIFieldOperatorAddRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &fielOperatorAddReq)

		if isPrepareExeSuccess == false {
			return false, successErrorData
		}

		isSuccess, resultData = FieldoperatorService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.Add(fielOperatorAddReq)

		break

	case constants.API_FIELD_OPERATOR_INFO_MASTER:

		recReq := gmodels.APIRecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = FieldoperatorService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.SelectById(recReq.RecId)

		break

	case constants.API_FIELD_OPERATOR_LIST:

		listReq := gmodels.APIDataListRequest{}
		listReq.Filter = &hktmodels.DBSearchFieldOperatorRequestFilterDataModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &listReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = FieldoperatorService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetFieldOperatorList(listReq)

		break

	case constants.API_FIELD_OPERATOR_UPDATE:

		reqData := &hktmodels.DBFieldOperatorUpdateRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = FieldoperatorService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.Update(reqData)

		break

	case constants.API_FIELD_OPERATOR_LIST_SHORT:

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionData(repo.Instance().Context, pContext)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = FieldoperatorService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.FieldOperatorShortDataList()

		break

	case constants.API_FIELD_OPERATOR_ASSOCIATE_SP:

		reqData := lmodels.APIFopSpAddRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			return false, successErrorData
		}

		isSuccess, resultData = FieldoperatorService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.FopSpAdd(reqData)

		break

	case constants.API_FIELD_OPERATOR_SP_ASSOCIATION_REMOVE:

		reqData := &lmodels.APIFopSpDeleteRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = FieldoperatorService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.FopSpDelete(reqData)

		break

	case constants.API_FIELD_OPERATOR_ASSOCIATE_SP_INFO:

		recReq := gmodels.APIRecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = FieldoperatorService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetFopSpAssociation(recReq.RecId)

		break

	}

	return isSuccess, resultData
}
