package dbaccess

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	dbmgr "opensoach.com/manager/dbmanager"
	gmodel "opensoach.com/models"
)

var dbMasterEngine *sqlx.DB

func Init(dbconfig gmodel.DatabaseSettings) bool {
	var connErr error
	dbMasterEngine, connErr = sqlx.Connect("mysql", dbconfig.DBConnection)

	if connErr != nil {
		return false
	}

	return true
}

func ValidateLogin(username string, password string) (error, *[]gmodel.DBMasterUserRowModel) {

	data := &[]gmodel.DBMasterUserRowModel{}
	selDBCtx := dbmgr.SelectProcContext{}
	selDBCtx.Engine = dbMasterEngine
	selDBCtx.SPName = "sp_mst_chk_user_login"
	selDBCtx.Dest = data

	selErr := selDBCtx.Select(username, password)

	if selErr != nil {
		return selErr, nil
	}

	return nil, data
}

func GetUserProducts(userid int64) (error, *[]gmodel.DBMasterProductBriefRowModel) {
	data := &[]gmodel.DBMasterProductBriefRowModel{}
	selDBCtx := dbmgr.SelectProcContext{}
	selDBCtx.Engine = dbMasterEngine
	selDBCtx.SPName = "sp_mst_get_usr_products"
	selDBCtx.Dest = data

	selErr := selDBCtx.Select(userid)

	if selErr != nil {
		return selErr, nil
	}

	return nil, data
}
