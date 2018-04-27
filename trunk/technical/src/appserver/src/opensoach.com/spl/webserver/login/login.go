package login

import (
	"opensoach.com/spl/constants"
	lmodels "opensoach.com/spl/models"
)

func Init(config *lmodels.WebServerConfiguration) {

	allGroup := config.AuthorizedRouterHandler[constants.API_AUTHORIZATION_ROUTER_GROUP_KEY]

	registerRouters(allGroup)

	//dbaccess.Init("root:welcome@tcp(localhost:3306)/spl_master?parseTime=true")

}
