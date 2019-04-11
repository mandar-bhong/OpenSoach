package user

import (
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	"opensoach.com/spl/api/constants"
	lmodels "opensoach.com/spl/api/models"
	repo "opensoach.com/spl/api/repository"
	"opensoach.com/spl/api/webserver/user/dbaccess"
)

var SUB_MODULE_NAME = "SPL.User"

type UserService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service UserService) AddUser(userData lmodels.DBSplMasterUserRowModel) (isSuccess bool, successErrorData interface{}) {

	userData.UsrPassword = ghelper.GetUserPassword()
	userData.UsrStateSince = ghelper.GetCurrentTime()

	dbTxErr, tx := dbaccess.GetDBTransaction(repo.Instance().Context.Master.DBConn)

	if dbTxErr != nil {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbErr, userInsertedId := dbaccess.SplMasterUserTableInsert(tx, userData)
	if dbErr != nil {

		txErr := tx.Rollback()

		if txErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to rollback transaction", txErr)
		}

		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while adding new user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	txErr := tx.Commit()
	if txErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to commit transaction", txErr)
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	response := gmodels.APIRecordIdResponse{}
	response.RecId = userInsertedId

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "User added successfully.")

	return true, response
}

func (service UserService) AddCUUser(userData lmodels.DBSplMasterUserRowModel) (isSuccess bool, successErrorData interface{}) {

	usrcpm := lmodels.DBUsrCpmRowModel{}
	usrcpm.UroleId = *userData.UroleId
	userData.UroleId = nil

	userData.UsrPassword = ghelper.GetUserPassword()
	userData.UsrCategory = constants.DB_USER_CATEGORY_CUSTOMER
	userData.UsrState = constants.DB_USER_STATE_ACTIVE
	userData.UsrStateSince = ghelper.GetCurrentTime()

	dbTxErr, tx := dbaccess.GetDBTransaction(repo.Instance().Context.Master.DBConn)

	if dbTxErr != nil {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbErr, userInsertedId := dbaccess.SplMasterUserTableInsert(tx, userData)
	if dbErr != nil {

		txErr := tx.Rollback()

		if txErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to rollback transaction", txErr)
		}

		errModel := gmodels.APIResponseError{}
		errHandledIsSuccess, errorCode := ghelper.GetApplicationErrorCodeFromDBError(dbErr)

		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while adding user.", dbErr)

		if errHandledIsSuccess == true {
			errModel.Code = errorCode
			return false, errModel
		}

		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "User added successfully.")

	usrcpm.CpmId = service.ExeCtx.SessionInfo.Product.CustProdID
	usrcpm.UserId = userInsertedId
	usrcpm.UcpmState = constants.DB_USER_CPM_STATE_ACTIVE
	usrcpm.UcpmStateSince = ghelper.GetCurrentTime()

	dberr, _ := dbaccess.SplMasterUserCpmTableInsert(tx, usrcpm)
	if dberr != nil {

		txErr := tx.Rollback()

		if txErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to rollback transaction", txErr)
		}

		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while associating cpm with user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	response := gmodels.APIRecordIdResponse{}
	response.RecId = userInsertedId
	txErr := tx.Commit()

	if txErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to commit transaction", txErr)
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully Add user and associated with customer product")

	return true, response
}

func (service UserService) UpdateUserDetails(userData lmodels.DBSplMasterUsrDetailsRowModel) (isSuccess bool, successErrorData interface{}) {

	dbErr, userDetailsData := dbaccess.GetSplMasterUserDetailsTableById(repo.Instance().Context.Master.DBConn, userData.UsrId)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "DB Error occured while getting user details by id.", dbErr)
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbUserDetailsRecord := *userDetailsData

	if len(dbUserDetailsRecord) < 1 {
		dbErr, userInsertedId := dbaccess.SplMasterUserDetailsTableInsert(repo.Instance().Context.Master.DBConn, userData)
		if dbErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while adding user details info.", dbErr)

			errModel := gmodels.APIResponseError{}
			errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
			return false, errModel
		}

		response := gmodels.APIRecordIdResponse{}
		response.RecId = userInsertedId

		logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "User details inserted successfully.")

		return true, response

	} else {
		dbErr, userAffectedRow := dbaccess.SplMasterUserDetailsTableUpdate(repo.Instance().Context.Master.DBConn, userData)
		if dbErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating user details info.", dbErr)

			errModel := gmodels.APIResponseError{}
			errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
			return false, errModel
		}

		if userAffectedRow == 0 {
			logger.Context().WithField("InputRequest", userData).LogError(SUB_MODULE_NAME, logger.Normal, "Update request has no updated data.", dbErr)

			errModel := gmodels.APIResponseError{}
			errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
			return false, errModel
		}

		logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "User details updated Successfully.")

		return true, nil
	}

}

