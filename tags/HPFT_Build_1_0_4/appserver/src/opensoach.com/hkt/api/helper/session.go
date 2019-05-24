package helper

import (
	"github.com/gin-gonic/gin"
	gcore "opensoach.com/core"
	gmodels "opensoach.com/models"
	pchelper "opensoach.com/prodcore/helper"
)

func SessionCreate(osContext *gcore.Context, pSessionData *gmodels.UserSessionInfo) (bool, string) {
	return pchelper.SessionCreate(osContext, pSessionData)
}

func SessionGet(osContext *gcore.Context, ginContext *gin.Context) (bool, *gmodels.UserSessionInfo) {
	return pchelper.SessionGet(osContext, ginContext)
}

func SessionUpdate(osContext *gcore.Context, ginContext *gin.Context) bool {
	return pchelper.SessionUpdate(osContext, ginContext)
}

func SessionDelete(osContext *gcore.Context, ginContext *gin.Context) bool {
	return pchelper.SessionDelete(osContext, ginContext)
}
