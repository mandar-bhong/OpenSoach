package webserver

import (
	"fmt"

	"github.com/gin-gonic/gin"
	gmodels "opensoach.com/models"
	webauth "opensoach.com/webserver/modules/authorization"
	loginModule "opensoach.com/webserver/modules/login"
	webcontent "opensoach.com/webserver/modules/webcontent"
	"opensoach.com/webserver/webmodels"
)

func Init(config *gmodels.ConfigSettings) bool {

	ginEngine := gin.Default()

	webConfig := &webmodels.WebServerConfiguration{}
	webConfig.WebHandlerEngine = ginEngine
	webConfig.DBConfig = config.DatabaseConfig

	webcontent.Init(webConfig)
	webauth.Init(webConfig)
	loginModule.Init(webConfig)

	ginEngine.Run(fmt.Sprintf(":%d", config.WebConfig.ServicePort))

	return true
}
