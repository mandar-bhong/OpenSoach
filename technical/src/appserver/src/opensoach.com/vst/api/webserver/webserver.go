package webserver

import (
	"fmt"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	gmodels "opensoach.com/models"
	pcmodels "opensoach.com/prodcore/models"
	pcwebsermid "opensoach.com/prodcore/webserver/middleware"
	"opensoach.com/vst/api/constants"
	repo "opensoach.com/vst/api/repository"
	complaint "opensoach.com/vst/api/webserver/complaint"
	"opensoach.com/vst/api/webserver/dashboard"
	device "opensoach.com/vst/api/webserver/device"
	feedback "opensoach.com/vst/api/webserver/feedback"
	fieldoperator "opensoach.com/vst/api/webserver/fieldoperator"
	job "opensoach.com/vst/api/webserver/job"
	master "opensoach.com/vst/api/webserver/master"
	report "opensoach.com/vst/api/webserver/report"
	service "opensoach.com/vst/api/webserver/service"
	servicepoint "opensoach.com/vst/api/webserver/servicepoint"
	splprod "opensoach.com/vst/api/webserver/splprod"
	task "opensoach.com/vst/api/webserver/task"
	vehicle "opensoach.com/vst/api/webserver/vehicle"
	"opensoach.com/vst/api/webserver/webcontent"
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

	pcwebsermid.Init(repo.Instance().Context, webConfig, AuthorizationFilter) // all api need to validated

	task.Init(webConfig)
	fieldoperator.Init(webConfig)
	complaint.Init(webConfig)
	service.Init(webConfig)
	master.Init(webConfig)
	servicepoint.Init(webConfig)
	device.Init(webConfig)
	splprod.Init(webConfig)
	report.Init(webConfig)
	dashboard.Init(webConfig)
	feedback.Init(webConfig)
	vehicle.Init(webConfig)
	job.Init(webConfig)

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
		RequestHeaders:  "Origin, Authorization, Content-Type, Cookies,responseType",
		MaxAge:          5000 * time.Second, // original value was 50
		Credentials:     true,
		ValidateHeaders: false,
		ExposedHeaders:  "Cache-Control, Content-Language, Content-Type, Expires, Last-Modified, Pragma",
	}))
}

func AuthorizationFilter(reqURL string) (isAuthorizationRequred bool) {
	switch reqURL {
	case constants.API_SPL_PROD_BASE_URL:
		return false
	}
	return true
}
