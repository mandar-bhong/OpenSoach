package webserver

import (
	"github.com/gin-gonic/gin"
	loginModule "opensoach.com/webserver/modules/login"
	webcontent "opensoach.com/webserver/modules/webcontent"
	"opensoach.com/webserver/webmodels"
)

func Init() bool {

	ginEngine := gin.Default()

	webConfig := &webmodels.WebServerConfiguration{}
	webConfig.WebHandlerEngine = ginEngine

	webcontent.Init(webConfig)

	loginModule.Init()

	return true
}
