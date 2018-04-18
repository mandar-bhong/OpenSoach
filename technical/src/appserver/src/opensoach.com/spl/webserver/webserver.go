package webserver

import (
	"fmt"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	gmodels "opensoach.com/models"
	lmodels "opensoach.com/spl/models"
	corporatemodule "opensoach.com/spl/webserver/corporate"
	custmodule "opensoach.com/spl/webserver/customer"
	loginModule "opensoach.com/spl/webserver/login"
	productmodule "opensoach.com/spl/webserver/product"
	usermodule "opensoach.com/spl/webserver/user"
	"opensoach.com/spl/webserver/webcontent"
)

func Init(configSetting *gmodels.ConfigSettings) error {

	ginEngine := gin.Default()
	pprof.Register(ginEngine)

	enableCrossDomain(ginEngine)

	webConfig := &lmodels.WebServerConfiguration{}
	webConfig.WebHandlerEngine = ginEngine
	webConfig.DBConfig = configSetting.DBConfig
	webConfig.WebConf = configSetting.WebConfig

	webcontent.Init(webConfig)
	//webauth.Init(webConfig)
	loginModule.Init(webConfig)
	productmodule.Init(webConfig)
	custmodule.Init(webConfig)
	usermodule.Init(webConfig)
	corporatemodule.Init(webConfig)

	err := ginEngine.Run(fmt.Sprintf("%s", configSetting.WebConfig.ServiceAddress))

	return err
}

func enableCrossDomain(c *gin.Engine) {
	c.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type, Cookies",
		MaxAge:          5000 * time.Second, // original value was 50
		Credentials:     true,
		ValidateHeaders: true,
		ExposedHeaders:  "Cache-Control, Content-Language, Content-Type, Expires, Last-Modified, Pragma",
	}))
}
