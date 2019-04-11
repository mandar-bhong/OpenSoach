package webserver

import (
	"fmt"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"opensoach.com/hpft/api/constants"
	repo "opensoach.com/hpft/api/repository"
	complaint "opensoach.com/hpft/api/webserver/complaint"
	"opensoach.com/hpft/api/webserver/dashboard"
	device "opensoach.com/hpft/api/webserver/device"
	document "opensoach.com/hpft/api/webserver/document"
	feedback "opensoach.com/hpft/api/webserver/feedback"
	fieldoperator "opensoach.com/hpft/api/webserver/fieldoperator"
	master "opensoach.com/hpft/api/webserver/master"
	patient "opensoach.com/hpft/api/webserver/patient"
	report "opensoach.com/hpft/api/webserver/report"
	service "opensoach.com/hpft/api/webserver/service"
	servicepoint "opensoach.com/hpft/api/webserver/servicepoint"
	splprod "opensoach.com/hpft/api/webserver/splprod"
	task "opensoach.com/hpft/api/webserver/task"
	user "opensoach.com/hpft/api/webserver/user"
	"opensoach.com/hpft/api/webserver/webcontent"
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
	patient.Init(webConfig)
	document.Init(webConfig)
	user.Init(webConfig)

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
	case constants.API_SPL_PROD_BASE_URL,
		constants.API_DEVICE_DOCUMENT_DOWNLOAD:
		return false
	}
	return true
}
