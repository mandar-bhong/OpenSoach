package vehicle

import (
	pcconst "opensoach.com/prodcore/constants"
	pcmodels "opensoach.com/prodcore/models"
)

func Init(config *pcmodels.WebServerConfiguration) {

	authorizedGroup := config.AuthorizedRouterHandler[pcconst.API_AUTHORIZATION_ROUTER_GROUP_KEY]
	registerRouters(authorizedGroup)
}
