package corporate

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
	router.GET(constants.API_CORPORATE_OSU_LIST, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_CORPORATE_OSU_LIST_SHORT, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_CORPORATE_OSU_ADD, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_CORPORATE_OSU_UPDATE, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_CORPORATE_OSU_INFO_MASTER, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
}

func requestHandler(pContext *gin.Context) (bool, interface{}) {
	var resultData interface{}
	isSuccess := false

	logger.Context().WithField("Request: ", pContext.Request.URL.Path).LogDebug(SUB_MODULE_NAME, logger.Normal, "API request received.")

	switch pContext.Request.URL.Path {

	case constants.API_CORPORATE_OSU_LIST:

		corpListReq := gmodels.APIDataListRequest{}
		corpListReq.Filter = &lmodels.DBSearchCorpRequestFilterDataModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &corpListReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = CorporateService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetCorpDataList(corpListReq)

		break

	case constants.API_CORPORATE_OSU_LIST_SHORT:

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionData(repo.Instance().Context, pContext)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = CorporateService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetCorpShortDataList()

		break

	case constants.API_CORPORATE_OSU_ADD:

		reqData := &lmodels.DBSplCorpRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = CorporateService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.AddCorp(reqData)

		break

	case constants.API_CORPORATE_OSU_UPDATE:

		reqData := &lmodels.DBSplCorpRowModel{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &reqData)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = CorporateService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.UpdateCorp(reqData)

		break

	case constants.API_CORPORATE_OSU_INFO_MASTER:

		recReq := gmodels.APIRecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = CorporateService.GetCorporateInfo(CorporateService{}, recReq.RecId)

		break

	}

	return isSuccess, resultData
}
