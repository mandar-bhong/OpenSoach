package middleware

import (
	"net/http"

	"opensoach.com/spl/constants"

	"github.com/gin-gonic/gin"
	gmodels "opensoach.com/models"
)

func authorizationHandler(c *gin.Context) {

	var isSuccess bool
	var successErrorData interface{}

	responsePayload := gmodels.APIPayloadResponse{}

	isSuccess, successErrorData = requestHandler(c)
	if isSuccess {
		c.Next()
	} else {
		responsePayload.Success = isSuccess
		responsePayload.Error = successErrorData
		c.JSON(http.StatusUnauthorized, responsePayload)
		c.Abort()
	}
}

func requestHandler(c *gin.Context) (bool, interface{}) {
	var isSuccess bool
	var successErrorData interface{}

	switch c.Request.RequestURI {

	case constants.API_USER_LOGIN,
		constants.API_USER_LOGOUT,
		constants.API_VALIDATE_AUTH_TOKEN:
		return true, nil
		break

	default:
		isSuccess, successErrorData = AuthorizationService.ValidateUserAuthorization(AuthorizationService{}, c)
		break

	}

	return isSuccess, successErrorData
}
