package endpoint

import (
	"opensoach.com/spl/constants"
	lmodels "opensoach.com/spl/models"
)

func Init(config *lmodels.WebServerConfiguration) {

	allGroup := config.AuthorizedRouterHandler[constants.API_AUTHORIZATION_ROUTER_GROUP_KEY]

	registerRouters(allGroup)
}
