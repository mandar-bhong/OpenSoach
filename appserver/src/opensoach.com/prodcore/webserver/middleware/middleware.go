package middleware

import (
	"github.com/gin-gonic/gin"
	"opensoach.com/core"
	pcmodels "opensoach.com/prodcore/models"
)

type AuthFilterHandler func(url string) bool

var oscontext *core.Context
var authFilter AuthFilterHandler

func Init(osctx *core.Context, config *pcmodels.WebServerConfiguration, authorizationFilter AuthFilterHandler) {
	config.AuthorizedRouterHandler = make(map[string]*gin.RouterGroup)
	oscontext = osctx
	authFilter = authorizationFilter
	registerRouters(config)
}
