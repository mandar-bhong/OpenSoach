package authorization

import (
	"net/http"

	"github.com/gin-gonic/gin"
	wmodels "opensoach.com/webserver/webmodels"
)

func registerRouters(config *wmodels.WebServerConfiguration) {
	allAuthorizedRouter := config.WebHandlerEngine.Group("/")
	allAuthorizedRouter.Use(commonHandler)
	config.AuthorizedRouterHandler["ALL"] = allAuthorizedRouter
}

func commonHandler(c *gin.Context) {
	var isSuccess bool
	var successErrorData interface{}

	responsePayload := wmodels.PayloadResponse{}

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

	case "/login":
		return true, nil
		break

	default:
		isSuccess, successErrorData = AuthorizationService.ValidateUserAuthorization(AuthorizationService{}, c)
		break
	}

	return isSuccess, successErrorData
}
