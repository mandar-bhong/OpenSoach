package dbaccess

import (
	dbmgr "opensoach.com/core/manager/db"
	hktmodels "opensoach.com/hpft/models"
	"opensoach.com/hpft/server/constants/dbquery"
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

func TaskGetFieldOperatorByFopId(instDBConn string, fopID int64) (error, hktmodels.DBDevFieldOperatorDataModel) {

	dbDevFieldOperatorDataModel := hktmodels.DBDevFieldOperatorDataModel{}
	selectContext := dbmgr.SelectContext{}
	selectContext.DBConnection = instDBConn
	selectContext.Query = dbquery.QUERY_GET_FIELD_OPERATOR_BY_ID
	selectContext.QueryType = dbmgr.Query
	selectContext.Dest = &dbDevFieldOperatorDataModel

	dbErr := selectContext.Get(fopID)

	return dbErr, dbDevFieldOperatorDataModel
}

func TaskGetServicePointByDevId(instDBConn string, devID int64) (error, hktmodels.DBDeviceServicePointDataModel) {

	dbDeviceServicePointDataModel := hktmodels.DBDeviceServicePointDataModel{}
	selectContext := dbmgr.SelectContext{}
	selectContext.DBConnection = instDBConn
	selectContext.Query = dbquery.QUERY_GET_SERVICE_POINT_BY_DEV_ID
	selectContext.QueryType = dbmgr.Query
	selectContext.Dest = &dbDeviceServicePointDataModel

	dbErr := selectContext.Get(devID)

	return dbErr, dbDeviceServicePointDataModel
}

func TaskGetDeviceBySpID(instDBConn string, spId int64) (error, []hktmodels.DBDeviceServicePointDataModel) {

	dbDeviceServicePointDataModel := []hktmodels.DBDeviceServicePointDataModel{}
	selectContext := dbmgr.SelectContext{}
	selectContext.DBConnection = instDBConn
	selectContext.Query = dbquery.QUERY_GET_DEVICES_BY_SERVICE_POINT_ID
	selectContext.QueryType = dbmgr.Query
	selectContext.Dest = &dbDeviceServicePointDataModel

	dbErr := selectContext.Select(spId)

	return dbErr, dbDeviceServicePointDataModel
}

func TaskGetDeviceByPatientID(instDBConn string, spId int64) (error, []hktmodels.DBDeviceServicePointDataModel) {

	dbDeviceServicePointDataModel := []hktmodels.DBDeviceServicePointDataModel{}
	selectContext := dbmgr.SelectContext{}
	selectContext.DBConnection = instDBConn
	selectContext.Query = dbquery.QUERY_GET_DEVICES_BY_PATIENT_ID
	selectContext.QueryType = dbmgr.Query
	selectContext.Dest = &dbDeviceServicePointDataModel

	dbErr := selectContext.Select(spId)

	return dbErr, dbDeviceServicePointDataModel
}

func TaskGetPatientConfDetails(instDBConn string, patientId int64) (error, []hktmodels.DBDevicePatientConfigModel) {

	dbDevicePatientConfigModel := []hktmodels.DBDevicePatientConfigModel{}
	selectContext := dbmgr.SelectContext{}
	selectContext.DBConnection = instDBConn
	selectContext.Query = dbquery.QUERY_GET_PATIENT_CONFIG
	selectContext.QueryType = dbmgr.Query
	selectContext.Dest = &dbDevicePatientConfigModel

	dbErr := selectContext.Select(patientId)

	return dbErr, dbDevicePatientConfigModel
}
