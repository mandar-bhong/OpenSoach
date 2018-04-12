package dbaccess

import (
	dbmgr "opensoach.com/core/manager/db"

	"opensoach.com/spl/constants"
	"opensoach.com/spl/constants/dbquery"
	lmodels "opensoach.com/spl/models"
)

var SUB_MODULE_NAME = "SPL.Login.DB"

func ValidateAuth(dbConn string, username, password string) (error, *[]lmodels.DBMasterUserRowModel) {
	filter := lmodels.AuthRequest{}
	filter.UserName = username
	filter.Password = password
	data := &[]lmodels.DBMasterUserRowModel{}
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

func GetUserAuthInfo(dbConn string, prodcode string) (error, *[]lmodels.DBUserAuthInfo) {
	filter := lmodels.DBUserAuthInfo{}
	filter.ProdCode = prodcode
	selDBCtx := dbmgr.SelectContext{}
	data := &[]lmodels.DBUserAuthInfo{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_GET_USER_AUTH_INFO
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selDBCtx.TableName = constants.DB_TABLE_PRODUCT_TBL
	selErr := selDBCtx.SelectByFilter(filter, "prod_code")
	if selErr != nil {
		return selErr, &[]lmodels.DBUserAuthInfo{}
	}
	return nil, data
}

func GetUserLoginInfo(dbConn string, userid int64) (error, *lmodels.DBUserInfoMinDataModel) {
	selDBCtx := dbmgr.SelectContext{}
	data := &lmodels.DBUserInfoMinDataModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_GET_USER_LOGIN_INFO
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selDBCtx.TableName = constants.DB_TABLE_USER_DETAILS_TBL
	selErr := selDBCtx.Get(userid)
	if selErr != nil {
		return selErr, &lmodels.DBUserInfoMinDataModel{}
	}
	return nil, data
}

func GetCustomerLoginInfo(dbConn string, customerId int64) (error, *lmodels.DBCustomerLoginInfoDataModel) {
	selDBCtx := dbmgr.SelectContext{}
	data := &lmodels.DBCustomerLoginInfoDataModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_GET_CUSTOMER_LOGIN_INFO
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selDBCtx.TableName = constants.DB_TABLE_CUSTOMER_TBL
	selErr := selDBCtx.Get(customerId)
	if selErr != nil {
		return selErr, &lmodels.DBCustomerLoginInfoDataModel{}
	}
	return nil, data
}
