package fieldoperator

import (
	"github.com/gin-gonic/gin"

	"opensoach.com/core/logger"
	"opensoach.com/hkt/api/constants"
	lhelper "opensoach.com/hkt/api/helper"
	lmodels "opensoach.com/hkt/api/models"
	repo "opensoach.com/hkt/api/repository"
	hktmodels "opensoach.com/hkt/models"
	gmodels "opensoach.com/models"
)

func registerRouters(router *gin.RouterGroup) {
	router.POST(constants.API_FOP_ADD, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_FIELD_OPERATOR_INFO_MASTER, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_FIELD_OPERATOR_LIST, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
}

func requestHandler(pContext *gin.Context) (bool, interface{}) {
	var resultData interface{}
	isSuccess := false

	logger.Context().WithField("Request: ", pContext.Request.URL.Path).LogDebug(SUB_MODULE_NAME, logger.Normal, "API request received.")

	switch pContext.Request.URL.Path {

	case constants.API_FOP_ADD:
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

		recReq := gmodels.RecordIdRequest{}

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

		listReq := gmodels.DataListRequest{}
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

	}

	return isSuccess, resultData
}
