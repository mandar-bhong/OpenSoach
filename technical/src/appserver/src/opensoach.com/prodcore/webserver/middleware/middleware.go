package middleware

import (
	"github.com/gin-gonic/gin"
	"opensoach.com/core"
	pcmodels "opensoach.com/prodcore/models"
)

var oscontext *core.Context

func Init(osctx *core.Context, config *pcmodels.WebServerConfiguration) {
	config.AuthorizedRouterHandler = make(map[string]*gin.RouterGroup)
	oscontext = osctx
	registerRouters(config)
}
