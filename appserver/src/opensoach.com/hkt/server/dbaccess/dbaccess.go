package dbaccess

import (
	dbmgr "opensoach.com/core/manager/db"
	"opensoach.com/hkt/server/constants/dbquery"
)

var SUB_MODULE_NAME = "HKT.Server.DB"

func GetInstanceDBConn(mstDBConn string, cpmid int64) (error, string) {

	dbInstConn := ""

	selectContext := dbmgr.SelectContext{}
	selectContext.DBConnection = mstDBConn
	selectContext.Dest = &dbInstConn
	selectContext.Query = dbquery.QUERY_GET_DB_CONN_BY_CPM_ID
	selectContext.QueryType = dbmgr.Query

	err := selectContext.Get(cpmid)

	return err, dbInstConn
}
