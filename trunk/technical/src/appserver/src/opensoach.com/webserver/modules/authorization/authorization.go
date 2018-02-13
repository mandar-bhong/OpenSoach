package authorization

import (
	"github.com/gin-gonic/gin"
	wmodels "opensoach.com/webserver/webmodels"
)

func Init(config *wmodels.WebServerConfiguration) bool {

	config.AuthorizedRouterHandler = make(map[string]*gin.RouterGroup)
	registerRouters(config)

	return true
}
