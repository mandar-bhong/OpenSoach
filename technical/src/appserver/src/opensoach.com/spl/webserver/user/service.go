package user

import (
	"time"

	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	"opensoach.com/spl/constants"
	lmodels "opensoach.com/spl/models"
	repo "opensoach.com/spl/repository"
	"opensoach.com/spl/webserver/user/dbaccess"
)

var SUB_MODULE_NAME = "SPL.User"

type UserService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service UserService) AddUser(userData lmodels.DBSplMasterUserRowModel) (isSuccess bool, successErrorData interface{}) {

	userData.UsrStateSince = time.Now()

	dbErr, userInsertedId := dbaccess.SplMasterUserTableInsert(repo.Instance().Context.Master.DBConn, userData)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	response := lmodels.RecordIdResponse{}
	response.RecId = userInsertedId

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "User added successfully.")

	return true, response
}

func (service UserService) UpdateUserDetails(userData lmodels.DBSplMasterUsrDetailsRowModel) (isSuccess bool, successErrorData interface{}) {

	dbErr, userDetailsData := dbaccess.GetSplMasterUserDetailsTableById(repo.Instance().Context.Master.DBConn, userData.UsrId)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "DB Error occured while login.", dbErr)
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbUserDetailsRecord := *userDetailsData

	if len(dbUserDetailsRecord) < 1 {
		dbErr, userInsertedId := dbaccess.SplMasterUserDetailsTableInsert(repo.Instance().Context.Master.DBConn, userData)
		if dbErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

			errModel := gmodels.APIResponseError{}
			errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
			return false, errModel
		}

		response := lmodels.RecordIdResponse{}
		response.RecId = userInsertedId

		logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "User details inserted successfully.")

		return true, response

	} else {
		dbErr, userAffectedRow := dbaccess.SplMasterUserDetailsTableUpdate(repo.Instance().Context.Master.DBConn, userData)
		if dbErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

			errModel := gmodels.APIResponseError{}
			errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
			return false, errModel
		}

		if userAffectedRow == 0 {
			logger.Context().WithField("InputRequest", userData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

			errModel := gmodels.APIResponseError{}
			errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
			return false, errModel
		}

		logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "User details updated Successfully.")

		return true, nil
	}

}

func (service UserService) UpdateUserState(userData lmodels.DBSplMasterUserRowModel) (isSuccess bool, successErrorData interface{}) {

	userData.UsrStateSince = time.Now()

	dbErr, _ := dbaccess.UpdateUsrState(repo.Instance().Context.Master.DBConn, userData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", userData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "User state updated successfully.")

	return true, nil
}

func (service UserService) ChangeUserPassword(passData lmodels.UpdatePasswordRequest, userid int64) (isSuccess bool, successErrorData interface{}) {

	dbErr, userData := dbaccess.CheckOldPasswordExists(repo.Instance().Context.Master.DBConn, userid, passData.OldPassword)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "DB Error occured while login.", dbErr)
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbUserRecord := *userData

	if len(dbUserRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = constants.MOD_ERR_USER_PASSWORD_NOT_MATCH
		return false, errModel
	} else {

		updateUserData := lmodels.DBSplMasterUserRowModel{}
		updateUserData.UsrId = userid
		updateUserData.UsrPassword = passData.NewPassword

		dbErr, _ := dbaccess.UpdateUsrPassword(repo.Instance().Context.Master.DBConn, updateUserData)
		if dbErr != nil {
			logger.Context().WithField("InputRequest", passData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

			errModel := gmodels.APIResponseError{}
			errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
			return false, errModel
		}

		logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "User password changed successfully.")

		return true, nil
	}
}

func (UserService) GetUserDataList(usrListReqData lmodels.DataListRequest) (bool, interface{}) {

	dataListResponse := lmodels.DataListResponse{}

	filterModel := usrListReqData.Filter.(*lmodels.DBSearchUserRequestFilterDataModel)

	dbErr, userFilteredRecords := dbaccess.GetUsrFilterRecordsCount(repo.Instance().Context.Master.DBConn, filterModel)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}
	dbUserFilteredRecords := *userFilteredRecords
	dataListResponse.FilteredRecords = dbUserFilteredRecords.TotalRecords

	CurrentPage := usrListReqData.CurrentPage
	startingRecord := ((CurrentPage - 1) * usrListReqData.Limit)

	dbErr, usrFilterData := dbaccess.GetUserList(repo.Instance().Context.Master.DBConn, usrListReqData, filterModel, startingRecord)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbUserFilterRecord := *usrFilterData
	dataListResponse.Records = dbUserFilterRecord

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched user list data.")

	return true, dataListResponse

}

func (service UserService) AssociateUserWithCust(reqData *lmodels.CustomerAssociateUserRequest) (isSuccess bool, successErrorData interface{}) {

	dbErr, rsltData := dbaccess.GetUserIdByUserName(repo.Instance().Context.Master.DBConn, reqData.UserName)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *rsltData

	if len(dbRecord) < 1 {

		errModel := gmodels.APIResponseError{}
		errModel.Code = constants.MOD_ERR_USER_NAME_NOT_FOUND
		return false, errModel

	} else {

		reqData.UserId = dbRecord[0].UserId

		usrcpm := lmodels.DBUsrCpmRowModel{}
		usrcpm.CpmId = reqData.CpmId
		usrcpm.UroleId = reqData.UroleId
		usrcpm.UserId = reqData.UserId

		dbErr, insertedId := dbaccess.SplMasterUserCpmTableInsert(repo.Instance().Context.Master.DBConn, usrcpm)
		if dbErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

			errModel := gmodels.APIResponseError{}
			errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
			return false, errModel
		}

		response := lmodels.RecordIdResponse{}
		response.RecId = insertedId

		logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "User associated with customer successfully.")

		return true, response
	}
}
