package login

import (
	"opensoach.com/webserver/modules/login/dbaccess"
	wmodels "opensoach.com/webserver/webmodels"
)

func Init(config *wmodels.WebServerConfiguration) bool {

	isSuccess := dbaccess.Init(config.DBConfig)

	return isSuccess
}
