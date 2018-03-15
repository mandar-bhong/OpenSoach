package product

import (
	lmodels "opensoach.com/spl/models"
)

func Init(config *lmodels.WebServerConfiguration) {

	allGroup := config.WebHandlerEngine.Group("/")

	registerRouters(allGroup)

	//	dbaccess.Init("root:welcome@tcp(localhost:3306)/spl_master?parseTime=true")

}
