package login

import (
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	lhelper "opensoach.com/spl/helper"
	lmodels "opensoach.com/spl/models"
	repo "opensoach.com/spl/repository"
	"opensoach.com/spl/webserver/login/dbaccess"
)

var SUB_MODULE_NAME = "SPL.Login"

type Service interface {
}

type AuthService struct {
}

func (AuthService) Auth(username, password, prodcode string) (bool, interface{}) {

	dbErr, dbData := dbaccess.ValidateAuth(repo.Instance().Context.Dynamic.DB, username, password)

	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *dbData

	if len(dbRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = lmodels.MOD_ERR_LOGIN_INVALID_USER
		return false, errModel
	}

	userRecordItem := dbRecord[0]

	if userRecordItem.UserState != lmodels.DB_USER_STATE_ACTIVE {
		errModel := gmodels.APIResponseError{}
		errModel.Code = lmodels.MOD_ERR_LOGIN_INACTIVE_USER_STATE
		return false, errModel
	}

	if userRecordItem.UserCategory == lmodels.DB_USER_CATEGORY_CUSTOMER {
		errModel := gmodels.APIResponseError{}
		errModel.Code = lmodels.MOD_ERR_INVALID_USER_CATEGORY
		return false, errModel
	}

	dbErr, authData := dbaccess.GetUserAuthInfo(repo.Instance().Context.Dynamic.DB, prodcode)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "DB Error occured while login.", dbErr)
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbAuthRecord := *authData
	authRecordItem := dbAuthRecord[0]

	if authRecordItem.CpmId == 0 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = lmodels.MOD_ERR_CUSTOMER_PRODUCT_MAPPING
		return false, errModel
	}

	authResponse := lmodels.AuthResponse{}

	userSessionContext := gmodels.UserSessionInfo{}
	userSessionContext.CpmID = authRecordItem.CpmId
	userSessionContext.CustomerID = authRecordItem.CustomerId
	userSessionContext.UserRoleID = userRecordItem.UserRoleId
	userSessionContext.UserID = userRecordItem.ID
	userSessionContext.ModDB = gmodels.ConfigDB{ConnectionString: authRecordItem.Connectionstring, DBDriver: lmodels.DB_DRIVER_NAME}

	isSuccess, token := lhelper.SessionCreate(repo.Instance().Context, &userSessionContext)
	if isSuccess == false {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_SERVER
		return false, errModel
	}

	authResponse.Token = token
	authResponse.UroleCode = authRecordItem.UserRoleCode

	return true, authResponse
}
