package dbaccess

import (
	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	"opensoach.com/splserver/constants"
	dbquery "opensoach.com/splserver/constants/dbquery"
	lmodels "opensoach.com/splserver/models"
)

func GetEmailTemplate(dbConn string, code string) (error, *lmodels.DBEmailTemplateRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing get email template")

	selDBCtx := dbmgr.SelectContext{}
	data := &lmodels.DBEmailTemplateRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_SELECT_EMAIL_TML_BY_CODE
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data

	err := selDBCtx.Get(code)

	return err, data

}

func SaveEmail(dbConn string, dbEmailRowModel lmodels.DBEmailRowModel) (error, int64) {

	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = dbEmailRowModel
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_TABLE_SPL_MASTER_EMAIL_TBL
	insErr := insDBCtx.Insert()
	if insErr != nil {
		return insErr, 0
	}
	return nil, insDBCtx.InsertID

}

func UpdateEmailStatus(dbConn string, dbEmailRowModel lmodels.DBEmailRowModel) error {
	updateDBCtx := dbmgr.UpdateDeleteContext{}
	updateDBCtx.Args = dbEmailRowModel
	updateDBCtx.DBConnection = dbConn
	updateDBCtx.Query = dbquery.QUERY_UPDATE_EMAIL_EMAIL_STATUS
	updateDBCtx.QueryType = dbmgr.Query

	err := updateDBCtx.Update()

	return err

}
