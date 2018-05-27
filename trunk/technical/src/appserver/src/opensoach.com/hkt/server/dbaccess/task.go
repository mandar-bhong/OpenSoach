package dbaccess

import (
	dbmgr "opensoach.com/core/manager/db"
	//"opensoach.com/hkt/server/constants/dbquery"
)

func TaskGetSerInstDetailsById(instDBConn string, servConfIn int64) {

	selectContext := dbmgr.SelectContext{}
	selectContext.DBConnection = instDBConn
	selectContext.Query = ""
	selectContext.QueryType = dbmgr.Query
	//selectContext.Dest

	//selectContext.Select()

}
