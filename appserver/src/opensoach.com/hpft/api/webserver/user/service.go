package user

import (
	"opensoach.com/core/logger"
	repo "opensoach.com/hpft/api/repository"
	"opensoach.com/hpft/api/webserver/user/dbaccess"
	gmodels "opensoach.com/models"
)

var SUB_MODULE_NAME = "HPFT.API.User"

type UserService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service UserService) SelectDoctorUsers() (bool, interface{}) {

	cmpID := service.ExeCtx.SessionInfo.Product.CustProdID

	dbErr, data := dbaccess.GetDoctorUsers(repo.Instance().Context.Master.DBConn, cmpID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting doctor users.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecords := *data

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched doctor users")
	return true, dbRecords
}
