package dbaccess

import (
	dbmgr "opensoach.com/core/manager/db"
	"opensoach.com/spl/constants"
	"opensoach.com/spl/constants/dbquery"
	lmodels "opensoach.com/spl/models"
)

func SplMasterUserTableInsert(dbConn string, insrtStruct lmodels.DBSplMasterUserTableRowModel) (error, int64) {
	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = insrtStruct
	insDBCtx.Query = dbquery.QUERY_SPL_MASTER_USER_TABLE_INSERT
	insDBCtx.QueryType = dbmgr.Query
	insDBCtx.TableName = constants.DB_TABLE_USER_TBL
	insErr := insDBCtx.Insert()
	if insErr != nil {
		return insErr, 0
	}
	return nil, insDBCtx.InsertID
}
