package customer

import (
	"github.com/gin-gonic/gin"

	gmodels "opensoach.com/models"
	"opensoach.com/spl/constants"
	lhelper "opensoach.com/spl/helper"
	lmodels "opensoach.com/spl/models"
)

func registerRouters(router *gin.RouterGroup) {
	//Register
}

func requestHandler(pContext *gin.Context) (bool, interface{}) {
	var resultData interface{}
	isSuccess := false

	switch pContext.Request.RequestURI {

	case constants.API_CUSTOMER_OSU_UPDATE_DETAILS:
		customerService := CustomerService{}

		break

	case constants.API_CUSTOMER_CU_UPDATE_DETAILS:
		break
	}

	return isSuccess, resultData
}
