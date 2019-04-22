package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gmodels "opensoach.com/models"

	"opensoach.com/core/logger"
)

func AuthorizationHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var isSuccess bool
		var successErrorData interface{}

		logger.Context().LogDebug("Authorization", logger.Instrumentation, "Executing authorization")
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
}

func requestHandler(c *gin.Context) (bool, interface{}) {

	requiredValidation := authFilter(c.Request.URL.Path)

	if requiredValidation {
		return AuthorizationService.ValidateUserAuthorization(AuthorizationService{}, c)
	}
	return true, nil
}
