package webserver

import (
	"fmt"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	gmodels "opensoach.com/models"
	pcmodels "opensoach.com/prodcore/models"
	pcmiddleware "opensoach.com/prodcore/webserver/middleware"
	repo "opensoach.com/spl/api/repository"
	corporatemodule "opensoach.com/spl/api/webserver/corporate"
	custmodule "opensoach.com/spl/api/webserver/customer"
	devicemodule "opensoach.com/spl/api/webserver/device"
	endpointmodule "opensoach.com/spl/api/webserver/endpoint"
	loginModule "opensoach.com/spl/api/webserver/login"
	"opensoach.com/spl/api/webserver/middleware"
	productmodule "opensoach.com/spl/api/webserver/product"
	usermodule "opensoach.com/spl/api/webserver/user"
	"opensoach.com/spl/api/webserver/webcontent"
)

func Init(configSetting *gmodels.ConfigSettings) error {

	ginEngine := gin.Default()
	pprof.Register(ginEngine)

	enableCrossDomain(ginEngine)

	webConfig := &pcmodels.WebServerConfiguration{}
	webConfig.WebHandlerEngine = ginEngine
	webConfig.DBConfig = configSetting.DBConfig
	webConfig.WebConf = configSetting.WebConfig
	webConfig.AuthorizedRouterHandler = make(map[string]*gin.RouterGroup)

	webcontent.Init(webConfig)
	pcmiddleware.Init(repo.Instance().Context, webConfig, middleware.AuthorizationFilter)
	//webauth.Init(webConfig)
	loginModule.Init(webConfig)
	productmodule.Init(webConfig)
	custmodule.Init(webConfig)
	usermodule.Init(webConfig)
	corporatemodule.Init(webConfig)
	devicemodule.Init(webConfig)
	endpointmodule.Init(webConfig)

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
