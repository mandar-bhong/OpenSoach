package webmodels

import (
	"github.com/gin-gonic/gin"
	gmodels "opensoach.com/models"
)

type WebServerConfiguration struct {
	WebHandlerEngine        *gin.Engine
	AuthorizedRouterHandler *gin.RouterGroup
	DBConfig                gmodels.DatabaseSettings
}
