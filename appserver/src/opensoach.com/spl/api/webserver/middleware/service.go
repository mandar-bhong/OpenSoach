package middleware

import (
	"github.com/gin-gonic/gin"
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	lhelper "opensoach.com/spl/api/helper"
	repo "opensoach.com/spl/api/repository"
)

var SUB_COMPONENT = "SPL.Authorization"

type Service interface {
}

type AuthorizationService struct {
}

func (AuthorizationService) ValidateUserAuthorization(pContext *gin.Context) (bool, interface{}) {
	logger.Context().Log(SUB_COMPONENT, logger.Normal, logger.Debug, "ValidateUserAuthorization")

	oscontext := repo.Instance().Context

	isGetSuccess, _ := lhelper.SessionGet(oscontext, pContext)

	if isGetSuccess == false {
		errorData := gmodels.APIResponseError{}
		errorData.Code = gmodels.MOD_OPER_UNAUTHORIZED
		return false, errorData
	}

	isUpdateSessionSuccess := lhelper.SessionUpdate(oscontext, pContext)
	if !isUpdateSessionSuccess {
		logger.Context().Log(SUB_COMPONENT, logger.Normal, logger.Error, "Failed to save session.")
		errorData := gmodels.APIResponseError{}
		errorData.Code = gmodels.MOD_OPER_ERR_SERVER
		return false, errorData
	}

	logger.Context().Log(SUB_COMPONENT, logger.Normal, logger.Debug, "Session saved successfully.")
	return isUpdateSessionSuccess, "Invalid user"
}
