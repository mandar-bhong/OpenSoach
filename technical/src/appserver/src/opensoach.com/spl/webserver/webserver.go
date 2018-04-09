package webserver

import (
	"fmt"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	gmodels "opensoach.com/models"
	lmodels "opensoach.com/spl/models"
	custmodule "opensoach.com/spl/webserver/customer"
	loginModule "opensoach.com/spl/webserver/login"
	productmodule "opensoach.com/spl/webserver/product"
)

func Init(configSetting *gmodels.ConfigSettings) error {

	ginEngine := gin.Default()
	pprof.Register(ginEngine)

	webConfig := &lmodels.WebServerConfiguration{}
	webConfig.WebHandlerEngine = ginEngine
	webConfig.DBConfig = configSetting.DBConfig
	webConfig.WebConf = configSetting.WebConfig

	//webcontent.Init(webConfig)
	//webauth.Init(webConfig)
	loginModule.Init(webConfig)
	productmodule.Init(webConfig)
	custmodule.Init(webConfig)

	err := ginEngine.Run(fmt.Sprintf("%s", configSetting.WebConfig.ServiceAddress))

	return err
}
