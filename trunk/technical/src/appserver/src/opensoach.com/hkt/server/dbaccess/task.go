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
