package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gmodels "opensoach.com/models"
)

func AuthorizationHandler(c *gin.Context) {

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

	switch c.Request.URL.Path {

	default:
		isSuccess, successErrorData = AuthorizationService.ValidateUserAuthorization(AuthorizationService{}, c)
		break
	}

	return isSuccess, successErrorData
}
