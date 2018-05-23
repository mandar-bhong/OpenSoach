package webserver

import (
	"fmt"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	repo "opensoach.com/hkt/api/repository"
	complaint "opensoach.com/hkt/api/webserver/complaint"
	device "opensoach.com/hkt/api/webserver/device"
	fieldoperator "opensoach.com/hkt/api/webserver/fieldoperator"
	master "opensoach.com/hkt/api/webserver/master"
	service "opensoach.com/hkt/api/webserver/service"
	servicepoint "opensoach.com/hkt/api/webserver/servicepoint"
	task "opensoach.com/hkt/api/webserver/task"
	"opensoach.com/hkt/api/webserver/webcontent"
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

	webcontent.Init(webConfig)

	pcwebsermid.Init(repo.Instance().Context, webConfig)

	task.Init(webConfig)
	fieldoperator.Init(webConfig)
	complaint.Init(webConfig)
	service.Init(webConfig)
	master.Init(webConfig)
	servicepoint.Init(webConfig)
	device.Init(webConfig)

	var webServerStartErr error

	go func() {
		webServerStartErr = ginEngine.Run(fmt.Sprintf("%s", configSetting.WebConfig.ServiceAddress))
	}()

	time.Sleep(time.Second * 2)

	return webServerStartErr
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
