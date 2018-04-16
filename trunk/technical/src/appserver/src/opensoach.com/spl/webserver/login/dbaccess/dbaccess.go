package dbaccess

import (
	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"

	"opensoach.com/spl/constants"
	"opensoach.com/spl/constants/dbquery"
	lmodels "opensoach.com/spl/models"
)

var SUB_MODULE_NAME = "SPL.Login.DB"

func ValidateAuth(dbConn string, username, password string) (error, *[]lmodels.DBSplMasterUserTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing ValidateAuth")

	filter := lmodels.AuthRequest{}
	filter.UserName = username
	filter.Password = password
	data := &[]lmodels.DBSplMasterUserTableRowModel{}
	selDBCtx := dbmgr.SelectContext{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_MUST_CHECK_USER_LOGIN
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selDBCtx.TableName = constants.DB_TABLE_USER_TBL

	selErr := selDBCtx.SelectByFilter(filter, "usr_name", "usr_password")

	if selErr != nil {
		return selErr, nil
	}

	return nil, data
}

func GetUserAuthInfo(dbConn string, prodcode string, userid int64) (error, *[]lmodels.DBUserAuthInfo) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetUserAuthInfo")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]lmodels.DBUserAuthInfo{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_GET_USER_AUTH_INFO
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data

	selErr := selDBCtx.Select(prodcode, userid)
	if selErr != nil {
		return selErr, &[]lmodels.DBUserAuthInfo{}
	}
	return nil, data
}

func GetUserLoginInfo(dbConn string, userid int64) (error, *lmodels.DBUserInfoMinDataModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetUserLoginInfo")

	selDBCtx := dbmgr.SelectContext{}
	data := &lmodels.DBUserInfoMinDataModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_GET_USER_LOGIN_INFO
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Get(userid)
	if selErr != nil {
		return selErr, &lmodels.DBUserInfoMinDataModel{}
	}
	return nil, data
}

func GetCustomerLoginInfo(dbConn string, customerId int64) (error, *lmodels.DBCustomerLoginInfoDataModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetCustomerLoginInfo")

	selDBCtx := dbmgr.SelectContext{}
	data := &lmodels.DBCustomerLoginInfoDataModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_GET_CUSTOMER_LOGIN_INFO
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Get(customerId)
	if selErr != nil {
		return selErr, &lmodels.DBCustomerLoginInfoDataModel{}
	}
	return nil, data
}
