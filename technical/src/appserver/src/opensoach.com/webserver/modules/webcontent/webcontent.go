package webcontent

import (
	"opensoach.com/webserver/webmodels"
)

func Init(conf *webmodels.WebServerConfiguration) bool {

	registerRouters(conf.WebHandlerEngine)
	return true
}
