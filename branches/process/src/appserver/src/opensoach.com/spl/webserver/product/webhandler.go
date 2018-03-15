package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ghelper "opensoach.com/core/helper"
	gmodels "opensoach.com/models"
	lhelper "opensoach.com/spl/helper"
	lmodels "opensoach.com/spl/models"
	repo "opensoach.com/spl/repository"
)

func registerRouters(router *gin.RouterGroup) {

	//logger.Instance.Debug("Registering product module")

	router.GET("/getproducts", commonHandler)
	router.POST("/selectproduct", commonHandler)

	return
}

func commonHandler(pContext *gin.Context) {
	var isSuccess bool
	var successErrorData interface{}
	ginRetStatus := http.StatusOK

	responsePayload := gmodels.APIPayloadResponse{}

	ghelper.Block{
		Try: func() {
			isSuccess, successErrorData = requestHandler(pContext)
		},

		Catch: func(e ghelper.Exception) {
			panic(e)
			//logger.Log(helper.MODULENAME, logger.ERROR, "Exception occured while processing websocket data: %#v", e)
			isSuccess = false
			errorData := gmodels.APIResponseError{}
			errorData.Code = gmodels.MOD_OPER_ERR_SERVER
			successErrorData = errorData
		},

		Finally: func() {
			//Do something if required
		},
	}.Do()

	responsePayload.Success = isSuccess
	if isSuccess {
		responsePayload.Data = successErrorData
	} else {
		responsePayload.Error = successErrorData
	}

	pContext.JSON(ginRetStatus, responsePayload)

	return
}

func requestHandler(pContext *gin.Context) (bool, interface{}) {
	var resultData interface{}
	isSuccess := false

	switch pContext.Request.RequestURI {
	case "/getproducts":
		isExecutionDataSuccess, successErrorData := lhelper.PrepareExecutionData(repo.Instance().Context, pContext)

		if !isExecutionDataSuccess {
			return false, successErrorData
		}

		exeContext := successErrorData.(*gmodels.ExecutionContext)
		isSuccess, resultData = ProductService.GetProducts(ProductService{}, exeContext)
		break

	case "/selectproduct":

		isExecutionDataSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &lmodels.APIProductSelectRequest{})

		if !isExecutionDataSuccess {
			return false, successErrorData
		}

		exeContext := successErrorData.(*gmodels.ExecutionContext)
		isSuccess, resultData = ProductService.SelectProduct(ProductService{}, exeContext)
		break

	}

	return isSuccess, resultData
}
