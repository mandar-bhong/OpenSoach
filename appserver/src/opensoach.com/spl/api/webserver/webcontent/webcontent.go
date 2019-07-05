package webcontent

import (
	pcmodels "opensoach.com/prodcore/models"
)

func Init(config *pcmodels.WebServerConfiguration) {

	allGroup := config.WebHandlerEngine.Group("/")

	registerRouters(allGroup)
}
