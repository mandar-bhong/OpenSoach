package task

import (
	lmodels "opensoach.com/spl/models"
)

func Init(config *lmodels.WebServerConfiguration) {

	allGroup := config.WebHandlerEngine.Group("/")

	registerRouters(allGroup)
}
