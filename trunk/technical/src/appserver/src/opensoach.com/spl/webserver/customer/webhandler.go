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
	
	router.GET(constants.API_CUSTOMER_MASTER_OSU_DETAILS, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_CUSTOMER_MASTER_CU_DETAILS, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
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

	case constants.API_CUSTOMER_MASTER_OSU_DETAILS:

		recReq := lmodels.RecordIdRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &recReq)

		if isPrepareExeSuccess == false {
			return false, successErrorData
		}

		isSuccess, resultData = CustomerService.GetCustomerDetails(CustomerService{}, recReq.RecId)

		break

	case constants.API_CUSTOMER_MASTER_CU_DETAILS:

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionData(repo.Instance().Context, pContext)

		if isPrepareExeSuccess == false {
			return false, successErrorData
		}

		isSuccess, resultData = CustomerService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.GetCustomerDetails(successErrorData.(*gmodels.ExecutionContext).SessionInfo.CustomerID)

		break
	}

	return isSuccess, resultData
}
