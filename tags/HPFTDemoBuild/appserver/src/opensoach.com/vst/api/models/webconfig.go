package models

import (
	"github.com/gin-gonic/gin"
	gmodels "opensoach.com/models"
)

type WebServerConfiguration struct {
	WebHandlerEngine        *gin.Engine
	AuthorizedRouterHandler map[string]*gin.RouterGroup
	WebConf                 *gmodels.ConfigWebSettings `json:"webconfig"`
	DBConfig                *gmodels.ConfigDB
}
