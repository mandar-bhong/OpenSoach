package dbaccess

import (
	dbmgr "opensoach.com/core/manager/db"
	hktmodels "opensoach.com/hkt/models"
	"opensoach.com/hkt/server/constants/dbquery"
)

func TaskGetSerConfDetailsByConfInstId(instDBConn string, servConfInID int64) (error, []hktmodels.DBDeviceSerConfigModel) {

	dbDeviceSerConfigModel := []hktmodels.DBDeviceSerConfigModel{}
	selectContext := dbmgr.SelectContext{}
	selectContext.DBConnection = instDBConn
	selectContext.Query = dbquery.QUERY_GET_SER_CONFIG_BY_SER_CONF_INS_ID
	selectContext.QueryType = dbmgr.Query
	selectContext.Dest = &dbDeviceSerConfigModel

	dbErr := selectContext.Select(servConfInID)

	return dbErr, dbDeviceSerConfigModel
}

func TaskGetSerConfDetails(instDBConn string, cpmid int64, devid int64, spid int64) (error, []hktmodels.DBDeviceSerConfigModel) {

	dbDeviceSerConfigModel := []hktmodels.DBDeviceSerConfigModel{}
	selectContext := dbmgr.SelectContext{}
	selectContext.DBConnection = instDBConn
	selectContext.Query = dbquery.QUERY_GET_SER_CONFIG_BY_CPM_DEV_SP
	selectContext.QueryType = dbmgr.Query
	selectContext.Dest = &dbDeviceSerConfigModel

	dbErr := selectContext.Select(cpmid, devid, spid)

	return dbErr, dbDeviceSerConfigModel
}

func TaskGetSerConfDetailsByConfId(instDBConn string, servConfID int64) (error, []hktmodels.DBDeviceSerConfigModel) {

	dbDeviceSerConfigModel := []hktmodels.DBDeviceSerConfigModel{}
	selectContext := dbmgr.SelectContext{}
	selectContext.DBConnection = instDBConn
	selectContext.Query = dbquery.QUERY_GET_SER_CONFIG_BY_SER_CONF_ID
	selectContext.QueryType = dbmgr.Query
	selectContext.Dest = &dbDeviceSerConfigModel

	dbErr := selectContext.Select(servConfID)

	return dbErr, dbDeviceSerConfigModel
}

func TaskGetFieldOperatorDetailsByFopId(instDBConn string, fopID int64) (error, []hktmodels.DBDeviceFieldOperatorDataModel) {

	dbDeviceFieldOperatorDataModel := []hktmodels.DBDeviceFieldOperatorDataModel{}
	selectContext := dbmgr.SelectContext{}
	selectContext.DBConnection = instDBConn
	selectContext.Query = dbquery.QUERY_GET_FIELD_OPERATOR_BY_FOP_ID
	selectContext.QueryType = dbmgr.Query
	selectContext.Dest = &dbDeviceFieldOperatorDataModel

	dbErr := selectContext.Select(fopID)

	return dbErr, dbDeviceFieldOperatorDataModel
}
