package login

import (
	pcmodels "opensoach.com/prodcore/models"
	"opensoach.com/spl/api/constants"
)

func Init(config *pcmodels.WebServerConfiguration) {

	allGroup := config.AuthorizedRouterHandler[constants.API_AUTHORIZATION_ROUTER_GROUP_KEY]

	registerRouters(allGroup)

	//dbaccess.Init("root:welcome@tcp(localhost:3306)/spl_master?parseTime=true")

}
