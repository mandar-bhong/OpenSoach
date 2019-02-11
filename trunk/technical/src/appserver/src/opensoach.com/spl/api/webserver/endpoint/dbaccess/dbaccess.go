package dbaccess

import (
	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	"opensoach.com/spl/api/constants"
	"opensoach.com/spl/api/constants/dbquery"
	lmodels "opensoach.com/spl/api/models"
)

var SUB_MODULE_NAME = "SPL.Endpoint.DB"

func ValidateDevice(dbConn, serialno string) (error, *[]lmodels.DBSplMasterDeviceTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing ValidateDevice")

	data := &[]lmodels.DBSplMasterDeviceTableRowModel{}
	selDBCtx := dbmgr.SelectContext{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_GET_DEVICE_INFO_BY_DEVICE_SERIAL_NO
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selDBCtx.TableName = constants.DB_TABLE_USER_TBL

	selErr := selDBCtx.Select(serialno)

	if selErr != nil {
		return selErr, nil
	}

	return nil, data
}

func GetDeviceAuthInfo(dbConn string, devid int64, prodcode string) (error, *[]lmodels.DBDeviceAuthInfoModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetDeviceAuthInfo")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]lmodels.DBDeviceAuthInfoModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_GET_DEVICE_AUTH_INFO
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data

	selErr := selDBCtx.Select(devid, prodcode)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func GetDeviceUserAuthInfo(dbConn string, usrname string, password string, cpmid int64) (error, *[]lmodels.DBDeviceUserAuthInfoModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetDeviceUserAuthInfo")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]lmodels.DBDeviceUserAuthInfoModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_GET_DEVICE_USER_INFO
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data

	selErr := selDBCtx.Select(usrname, password, cpmid)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}
