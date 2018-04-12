package customer

import (
	"github.com/gin-gonic/gin"

	gmodels "opensoach.com/models"
	"opensoach.com/spl/constants"
	lhelper "opensoach.com/spl/helper"
	lmodels "opensoach.com/spl/models"
	repo "opensoach.com/spl/repository"
)

func registerRouters(router *gin.RouterGroup) {

	router.GET(constants.API_CUSTOMER_OSU_INFO_MASTER, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_CUSTOMER_CU_INFO_MASTER, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_CUSTOMER_OSU_INFO_DETAILS, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_CUSTOMER_CU_INFO_DETAILS, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_CUSTOMER_OSU_CORPORATE_INFO, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_CUSTOMER_CU_CORPORATE_INFO, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
}

func requestHandler(pContext *gin.Context) (bool, interface{}) {
	var resultData interface{}
	isSuccess := false

	switch pContext.Request.URL.Path {

	case constants.API_CUSTOMER_OSU_UPDATE_DETAILS:

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionData(repo.Instance().Context, pContext)

		if isPrepareExeSuccess == false {
			return false, successErrorData
		}

		isSuccess, resultData = CustomerService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.UpdateCustomerDetails()

		break

	case constants.API_CUSTOMER_CU_UPDATE_DETAILS:

		break

	case constants.API_CUSTOMER_OSU_INFO_MASTER:

		recReq := lmodels.RecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			return false, successErrorData
		}

		isSuccess, resultData = CustomerService.GetCustomerInfo(CustomerService{}, recReq.RecId)

		break

	case constants.API_CUSTOMER_CU_INFO_MASTER:

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionData(repo.Instance().Context, pContext)

		if isPrepareExeSuccess == false {
			return false, successErrorData
		}

		isSuccess, resultData = CustomerService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetCustomerInfo(successErrorData.(*gmodels.ExecutionContext).SessionInfo.CustomerID)

		break

	case constants.API_CUSTOMER_OSU_INFO_DETAILS:

		recReq := lmodels.RecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			return false, successErrorData
		}

		isSuccess, resultData = CustomerService.GetCustomerDetailsInfo(CustomerService{}, recReq.RecId)

		break

	case constants.API_CUSTOMER_CU_INFO_DETAILS:

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionData(repo.Instance().Context, pContext)

		if isPrepareExeSuccess == false {
			return false, successErrorData
		}

		isSuccess, resultData = CustomerService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetCustomerDetailsInfo(successErrorData.(*gmodels.ExecutionContext).SessionInfo.CustomerID)

		break

	case constants.API_CUSTOMER_OSU_CORPORATE_INFO:

		recReq := lmodels.RecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			return false, successErrorData
		}

		isSuccess, resultData = CustomerService.GetCorpInfo(CustomerService{}, recReq.RecId)

		break

	case constants.API_CUSTOMER_CU_CORPORATE_INFO:

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionData(repo.Instance().Context, pContext)

		if isPrepareExeSuccess == false {
			return false, successErrorData
		}

		isSuccess, resultData = CustomerService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetCorpInfo(successErrorData.(*gmodels.ExecutionContext).SessionInfo.CustomerID)

		break

	}

	return isSuccess, resultData
}
