package middleware

import (
	//pcmiddleware "opensoach.com/prodcore/webserver/middleware"
	"opensoach.com/spl/constants"
	lmodels "opensoach.com/spl/models"
)

func registerRouters(config *lmodels.WebServerConfiguration) {
	allAuthorizedRouter := config.WebHandlerEngine.Group("/")

	config.AuthorizedRouterHandler[constants.API_AUTHORIZATION_ROUTER_GROUP_KEY] = allAuthorizedRouter
}
