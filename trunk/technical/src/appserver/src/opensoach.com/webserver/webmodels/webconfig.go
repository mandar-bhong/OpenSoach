package webmodels

import (
	"github.com/gin-gonic/gin"
)

type WebServerConfiguration struct {
	WebHandlerEngine        *gin.Engine
	AuthorizedRouterHandler *gin.RouterGroup
}
