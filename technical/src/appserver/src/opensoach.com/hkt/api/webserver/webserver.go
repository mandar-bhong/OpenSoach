package webserver

import (
	"fmt"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	repo "opensoach.com/hkt/api/repository"
	"opensoach.com/hkt/api/webserver/task"
	gmodels "opensoach.com/models"
	pcmodels "opensoach.com/prodcore/models"
	pcwebsermid "opensoach.com/prodcore/webserver/middleware"
)

func Init(configSetting *gmodels.ConfigSettings) error {

	ginEngine := gin.Default()
	pprof.Register(ginEngine)

	enableCrossDomain(ginEngine)

	webConfig := &pcmodels.WebServerConfiguration{}
	webConfig.WebHandlerEngine = ginEngine
	webConfig.DBConfig = configSetting.DBConfig
	webConfig.WebConf = configSetting.WebConfig

	pcwebsermid.Init(repo.Instance().Context, webConfig)

	task.Init(webConfig)

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
