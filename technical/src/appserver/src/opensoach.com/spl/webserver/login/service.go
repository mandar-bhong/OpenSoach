package login

import (
	"github.com/gin-gonic/gin"
	gcore "opensoach.com/core"
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	"opensoach.com/spl/constants"
	lhelper "opensoach.com/spl/helper"
	lmodels "opensoach.com/spl/models"
	repo "opensoach.com/spl/repository"
	"opensoach.com/spl/webserver/login/dbaccess"
)

var SUB_MODULE_NAME = "SPL.Login"

type Service interface {
}

type AuthService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (AuthService) Auth(username, password, prodcode string) (bool, interface{}) {

	dbErr, dbData := dbaccess.ValidateAuth(repo.Instance().Context.Master.DBConn, username, password)

	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *dbData

	if len(dbRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = constants.MOD_ERR_LOGIN_INVALID_USER
		return false, errModel
	}

	userRecordItem := dbRecord[0]

	if userRecordItem.UsrState != constants.DB_USER_STATE_ACTIVE {
		errModel := gmodels.APIResponseError{}
		errModel.Code = constants.MOD_ERR_LOGIN_INACTIVE_USER_STATE
		return false, errModel
	}

	authResponse := lmodels.AuthResponse{}

	if userRecordItem.UsrCategory == constants.DB_USER_CATEGORY_OS {

		dbErr, authData := dbaccess.GetUserAuthInfo(repo.Instance().Context.Master.DBConn, userRecordItem.UroleId)

		if dbErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "DB Error occured while login.", dbErr)
			errModel := gmodels.APIResponseError{}
			errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
			return false, errModel
		}

		dbAuthData := *authData

		userSessionContext := gmodels.UserSessionInfo{}
		userSessionContext.UserRoleID = *userRecordItem.UroleId
		userSessionContext.UserID = userRecordItem.UserId

		isSuccess, token := lhelper.SessionCreate(repo.Instance().Context, &userSessionContext)
		if isSuccess == false {
			errModel := gmodels.APIResponseError{}
			errModel.Code = gmodels.MOD_OPER_ERR_SERVER
			return false, errModel
		}

		authResponse.Token = token
		authResponse.UroleCode = dbAuthData.UserRoleCode

		logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "User login successfull")

	} else if userRecordItem.UsrCategory == constants.DB_USER_CATEGORY_CUSTOMER {
		dbErr, authData := dbaccess.GetUserAuthInfoCategoryCustomer(repo.Instance().Context.Master.DBConn, prodcode, userRecordItem.UserId)

		if dbErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "DB Error occured while login.", dbErr)
			errModel := gmodels.APIResponseError{}
			errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
			return false, errModel
		}

		dbAuthRecord := *authData

		if len(dbAuthRecord) < 1 {
			errModel := gmodels.APIResponseError{}
			errModel.Code = constants.MOD_ERR_CUSTOMER_PRODUCT_MAPPING
			return false, errModel

		}

		authRecordItem := dbAuthRecord[0]

		userSessionContext := gmodels.UserSessionInfo{}
		userSessionContext.CpmID = authRecordItem.CpmId
		userSessionContext.CustomerID = authRecordItem.CustomerId
		userSessionContext.UserRoleID = authRecordItem.UserRoleId
		userSessionContext.UserID = userRecordItem.UserId
		userSessionContext.ModDB = gmodels.ConfigDB{ConnectionString: authRecordItem.Connectionstring, DBDriver: constants.DB_DRIVER_NAME}

		isSuccess, token := lhelper.SessionCreate(repo.Instance().Context, &userSessionContext)
		if isSuccess == false {
			errModel := gmodels.APIResponseError{}
			errModel.Code = gmodels.MOD_OPER_ERR_SERVER
			return false, errModel
		}

		authResponse.Token = token
		authResponse.UroleCode = authRecordItem.UserRoleCode

		logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "User login successfull : User Category Customer")
	}
	return true, authResponse
}

func (service AuthService) GetUserLoginDetails() (bool, interface{}) {

	userId := service.ExeCtx.SessionInfo.UserID

	dbErr, userLoginInfo := dbaccess.GetUserLoginInfo(repo.Instance().Context.Master.DBConn, userId)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}
	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched minimum user login details")
	return true, userLoginInfo
}

func (service AuthService) UserLogout(pContext *gin.Context) bool {
	isSuccess := lhelper.SessionDelete(repo.Instance().Context, pContext)
	if isSuccess != false {
		logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "User Logged out succesfully")
	}
	return isSuccess
}

func (service AuthService) GetCustomerLoginDetails() (bool, interface{}) {
	custId := service.ExeCtx.SessionInfo.CustomerID
	dbErr, customerLoginInfo := dbaccess.GetCustomerLoginInfo(repo.Instance().Context.Master.DBConn, custId)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched minimum customer details.")

	return true, customerLoginInfo
}

func (service AuthService) ValidateAuthToken(token string, osContext *gcore.Context) (bool, interface{}) {

	isSuccess, _ := osContext.Master.Cache.Get(token)
	if isSuccess == false {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_USER_TOKEN_NOT_AVAILABLE
		return isSuccess, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully Validated Auth Token.")

	return isSuccess, nil
}
