package middleware

import (
	"opensoach.com/prodcore/constants"
	lmodels "opensoach.com/prodcore/models"
)

func registerRouters(config *lmodels.WebServerConfiguration) {
	allAuthorizedRouter := config.WebHandlerEngine.Group("/")

	allAuthorizedRouter.Use(RequestLogger())
	allAuthorizedRouter.Use(AuthorizationHandler())

	config.AuthorizedRouterHandler[constants.API_AUTHORIZATION_ROUTER_GROUP_KEY] = allAuthorizedRouter
}
