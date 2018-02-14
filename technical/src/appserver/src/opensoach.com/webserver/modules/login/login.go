package login

import (
	"opensoach.com/webserver/modules/login/dbaccess"
	wmodels "opensoach.com/webserver/webmodels"
)

func Init(config *wmodels.WebServerConfiguration) bool {
	routerGroup := config.AuthorizedRouterHandler["ALL"]
	isSuccess := dbaccess.Init(config.DBConfig)

	registerRouters(routerGroup)

	return isSuccess
}
