package customer

import (
	pcmodels "opensoach.com/prodcore/models"
	"opensoach.com/spl/constants"
)

func Init(config *pcmodels.WebServerConfiguration) {

	allGroup := config.AuthorizedRouterHandler[constants.API_AUTHORIZATION_ROUTER_GROUP_KEY]

	registerRouters(allGroup)
}
