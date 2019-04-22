package middleware

import (
	"net/http"

	"opensoach.com/spl/api/constants"

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

func AuthorizationFilter(reqURL string) (isAuthorizationRequred bool) {
	switch reqURL {
	case constants.API_USER_LOGIN,
		constants.API_USER_LOGOUT,
		constants.API_ENDPOINT_DEVICE_AUTH,
		constants.API_VALIDATE_AUTH_TOKEN,
		constants.API_USER_ACTIVATION,
		constants.API_ENDPOINT_DEVICE_USER_AUTH,
		constants.API_ENDPOINT_DEVICE_USER_LIST,
		constants.API_USER_CREATE_PASSWORD:
		return false
	}

	return true
}

func requestHandler(c *gin.Context) (bool, interface{}) {
	var isSuccess bool
	var successErrorData interface{}

	switch c.Request.URL.Path {

	case constants.API_USER_LOGIN,
		constants.API_USER_LOGOUT,
		constants.API_ENDPOINT_DEVICE_AUTH,
		constants.API_VALIDATE_AUTH_TOKEN:
		return true, nil
		break

	default:
		isSuccess, successErrorData = AuthorizationService.ValidateUserAuthorization(AuthorizationService{}, c)
		break

	}

	return isSuccess, successErrorData
}
