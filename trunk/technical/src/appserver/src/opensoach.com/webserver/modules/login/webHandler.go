package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"opensoach.com/utility/logger"

	ghelper "opensoach.com/utility/helper"
	"opensoach.com/webserver/modules/login/helper"
	whelper "opensoach.com/webserver/webhelper"
	wmodels "opensoach.com/webserver/webmodels"
)

func registerRouters(router *gin.RouterGroup) {

	//logger.Instance.Debug("Registering log module")

	router.POST("/login", commonHandler)
	router.POST("/getproducts", commonHandler)
	router.POST("/selectproduct", commonHandler)

	return
}

func commonHandler(pContext *gin.Context) {
	var isSuccess bool
	var successErrorData interface{}
	ginRetStatus := http.StatusOK

	responsePayload := wmodels.PayloadResponse{}

	ghelper.Block{
		Try: func() {
			isSuccess, successErrorData = requestHandler(pContext)
		},

		Catch: func(e ghelper.Exception) {
			panic(e)
			//logger.Log(helper.MODULENAME, logger.ERROR, "Exception occured while processing websocket data: %#v", e)
			isSuccess = false
			errorData := wmodels.ResponseError{}
			errorData.Code = whelper.MOD_OPER_ERR_SERVER
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

	logger.Debug(helper.MODULE_NAME, "API Request Received: %s", pContext.Request.RequestURI)

	switch pContext.Request.RequestURI {
	case "/login":

		loginReq := wmodels.LoginRequest{}

		if err := pContext.Bind(&loginReq); err != nil {
			return false, resultData
		}

		isSuccess, resultData = LoginService.Login(LoginService{}, loginReq.UserName, loginReq.Password)

		if isSuccess {
			sessionInfo := wmodels.UserSessionInfo{}
			loginResp := resultData.(*wmodels.LoginResponse)
			sessionInfo.UserID = loginResp.UserID
			sessionInfo.UserRoleID = loginResp.Category
			sessionInfo.UserType = loginResp.Category
			isSessionSuccess, token := whelper.SessionCreate(pContext, &sessionInfo)

			if isSessionSuccess {
				loginResp.Token = token
			} else {
				errModel := wmodels.ResponseError{}
				errModel.Code = whelper.MOD_OPER_ERR_SERVER
				resultData = errModel
				isSuccess = false
				return isSuccess, resultData
			}
		}

		break

	case "/getproducts":
		isExecutionDataSuccess, successErrorData := whelper.PrepareSessionData(pContext)

		if !isExecutionDataSuccess {
			return false, successErrorData
		}

		exeContext := successErrorData.(*wmodels.ExecutionContext)
		isSuccess, resultData = ProductService.GetProducts(ProductService{}, exeContext)
		break

	case "/selectproduct":

		isExecutionDataSuccess, successErrorData := whelper.PrepareExecutionData(pContext, &wmodels.APILoginSelectProductRequest{})

		if !isExecutionDataSuccess {
			return false, successErrorData
		}

		exeContext := successErrorData.(*wmodels.ExecutionContext)
		isSuccess, resultData = ProductService.SelectProduct(ProductService{}, exeContext)
		break
	}

	return isSuccess, resultData
}