func (service UserService) UpdateUserState(userData lmodels.DBSplMasterUserRowModel) (isSuccess bool, successErrorData interface{}) {

	userData.UsrStateSince = ghelper.GetCurrentTime()

	dbErr, _ := dbaccess.UpdateUsrState(repo.Instance().Context.Master.DBConn, userData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", userData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating user state.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "User state updated successfully.")

	return true, nil
}

func (service UserService) ChangeUserPassword(passData lmodels.APIUpdatePasswordRequest, userid int64) (isSuccess bool, successErrorData interface{}) {

	dbErr, userData := dbaccess.CheckOldPasswordExists(repo.Instance().Context.Master.DBConn, userid, passData.OldPassword)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "DB Error occured while validating user password.", dbErr)
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
			logger.Context().WithField("InputRequest", passData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating user password.", dbErr)

			errModel := gmodels.APIResponseError{}
			errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
			return false, errModel
		}

		logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "User password changed successfully.")

		return true, nil
	}
}

func (UserService) GetCUDataList(usrListReqData gmodels.APIDataListRequest) (bool, interface{}) {

	dataListResponse := gmodels.APIDataListResponse{}

	filterModel := usrListReqData.Filter.(*lmodels.DBSearchUserRequestFilterDataModel)

	CurrentPage := usrListReqData.CurrentPage
	startingRecord := ((CurrentPage - 1) * usrListReqData.Limit)

	dbErr, listData := dbaccess.GetCustUsrFilterList(repo.Instance().Context.Master.DBConn, filterModel, usrListReqData, startingRecord)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting user data list.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbListDataRecord := *listData

	dataListResponse.FilteredRecords = dbListDataRecord.RecordCount
	dataListResponse.Records = dbListDataRecord.RecordList

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched user list data.")

	return true, dataListResponse

}

