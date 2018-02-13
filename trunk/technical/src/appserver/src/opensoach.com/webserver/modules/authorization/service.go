package authorization

import (
	"github.com/gin-gonic/gin"
)

type Service interface {
}

type AuthorizationService struct {
}

func (AuthorizationService) ValidateUserAuthorization(pContext *gin.Context) (bool, interface{}) {
	//logger.Log(helper.MODULENAME, logger.DEBUG, "ValidateUserAuthorization")
	//isUpdateSessionSuccess := ghelper.SessionUpdateExpiration(pContext)
	//	if !isUpdateSessionSuccess {
	//		//logger.Log(helper.MODULENAME, logger.ERROR, "Failed to save session.")
	//	}

	//logger.Log(helper.MODULENAME, logger.DEBUG, "Session saved successfully.")
	//return isUpdateSessionSuccess, "Invalid user"

	return true, "Invalid user"
}