func (service UserService) GetUserDetailsInfo(userID int64) (bool, interface{}) {

	dbErr, userData := dbaccess.GetUserById(repo.Instance().Context.Master.DBConn, userID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting user by id.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbUserRecord := *userData

	if len(dbUserRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	dbErr, userDetails := dbaccess.GetUserDetailsById(repo.Instance().Context.Master.DBConn, userID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting user details by id.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *userDetails

	if len(dbRecord) < 1 {
		return true, nil
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched user details")
	return true, dbRecord[0]
}

func (UserService) GetOSUDataList(usrListReqData gmodels.APIDataListRequest) (bool, interface{}) {

	dataListResponse := gmodels.APIDataListResponse{}

	filterModel := usrListReqData.Filter.(*lmodels.DBSearchUserRequestFilterDataModel)

	CurrentPage := usrListReqData.CurrentPage
	startingRecord := ((CurrentPage - 1) * usrListReqData.Limit)

	dbErr, listData := dbaccess.GetOSUsrFilterList(repo.Instance().Context.Master.DBConn, filterModel, usrListReqData, startingRecord)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting user data list.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbListDataRecord := *listData

	dataListResponse.FilteredRecords = dbListDataRecord.RecordCount
	dataListResponse.Records = dbListDataRecord.RecordList

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched user list data.")

	return true, dataListResponse

}

func (service UserService) AssociateUserWithCust(reqData *lmodels.APICustomerAssociateUserRequest) (isSuccess bool, successErrorData interface{}) {

	if reqData.UserId == 0 {

		dbErr, rsltData := dbaccess.GetUserIdByUserName(repo.Instance().Context.Master.DBConn, reqData.UserName)
		if dbErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting usrid by user name.", dbErr)

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
		}

	}

	usrcpm := lmodels.DBUsrCpmRowModel{}
	usrcpm.CpmId = reqData.CpmId
	usrcpm.UroleId = reqData.UroleId
	usrcpm.UserId = reqData.UserId
	usrcpm.UcpmState = reqData.UcpmState
	usrcpm.UcpmStateSince = ghelper.GetCurrentTime()

	dbTxErr, tx := dbaccess.GetDBTransaction(repo.Instance().Context.Master.DBConn)

	if dbTxErr != nil {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbErr, insertedId := dbaccess.SplMasterUserCpmTableInsert(tx, usrcpm)
	if dbErr != nil {
		txErr := tx.Rollback()

		if txErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to rollback transaction", txErr)
		}
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while associating user with cpm.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	txErr := tx.Commit()
	if txErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to commit transaction", txErr)
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	response := gmodels.APIRecordIdResponse{}
	response.RecId = insertedId

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "User associated with customer successfully.")

	taskUserAssociatedModel := gmodels.TaskUserAssociatedModel{}
	taskUserAssociatedModel.UserID = usrcpm.UserId
	isSendSuccess := repo.Instance().SendTaskToServer(gmodels.TASK_API_USER_ASSOCIATED, service.ExeCtx.SessionToken, taskUserAssociatedModel)

	if isSendSuccess == false {
		logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Unable to submit task for user associated")
	}

	return true, response

}

func (UserService) GetUserRoleListOSU() (bool, interface{}) {

	dbErr, listData := dbaccess.GetUroleListOSU(repo.Instance().Context.Master.DBConn)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting user role data list.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched urole data list.")

	return true, listData

}

func (UserService) GetUserRoleList(prodCode string) (bool, interface{}) {

	dbErr, listData := dbaccess.GetUroleList(repo.Instance().Context.Master.DBConn, prodCode)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting user role data list.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched urole data list.")

	return true, listData

}

func (service UserService) GetUserProdAssociation(userID int64) (bool, interface{}) {

	dbErr, data := dbaccess.GetProdAssociationByUsrId(repo.Instance().Context.Master.DBConn, userID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting user product association.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecords := *data

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched User Product association list")
	return true, dbRecords
}

func (service UserService) UpdateUcpmState(reqData *lmodels.DBUsrCpmStateUpdateRowModel) (isSuccess bool, successErrorData interface{}) {

	reqData.UcpmStateSince = ghelper.GetCurrentTime()

	dbErr, _ := dbaccess.UcpmStateUpdate(repo.Instance().Context.Master.DBConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updatin ucpm state.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Ucpm state updated successfully.")

	return true, nil
}

func (service UserService) UpdateUser(reqData *lmodels.DBUserUpdateRowModel) (isSuccess bool, successErrorData interface{}) {

	dbErr, data := dbaccess.GetUserById(repo.Instance().Context.Master.DBConn, reqData.UserId)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting user by id.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbrecord := *data
	if reqData.UsrState == dbrecord[0].UsrState {
		reqData.UsrStateSince = dbrecord[0].UsrStateSince
	} else {
		reqData.UsrStateSince = ghelper.GetCurrentTime()
	}

	dbErr, affectedRow := dbaccess.UserUpdate(repo.Instance().Context.Master.DBConn, reqData)
	if dbErr != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating user info.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	if affectedRow == 0 {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Update request has no updated data.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "User data updated successfully.")

	return true, nil
}

func (service UserService) GetUserInfo(userID int64) (bool, interface{}) {

	dbErr, userData := dbaccess.GetUserById(repo.Instance().Context.Master.DBConn, userID)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting user info.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *userData

	if len(dbRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched user master details")
	return true, dbRecord[0]
}

func (service UserService) GetCUUserInfo(userID int64) (bool, interface{}) {

	cpmId := service.ExeCtx.SessionInfo.Product.CustProdID

	dbErr, userData := dbaccess.GetCUUserById(repo.Instance().Context.Master.DBConn, userID, cpmId)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting user info.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *userData

	if len(dbRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched user master details")
	return true, dbRecord[0]
}

func (service UserService) UpdateCUUser(reqData *lmodels.APICUUserUpdateRequestModel) (isSuccess bool, successErrorData interface{}) {

	dberror, data := dbaccess.GetUserById(repo.Instance().Context.Master.DBConn, reqData.UserId)
	if dberror != nil {
		logger.Context().WithField("InputRequest", reqData).LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting user by id.", dberror)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	userupdatemodel := &lmodels.DBCUUserUpateRowModel{}
	userupdatemodel.UserId = reqData.UserId
	userupdatemodel.UsrState = reqData.UsrState

	dbrecord := *data
	if reqData.UsrState == dbrecord[0].UsrState {
		userupdatemodel.UsrStateSince = dbrecord[0].UsrStateSince
	} else {
		userupdatemodel.UsrStateSince = ghelper.GetCurrentTime()
	}

	dbTxErr, tx := dbaccess.GetDBTransaction(repo.Instance().Context.Master.DBConn)

	if dbTxErr != nil {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbErr, _ := dbaccess.CUUserUpdate(tx, userupdatemodel)

	if dbErr != nil {
		txErr := tx.Rollback()

		if txErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to rollback transaction", txErr)
		}

		errModel := gmodels.APIResponseError{}
		errHandledIsSuccess, errorCode := ghelper.GetApplicationErrorCodeFromDBError(dbErr)

		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating user info.", dbErr)

		if errHandledIsSuccess == true {
			errModel.Code = errorCode
			return false, errModel
		}

		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "User data updated successfully.")

	ucpmupdatemodel := &lmodels.DBCUUcpmUpdateRowModel{}
	ucpmupdatemodel.UserId = reqData.UserId
	ucpmupdatemodel.UroleId = reqData.UroleId

	dberr, _ := dbaccess.CUUcpmUpdate(tx, ucpmupdatemodel)

	if dberr != nil {

		txErr := tx.Rollback()

		if txErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to rollback transaction", txErr)
		}

		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updatin ucpm info.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	txErr := tx.Commit()

	if txErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to commit transaction", txErr)
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully update user and user role id")

	return true, nil
}

func (service UserService) UserActivation(req lmodels.APIUserActivateRequestModel) (bool, interface{}) {

	dbErr, userData := dbaccess.ValidateUsrActivation(repo.Instance().Context.Master.DBConn, req.Code)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user activation.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *userData

	if len(dbRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	dbUserUpdateActivationDataModel := lmodels.DBUserUpdateActivationDataModel{}
	dbUserUpdateActivationDataModel.UserId = dbRecord[0].UsrId
	dbUserUpdateActivationDataModel.UsrState = constants.DB_USER_STATE_ACTIVE
	dbUserUpdateActivationDataModel.UsrPassword = req.UsrPassword

	dbTxErr, tx := dbaccess.GetDBTransaction(repo.Instance().Context.Master.DBConn)

	if dbTxErr != nil {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dberr, _ := dbaccess.UpdateUsrActivationInfo(tx, &dbUserUpdateActivationDataModel)
	if dberr != nil {

		txErr := tx.Rollback()

		if txErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to rollback transaction", txErr)
		}

		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while updating user activation info.", dberr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbUserDeleteActivationDataModel := lmodels.DBUserDeleteActivationDataModel{}
	dbUserDeleteActivationDataModel.Code = req.Code

	err, _ := dbaccess.SplMasterUsrActivationDelete(tx, &dbUserDeleteActivationDataModel)
	if err != nil {

		txErr := tx.Rollback()

		if txErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to rollback transaction", txErr)
		}

		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while deleting user activation info.", err)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	txErr := tx.Commit()

	if txErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to commit transaction", txErr)
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully updated user activation details")
	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully deleted user activation row")

	return true, nil
}
